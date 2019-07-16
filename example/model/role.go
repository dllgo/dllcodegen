package model

import "github.com/dllgo/dllkit/gins"

type Role struct {
	gins.Model
	Name string `gorm:"char(60)" form:"name" json:"name"`
	Mids string `gorm:"char(100)" form:"mids" json:"-"`
	Menu []Menu `json:"menu" gorm:"many2many:role_menu;"`
}
