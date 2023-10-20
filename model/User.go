package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(20)" json:"UserName"`
	PassWord string `gorm:"type:varchar(20)" json:"PassWord"`
	Role     int    `gorm:"type:int" json:"Role"`
}
