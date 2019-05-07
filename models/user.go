package models

import (
	"fmt"
	"go-restful-project/pkg/util"
	"time"

	"github.com/jinzhu/gorm"
)

type AddUserParam struct {
	UserName string `valid:"required,length(3|36)~仅允许3-36个字符的用户名"`
	Password string `valid:"required,length(3|36)~仅允许3-36个字符的密码"`
	Name     string `valid:"runelength(2|36)~仅允许2-36个字符的姓名"`
	Phone    string `valid:"phone~电话号码格式不正确"`
	Remarks  string `valid:"runelength(0|200)"`
	Role     int    `valid:"required~角色ID不能为空"`
}

type EditUserParam struct {
	Password string `valid:"length(3|36)~仅允许3-36个字符的密码“`
	Name     string `valid:"runelength(2|36)~仅允许2-36个字符的姓名"`
	Phone    string `valid:"phone~电话号码格式不正确"`
	Remarks  string `valid:"runelength(0|200)"`
	Role     int
}

type Auth struct {
	Username string `json:"username" valid:"Required; MaxSize(50)" binding:"required"`
	Password string `json:"password" valid:"Required; MaxSize(50)" binding:"required"`
}

type Userprofile struct {
	BaseModel
	Username    string        `gorm:"column:username;not null;unique;type:varchar(47)" json:"username"`
	Password    string        `gorm:"column:password;not null;type:varchar(128)" json:"password"`
	FirstName   string        `gorm:"column:first_name;not null;type:varchar(30)" json:"-"`
	LastName    string        `gorm:"column:last_name;not null;type:varchar(150)" json:"-"`
	Email       string        `gorm:"column:email;not null;type:varchar(254)" json:"email"`
	IsSuperuser bool          `gorm:"column:is_superuser;not null;" json:"is_superuser"`
	IsStaff     bool          `gorm:"column:is_staff;not null;" json:"-"`
	IsActive    bool          `gorm:"column:is_active;default:true;not null;" json:"-"`
	LastLogin   util.JSONTime `gorm:"column:last_login" json:"-"`
	DateJoined  util.JSONTime `gorm:"column:date_joined";not null; json:"-"`
	Name        string        `gorm:"column:name;not null;type:varchar(47)" json:"name"`
	Phone       string        `gorm:"column:phone;not null;type:varchar(18)" json:"phone"`
	Remarks     string        `gorm:"column:remarks;type:varchar(200)" json:"remarks"`
	RoleID      int           `gorm:"column:role_id;not null;default:1" json:"role_id"`
}

// Users 返回前端的用户列表
type Userslist struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	IsSuperuser bool   `json:"is_superuser"`
	Remarks     string `json:"remarks"`
	RoleName    string `json:"role_name"`
}

// User 用户详情
type UserDetails struct {
	ID          int           `json:"id"`
	Username    string        `json:"username"`
	Name        string        `json:"name"`
	Phone       string        `json:"phone"`
	IsSuperuser bool          `json:"is_superuser"`
	Remarks     string        `json:"remarks"`
	Role        int           `json:"role"`
	RoleName    string        `json:"role_name"`
	CreateTime  util.JSONTime `json:"create_time"`
	UpdateTime  util.JSONTime `json:"update_time"`
}

// CheckAuth 检查身份验证信息是否存在
func CheckAuth(username, password string) (bool, error) {
	// db.AutoMigrate(&User{})
	var user Userprofile
	err := db.Select("id").Where(Userprofile{Username: username, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

// GetUserTotal 获取用户总条数
func GetUserTotal(search string) (int, error) {
	var (
		count int
		err   error
	)
	if search == "" {
		err = db.Model(&Userprofile{}).Where("is_del = ?", false).Count(&count).Error
	} else {
		search = fmt.Sprintf("%%%s%%", search)
		fmt.Println(search)
		err = db.Model(&Userprofile{}).Where("name like ? AND is_del = ?", search, false).Or("username like ?", search).Or("phone like ?", search).Count(&count).Error
	}
	if err != nil {
		return 0, err
	}
	return count, nil
}

// ExistUserByID根据该ID确定是否存在标记
func ExistUserByID(id int) (bool, error) {
	var user Userprofile
	err := db.Select("id").Where("id = ? AND is_del = ? ", id, false).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

// ExistUserByID根据该ID确定是否存在标记
func ExistUserByUserName(username string) (bool, error) {
	var user Userprofile
	err := db.Select("id").Where("username = ? AND is_del = ? ", username, false).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetUsers 获取所有用户
func GetUsers(pageNum int, pageSize int, search string) ([]Userprofile, error) {

	var (
		users []Userprofile
		err   error
	)
	// filter := maps.(map[string]string)
	if search == "" {
		if pageSize > 0 {
			err = db.Where("is_del = ?", false).Offset(pageNum).Limit(pageSize).Order("id desc").Find(&users).Error
		} else {
			err = db.Where("is_del = ?", false).Order("id desc").Find(&users).Error
		}
	} else {
		search = fmt.Sprintf("%%%s%%", search)
		if pageSize > 0 {
			err = db.Where("name like ? AND is_del = ?", search, false).Or("username like ?", search).
				Or("phone like ?", search).Offset(pageNum).Limit(pageSize).Order("id desc").Find(&users).Error
		} else {
			err = db.Where("name like ? AND is_del = ?", search, false).Or("username like ?", search).
				Or("phone like ?", search).Order("id desc").Find(&users).Error
		}
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}

// GetUser 获取单个用户
func GetUser(id int) (*Userprofile, error) {
	var user Userprofile
	err := db.Where("id = ? AND is_del = ? ", id, false).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &user, nil
}

// AddUser 添加用户
// func AddUser(username, password, name, phone, remarks string, role int) (*Userprofile, error) {
func AddUser(user AddUserParam) (*Userprofile, error) {
	userprofile := Userprofile{
		Username: user.UserName,
		Password: user.Password,
		Name:     user.Name,
		Phone:    user.Phone,
		Remarks:  user.Remarks,
		RoleID:   user.Role,
	}
	if err := db.Create(&userprofile).Error; err != nil {
		return nil, err
	}

	return &userprofile, nil
}

// UpdateUser 修改用户信息
func UpdateUser(id int, data interface{}) error {
	if err := db.Model(&Userprofile{}).Where("id = ? AND is_del = ? ", id, false).Updates(data).Error; err != nil {
		return err
	}
	return nil
}

// DeleteUser 删除用户
func DeleteUser(id int) error {
	data := make(map[string]interface{})
	user, err := GetUser(id)
	if err != nil {
		return err
	}
	t := time.Now()
	data["username"] = fmt.Sprintf("%s_%d", user.Username, t.Unix())
	data["is_del"] = true
	data["del_time"] = t
	if err := db.Model(&Userprofile{}).Where("id = ? AND is_del = ? ", id, false).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
