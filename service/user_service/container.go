package user_service

import (
	"github.com/emicklei/go-restful"
)

type AddUserParam struct {
	username string
	password string
	name     string
	phone    string
	remarks  string
	role     int
}

type EditUserParam struct {
	username string
	name     string
	phone    string
	remarks  string
	role     int
}

func UsersRegister(container *restful.Container) {
	ws := new(restful.WebService)

	ws.
		Path("/api/v1/users").
		Doc("Users").
		Consumes(restful.MIME_JSON, restful.MIME_XML).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("").To(GetUsers).
		Doc("获取用户列表").
		Param(ws.QueryParameter("serach", "模糊搜索用户名，姓名，电话号码").DataType("string")))

	ws.Route(ws.GET("/{user-id}").To(GetUser).
		Doc("获取用户详情").
		Param(ws.PathParameter("user-id", "用户ID").DataType("string")))

	ws.Route(ws.POST("").To(CreateUser).
		Doc("添加用户").
		Reads(AddUserParam{}, "username:3-26个字母/数字/下划线/—_组成的字符\r\npassword:3-36\r\nname:3-36"))
	// Param(ws.BodyParameter("username", "用户名（3-26个字母/数字/下划线/—_组成的字符）").DataType("string")).
	// Param(ws.BodyParameter("password", "密码").DataType("string")).
	// Param(ws.BodyParameter("name", "姓名").Required(false).DataType("string")).
	// Param(ws.BodyParameter("role", "角色").DataType("int")).
	// Param(ws.BodyParameter("phone", "电话号码").Required(false).DataType("string")).
	// Param(ws.BodyParameter("remarks", "备注").Required(false).DataType("string")).

	ws.Route(ws.PUT("/{user-id}").To(UpdateUser).
		Doc("更新用户").
		Param(ws.PathParameter("user-id", "用户ID").DataType("string")).
		Reads(EditUserParam{}))

	ws.Route(ws.DELETE("/{user-id}").To(DeleteUser).
		Doc("删除用户").
		Param(ws.PathParameter("user-id", "用户ID").DataType("string")))

	container.Add(ws)

}
