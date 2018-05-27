package models

import (
	"strings"
	"strconv"
	"fmt"
	"time"
	"github.com/jinzhu/gorm"
)
type Scores struct {
	UserName string
	Score int
	CreatedAt	time.Time
}


//Result ...Result Model
type Result struct {
	gorm.Model
	// GORM SOESANOT SUPPORT DATABASE WELL, including that it cannot handle foreign key or drop column
	QnID     int `gorm:"ForeignKey:ID"`
	UserName  string
	Answer string
	Code  string
	Score int
	Status  bool
	Filename string
}

func (Result) TableName() string {
	return "result"
}

func (res Result) Save() (error) {
	err := db.Save(&res).Error
		return err
}

func BonusEligible(qn_id int) bool {
	var res int
	db.Raw("SELECT count(*) FROM result WHERE qn_id = ?", qn_id).Scan(&res)
	if res == 0 {
		return false
	}else{
		return true
	}
}


func QuestionAttemptedBy(qn_id uint) string {
	var res int
	var names []string
	var sliced string
	db.Raw("SELECT count(DISTINCT user_name) FROM result WHERE qn_id = ?", qn_id).Scan(&res)
	if res > 0{
		db.Raw("SELECT DISTINCT user_name FROM result WHERE qn_id = ? LIMIT 3", qn_id).Scan(&names)
		sliced = strings.Join(names, ",")
		if res > 3{
   		sliced= sliced + " and " + strconv.Itoa(res - 3) + " others Attempted"
		}else{sliced = sliced + " Attempted"}
	}else{sliced = "None Attempted"}
	return sliced
}

func QuestionCompletedBy(qn_id uint) string {
	var res int
	var names []string
	var sliced string
	err := db.Raw("SELECT count(DISTINCT user_name) FROM result WHERE qn_id = ? AND status = true  ", qn_id).Scan(&res).Error
	if err != nil {return "None Attempted" }
	if res > 0{
		db.Raw("SELECT DISTINCT user_name FROM result WHERE qn_id = ? LIMIT 3", qn_id).Scan(&names)
		sliced = strings.Join(names, ",")
		if res > 3{
   		sliced= sliced + " " + strconv.Itoa(res - 3) + " others Completed"
			}else{sliced = sliced + " Completed"}
		}else{sliced = "None Attempted"}
	return sliced
}

func BonusCapturedBy(qn_id uint) string {
	var user string
	err := db.Raw("SELECT user_name FROM result WHERE qn_id = ? AND status = true ORDER BY id LIMIT 1", qn_id).Scan(&user).Error
	if err != nil {user = "" }
	return user
}


func GetResults(condition interface{}) (res []Result, e error) {
	db = Model
	if err := db.Where(condition).Find(&res).Error; err != nil {
		return res, err
	}else{
		return res, err
	}
}


func TopScores() ([]Scores, error) {
	var meta []Scores
	err := db.Raw("SELECT user_name, sum(score) as score, max(created_at) as created_at FROM result GROUP BY user_name ORDER BY sum(score) DESC").Scan(&meta).Error
	fmt.Println(err)
	fmt.Println(meta)
	// fmt.Println(meta[0])
	// fmt.Println(meta[0].UserName)

	return meta, err
	}
