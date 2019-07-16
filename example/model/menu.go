package model

import "github.com/dllgo/dllkit/gins"

type Menu struct {
	gins.Model
	Name   string `gorm:"char(50)" form:"name" json:"name"`
	Path   string `gorm:"char(50)" form:"path" json:"path"`
	Method string `gorm:"char(50)" form:"method" json:"method"`
}
