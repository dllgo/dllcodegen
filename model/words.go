package model

import "github.com/jinzhu/gorm"

type Words struct {
	gorm.Model
	Name         string `gorm:"size:100;unique" json:"name" form:"name"`
	Phonetic     string `gorm:"size:100" json:"phonetic" form:"phonetic"`
	PhoneticLink string `gorm:"type:text" json:"phoneticlink" form:"phoneticlink"`
	Form         string `gorm:"size:100;" json:"form" form:"form"`
	Mean         string `gorm:"size:100" json:"mean" form:"mean"`
}
