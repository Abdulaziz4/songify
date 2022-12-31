package model

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Name string
	Url  string
}
