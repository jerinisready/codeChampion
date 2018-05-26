package models

import (
	"github.com/jinzhu/gorm"
)

//Question ...Question Model
type Question struct {
	gorm.Model
	Title     string
	Description  string
	Score int
	Bonus  int
	Input1  string
	Output1 string
	Input2  string
	Output2 string
	Input3  string
	Output3 string
}

func (Question) TableName() string {
	return "question"
}
