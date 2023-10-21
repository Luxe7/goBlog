package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string    `gorm:"type:varchar(20);not null" json:"name"`
	Articles []Article `gorm:"foreignKey:Cid"` // 一对多关联关系，外键为CategoryID
}
