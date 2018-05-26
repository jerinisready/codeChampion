package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Result ...Result Model
type Result struct {
	gorm.Model
	Qn_id     Question `gorm:"ForeignKey:ID"`
	User  string
	Answer string
	Code  string
	Created_time  time.Time
	Score int
	Status  bool
	Filename string
}
