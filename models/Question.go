package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
// var db  *gorm.DB

//Question ...Question Model
type Question struct {
	gorm.Model
	// Id          uint		`gorm:"primary_key"`
	Title     string		`gorm:"type:varchar(255)"`
	Description  string		`gorm:"type:varchar(1000)"`
	Score int
	Bonus  int
	Input1  string 	`gorm:"type:varchar(255)"`
	Output1 string	`gorm:"type:varchar(255)"`
	Input2  string 	`gorm:"type:varchar(255)"`
	Output2 string	`gorm:"type:varchar(255)"`

}

type QuestionSet struct {

	Title string
	Description string
	Score	int
	Bonus int
	AttemptedBy string
	SolvedBy string
	BonusCapturedBy string
}


func (Question) TableName() string {
	return "question"
}

func (cu Question) Save() (error) {
	err := db.Save(&cu).Error
		return err
}


func GetQuestions() ([]Question, error) {
	var qn []Question
	err := db.Find(&qn).Error
	fmt.Println(len(qn))
	fmt.Println("**********************")
	fmt.Println(err)
	return  qn, err
}

func GetQuestionWithID(id int)(qn Question, e error){
	err := db.First(&qn, id).Error
	return qn, err
}
