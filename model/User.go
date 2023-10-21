package model

import (
	"encoding/base64"
	"goBlog/utlis/errormsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20);column:username" json:"username"`
	PassWord string `gorm:"type:varchar(20);column:password" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
}

func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("UserName = ?", name).First(&user)
	if user.ID > 0 {
		return errormsg.ERROR_USERNAME_USED
	}
	return errormsg.SUCCESS
}
func CreateUser(data *User) int {
	data.PassWord = ScryptPw(data.PassWord) //可以使用grom的钩子函数完成同样的操作 BeforeSaved
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}
func GetUsers(pageSize int, pageNum int) []User {
	var Users []User
	// 不手动设置分页的话，使用Postman和grom是返回为nil的，但是某些情况下也许是有默认值的，建议手动设置分页参数
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&Users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return Users
}
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{76, 78, 02, 03, 32, 67, 43, 32}
	HashPw, err := scrypt.Key([]byte(password), salt, 1024, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}
func DeleteUser(id int) int {
	//可以删除不存在的用户并且返回ok，原因未知
	/*如果数据库中不存在某id的值，这段代码返回的错误值是nil而不是ErrRecordNotFound，
	这是因为gorm库的Delete方法只会在执行SQL语句时出现错误才会返回非nil的错误值。
	如果SQL语句执行成功，但没有匹配到任何记录，Delete方法仍然会返回nil错误，表示删除操作没有发生*/

	//已解决：使用RowsAffected判断受影响的行数
	var users User
	//err := db.Where("id = ?", id).Delete(&users).Error
	//if err != nil {
	//	return errormsg.ERROR
	//}
	//return errormsg.SUCCESS
	result := db.Where("id = ?", id).Delete(&users)
	err := result.Error
	if err != nil {
		return errormsg.ERROR
	} else {
		if result.RowsAffected == 0 {
			return errormsg.ERROR_USER_NOT_EXIST
		}
		return errormsg.SUCCESS
	}
}
