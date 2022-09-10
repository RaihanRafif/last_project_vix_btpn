package models

import "github.com/jinzhu/gorm"

type Photo struct {
	gorm.Model
	Title    string
	Caption  string
	PhotoUrl string

	User   User `gorm:"association_foreignkey:UserId"`
	UserId uint `gorm:"default:null"`
}
