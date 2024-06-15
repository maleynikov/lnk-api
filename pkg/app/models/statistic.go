package models

import (
	"gorm.io/gorm"
)

type Statistic struct {
	gorm.Model
	UrlID uint
	Url   Url `gorm:"foreignKey:UrlID"`
	IP    string
}
