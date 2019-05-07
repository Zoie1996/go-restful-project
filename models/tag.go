package models

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	BaseModel

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// ExistTagByName检查是否存在具有相同名称的标记
func ExistTagByName(name string) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("name = ? AND is_del = ? ", name, false).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// 添加一个标签
func AddTag(name string, state int, createdBy string) error {
	tag := Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	if err := db.Create(&tag).Error; err != nil {
		return err
	}

	return nil
}

// GetTags根据分页和约束获取标记列表
func GetTags(pageNum int, pageSize int, maps interface{}) ([]Tag, error) {
	var (
		tags []Tag
		err  error
	)

	if pageSize > 0 && pageNum >= 0 {
		err = db.Where(maps).Find(&tags).Limit(pageSize).Offset(pageNum).Error
	} else {
		err = db.Where(maps).Find(&tags).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return tags, nil
}

// GetTagTotal根据约束计算标签的总数
func GetTagTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Tag{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// ExistTagByID根据该ID确定是否存在标记
func ExistTagByID(id int) (bool, error) {
	var tag Tag
	err := db.Select("id").Where("id = ? AND is_del = ? ", id, false).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if tag.ID > 0 {
		return true, nil
	}

	return false, nil
}

// 删除标签
func DeleteTag(id int) error {
	if err := db.Where("id = ?", id).Delete(&Tag{}).Error; err != nil {
		return err
	}

	return nil
}

// EditTag修改单个标签
func EditTag(id int, data interface{}) error {
	if err := db.Model(&Tag{}).Where("id = ? AND is_del = ? ", id, false).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// 清除所有标记
func CleanAllTag() (bool, error) {
	if err := db.Unscoped().Where("is_del != ? ", false).Delete(&Tag{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
