package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	//Category Category `gorm:"foreignKey:Cid"`

	Title   string `gorm:"type:varchar(18);not null" json:"Title"`
	Cid     int    `gorm:"type:int;not null" json:"Cid"`
	Desc    string `gorm:"type:varchar(200)" json:"Desc"`
	Context string `gorm:"type:longtext" json:"Context"`
	Img     string `gorm:"type:varchar(100)" json:"Img"`
}
