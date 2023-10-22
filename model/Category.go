package model

import (
	"goBlog/utlis/errormsg"
	"gorm.io/gorm"
)

type Category struct {
	ID   int    `gorm:"primary_key;auto_increment"json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?", name).First(&cate)
	if cate.ID > 0 {
		return errormsg.ERROR_CATEGORY_USED
	}
	return errormsg.SUCCESS
}
func CreateCategory(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}
func GetCategory(pageSize int, pageNum int) []Category {
	var cate []Category
	// 不手动设置分页的话，使用Postman和grom是返回为nil的，但是某些情况下也许是有默认值的，建议手动设置分页参数
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return cate
}

func UpdateCategory(id int, data *Category) int {
	//gorm中用struct更新信息，不会更新值为0的参数
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	err := db.Model(&cate).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}
func DeleteCategory(id int) int {
	var cate Category

	result := db.Where("id = ?", id).Delete(&cate)
	err := result.Error
	if err != nil {
		return errormsg.ERROR
	} else {
		if result.RowsAffected == 0 {
			return errormsg.ERROR_CATEGORY_NOT_EXIST
		}
		return errormsg.SUCCESS
	}
}
