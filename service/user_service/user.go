package user_service

import (
	"go-restful-project/models"
	"go-restful-project/pkg/e"
	"go-restful-project/pkg/util"
	"go-restful-project/service/role_service"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/emicklei/go-restful"
)

func CheckUserExist(id int) int {
	exists, err := models.ExistUserByID(id)
	if err != nil {
		return e.ERROR
	}
	if !exists {
		return e.NOT_FOUND

	}
	return e.SUCCESS
}

// GetUser 根据ID获取单个用户
func GetUser(request *restful.Request, response *restful.Response) {
	// response.AddHeader("Content-Type", restful.MIME_JSON)
	id := request.PathParameter("user-id")
	user_id, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[GetUser] id %s", err)
	}

	// 检查用户是否存在
	code := CheckUserExist(user_id)
	if code != e.SUCCESS {
		response.WriteHeaderAndEntity(http.StatusNotFound, e.GetInfo(code))
		return
	}

	// 获取用户详情
	userprofile, err := models.GetUser(user_id)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusNotFound, e.GetInfo(err))
		return
	}
	user := models.UserDetails{
		ID:          userprofile.ID,
		Username:    userprofile.Username,
		Name:        userprofile.Name,
		Phone:       userprofile.Phone,
		IsSuperuser: userprofile.IsSuperuser,
		Remarks:     userprofile.Remarks,
		Role:        userprofile.RoleID,
		// RoleName:    role_name,
		CreateTime: userprofile.CreateTime,
		UpdateTime: userprofile.CreateTime,
	}
	response.WriteEntity(user)
}

// GetUsers 获取用户列表
func GetUsers(request *restful.Request, response *restful.Response) {
	search := request.PathParameter("search")
	page := util.GetPage(request)
	page_size := util.GetPageSize(request)

	users, err := models.GetUsers(page, page_size, search)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, e.GetInfo(e.ERROR_GET_USERS_FAIL))
		return
	}
	users_list := []models.Userslist{}
	for _, u := range users {
		role_name, err := models.GetRoleName(u.RoleID)
		if err != nil {
			response.WriteHeaderAndEntity(http.StatusBadRequest, e.GetInfo(e.ERROR_GET_ROLE_NAME_FAIL))
			return
		}
		user := models.Userslist{
			ID:          u.ID,
			Username:    u.Username,
			Name:        u.Name,
			Phone:       u.Phone,
			IsSuperuser: u.IsSuperuser,
			Remarks:     u.Remarks,
			RoleName:    role_name,
		}
		users_list = append(users_list, user)
	}

	response.WriteEntity(users_list)

}

// CreateUser 创建用户
func CreateUser(request *restful.Request, response *restful.Response) {
	usr := new(models.AddUserParam)
	err := request.ReadEntity(&usr)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusNotFound, e.GetInfo(err))
		return
	}
	govalidator.TagMap["phone"] = govalidator.Validator(func(str string) bool {
		match, _ := regexp.MatchString("^((0\\d{2,3}[-]?)?[2-9]\\d{6,7}([-]?\\d{1,4})?)|(1[345789]\\d{9})$", str)
		return match
	})
	_, err = govalidator.ValidateStruct(usr)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusNotFound, e.GetInfo(err))
		return
	}
	// 检查是否存在相同的用户名
	exists, err := models.ExistUserByUserName(usr.UserName)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusNotFound, e.GetInfo(err))
		return
	}
	if exists {
		response.WriteHeaderAndEntity(http.StatusNotFound, e.GetInfo("具有该名称的 UserName 已存在"))
		return
	}
	// 检查角色是否存在
	code := role_service.CheckRoleExist(usr.Role)
	if code != e.SUCCESS {
		response.WriteHeaderAndEntity(http.StatusNotFound, e.GetInfo(code))
		return
	}
	userprofile, err := models.AddUser(*usr)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, e.GetInfo(err))
		return
	}
	user, err := models.GetUser(userprofile.ID)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, e.GetInfo(e.ERROR_GET_USER_FAIL))
		return
	}
	response.WriteEntity(user)
}

// UpdateUser 更新用户
func UpdateUser(request *restful.Request, response *restful.Response) {
	// 获取用户id
	id := request.PathParameter("user-id")
	user_id, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[UpdateUser] id %s", err)
	}
	// 检查用户是否存在
	code := CheckUserExist(user_id)
	if code != e.SUCCESS {
		response.WriteHeaderAndEntity(http.StatusNotFound, e.GetInfo(code))
		return
	}

	usr := new(models.EditUserParam)
	err = request.ReadEntity(&usr)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, e.GetInfo(err))
		return
	}

	// 自定义电话号码验证
	govalidator.TagMap["phone"] = govalidator.Validator(func(str string) bool {
		match, _ := regexp.MatchString("^((0\\d{2,3}[-]?)?[2-9]\\d{6,7}([-]?\\d{1,4})?)|(1[345789]\\d{9})$", str)
		return match
	})
	// 验证请求体
	_, err = govalidator.ValidateStruct(usr)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusNotFound, e.GetInfo(err))
		return
	}

	// 检查角色是否存在
	if usr.Role != 0 {
		code = role_service.CheckRoleExist(usr.Role)
		if code != e.SUCCESS {
			response.WriteHeaderAndEntity(http.StatusBadRequest, e.GetInfo(code))
			return
		}
	}

	data := make(map[string]interface{})
	if len(usr.Password) > 0 {
		data["password"] = usr.Password
	}
	data["name"] = usr.Name
	data["phone"] = usr.Phone
	data["role"] = usr.Role
	data["remarks"] = usr.Remarks

	err = models.UpdateUser(user_id, data)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, e.GetInfo(err))
		return
	}
	user, err := models.GetUser(user_id)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, e.GetInfo(e.ERROR_GET_USER_FAIL))
		return
	}
	response.WriteEntity(user)
}

// DeleteUser 删除用户
func DeleteUser(request *restful.Request, response *restful.Response) {
	// 获取用户id
	id := request.PathParameter("user-id")
	user_id, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("[DeleteUser] id %s", err)
	}
	// 检查用户是否存在
	code := CheckUserExist(user_id)
	if code != e.SUCCESS {
		response.WriteHeaderAndEntity(http.StatusNotFound, e.GetInfo(code))
		return
	}
	err = models.DeleteUser(user_id)
	if err != nil {
		response.WriteHeaderAndEntity(http.StatusBadRequest, e.GetInfo(err))
		return
	}
	response.WriteHeader(http.StatusNoContent)
}
