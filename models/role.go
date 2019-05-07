package models

import (
	"go-restful-project/pkg/util"

	"github.com/jinzhu/gorm"
)

type Role struct {
	BaseModel
	Name          string `gorm:"column:name;not null;unique;type:varchar(47);" json:"name"`
	Remarks       string `gorm:"column:remarks;type:varchar(200);" json:"remarks"`
	Menus         string `gorm:"column:menus;not null;type:varchar(200);" json:"menus"`
	IsSystemRole  bool   `gorm:"column:is_system_role;default:false;" json:"is_system_role"`
	IsAppletLogin bool   `gorm:"column:is_applet_login;default:false;" json:"is_applet_login"`
	// 区域为空时,默认是所有区域
	Area    string `gorm:"column:areas;" json:"areas"`
	Devices string `gorm:"column:devices;" json:"devices"`
}
type RoleDetails struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Remarks       string   `json:"remarks"`
	Menus         []string `json:"menus"`
	IsSystemRole  bool     `json:"is_system_role"`
	IsAppletLogin bool     `json:"is_applet_login"`
	// 区域为空时,默认是所有区域
	Area       []map[string]interface{} `json:"areas"`
	Devices    []int                    `json:"devices"`
	CreateTime util.JSONTime            `json:"create_time"`
	UpdateTime util.JSONTime            `json:"update_time“`
}

// ExistRoleByID根据该ID确定是否存在标记
func ExistRoleByID(id int) (bool, error) {
	var role Role
	err := db.Select("id").Where("id = ? AND is_del = ? ", id, false).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if role.ID > 0 {
		return true, nil
	}

	return false, nil
}

// GetRole 通过ID获取单个角色
func GetRoleName(id int) (string, error) {
	var role Role
	err := db.Where("id = ? AND is_del = ? ", id, false).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "", err
	}
	return role.Name, nil
}

// GetRole 通过ID获取单个角色
func GetRole(id int) (*Role, error) {
	var role Role
	err := db.Where("id = ? AND is_del = ? ", id, false).First(&role).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &role, nil
}

// func AddRole(Rolename, password, name, phone, remarks string, role int) (*Roleprofile, error) {
func AddRole(name, remarks, area, menus, devices string, isSystemRole, isAppletLogin bool) (*Role, error) {
	db.AutoMigrate(&Userprofile{}, &Role{})
	Role := Role{
		Name:          name,
		Remarks:       remarks,
		Menus:         menus,
		IsSystemRole:  isSystemRole,
		IsAppletLogin: isAppletLogin,
		// 区域为空时,默认是所有区域
		Area:    area,
		Devices: devices,
	}
	if err := db.Create(&Role).Error; err != nil {
		return nil, err
	}

	return &Role, nil
}
