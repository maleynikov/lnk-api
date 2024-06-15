package models

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	Code  string `gorm:"unique"`
	Value string
}
