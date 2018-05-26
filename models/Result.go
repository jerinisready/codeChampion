package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Result ...Result Model
type Result struct {
	gorm.Model
	Qn_id     Question `gorm:"ForeignKey:ID"`
	User  CodeUser `gorm:"ForeignKey:ID"`
	Answer string
	Code  string
	Created_time  time.Time
	Score int
	Status  bool
	Filename string
}

func (Result) TableName() string {
	return "result"
}
