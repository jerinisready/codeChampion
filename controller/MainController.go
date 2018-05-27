package controllers

import (
	"errors"
	"fmt"
	"go-webapp/models"
	"go-webapp/compileunit"
	"go-webapp/common"
	"github.com/gin-gonic/gin"
	"go-webapp/middleware/authenticate"
	log "github.com/sirupsen/logrus"

)

type Login struct {
	Username     string `form:"username" json:"username" binding:"required"`
	Password 		 string `form:"password" json:"password" binding:"required"`
}

type AddQuestion struct {
	title     string `form:"Title" json:"Title" binding:"required"`
	description     string `form:"Description" json:"Description" binding:"required"`
	score     int `form:"Score" json:"Score" binding:"required"`
	bonus     int `form:"Bonus" json:"Bonus" binding:"required"`
	input1     string `form:"Input1" json:"Input1" binding:"required"`
	output1    string `form:"Output1" json:"Output1" binding:"required"`
	input2     string `form:"Input2" json:"Input2" binding:"required"`
	output2    string `form:"Output2" json:"Output2" binding:"required"`
}

type Results struct {
	qn_id    int	`form:"qn_id" json:"qn_id" binding:"required"`
	script	 string	`form:"script" json:"script" binding:"required"`
	filename string	`form:"filename" json:"filename" binding:"required"`
	lang		 string	`form:"lang" json:"lang" binding:"required"`
}

// Login API
func LoginAPI(c *gin.Context) {
	var json Login
	fmt.Println("")

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "`username` or `password` field is missing" , "success": false})
	} else {
	  userModel, err := models.CodeUserGet(&models.CodeUser{Username: json.Username})
		if err != nil {
			var userModel models.CodeUser
			userModel.Username = json.Username
			userModel.Password = json.Password
			 err := userModel.Save()
			if err != nil {
				c.JSON(400, common.NewError("register", errors.New("Something Went Wrong on Registring an Account")))
			}else{
				authentication.SetSession(c, userModel.Username)
				c.JSON(200, gin.H{"success": true, "error": "", "username": userModel.Username})
			}
			return
		} else {
			if userModel.Password == json.Password{
				authentication.SetSession(c, userModel.Username)
				c.JSON(200, gin.H{"success": true, "error": "", "username": userModel.Username})
				return
			}else{
				c.JSON(400, gin.H{"error": "`password` went wrong" , "success": false})
			}
		}
	}
}


// Login API
func QuestionSet(c *gin.Context) {
	var data []models.Question
	var qns []models.QuestionSet
	var qn models.QuestionSet

	data, err := models.GetQuestions()
	if err != nil{
			log.WithFields(log.Fields{
				"error": err,
			}).Warn("Get Questions")
	}
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Info("Get Questions")
		c.JSON(400, gin.H{"error": err })
	}else{
	for _, elem := range data {
		fmt.Println(elem.ID)
		fmt.Println("elem.ID")
		qn.ID = elem.ID
		qn.Title = elem.Title
		qn.Description = elem.Description
		qn.Score	= elem.Score
		qn.Bonus = elem.Bonus
		qn.AttemptedBy = models.QuestionAttemptedBy(elem.ID)
		qn.SolvedBy = models.QuestionCompletedBy(elem.ID)
		qn.BonusCapturedBy = models.BonusCapturedBy(elem.ID)
		qns = append(qns, qn)
		fmt.Println(len(qns))

	}
	log.WithFields(log.Fields{
		"data": qns[0],
		"length": len(qns),
	}).Info("Backend")
	c.JSON(200, gin.H{"question-set": qns })
	}
}



// Scoreboard API
func Scoreboard(c *gin.Context) {
	// var scoretable []models.Scores
	var str []models.Scores
	str, err := 	models.TopScores()
	if err != nil {
	c.JSON(200, gin.H{"error": err, "scoreboard": nil })
	}else{
		c.JSON(200, gin.H{"scoreboard": str,
			"error": nil })
	// ISSUE % EL
	}
}

// Scoreboard API
func Temp(c *gin.Context) {

 out := execution.Complier("hello.py", "print 'Hello World!'", "python", "", "Hello World!")


	// var scoretable []models.Scores
		c.JSON(200, gin.H{"output":out})
	// ISSUE % EL
}

func AddQuestionAPI(c *gin.Context) {
	var json AddQuestion
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error() , "success": false})
		return
	} else {
			var qn models.Question
			qn.Title = json.title
			qn.Description = json.description
			qn.Score = json.score
			qn.Bonus = json.bonus
			qn.Input1 = json.input1
			qn.Output1 = json.output1
			qn.Input2 = json.input2
			qn.Output2 = json.output2
			err := qn.Save()
			if err != nil{
					c.JSON(200, gin.H{"success": false, "error": err })
			}else{
					c.JSON(200, gin.H{"success": true, "error": nil })
			}
	}
}

func Fixture(c *gin.Context) {
	qn1 := models.Question{Title:"Hello World", Description:"Print Hello World", Score:10, Bonus: 5, Input1:"", Output1:"Hello World!"}
	_ = qn1.Save()
	qn6 := models.Question{Title:"Sumk of two numbers", Description:"Sumk of two numbers", Score:15, Bonus: 5, Input1:"2\n3", Output1:"5", Input2:"4\n6", Output2:"10"}
	_ = qn6.Save()
	qn2 := models.Question{Title:"Difference of two numbers", Description:"Difference of two numbers", Score:15, Bonus: 5, Input1:"2\n3", Output1:"-1", Input2:"6\n4", Output2:"2"}
	_ = qn2.Save()
	qn3 := models.Question{Title:"Square of  numbers", Description:"Square of a number", Score:20, Bonus: 10, Input1:"2", Output1:"4", Input2:"4", Output2:"16"}
	_ = qn3.Save()
	qn4 := models.Question{Title:"Square Root of  numbers", Description:"Square Root of a numbers", Score:10, Bonus: 5, Input1:"4", Output1:"2", Input2:"25", Output2:"5"}
	_ = qn4.Save()
	res1 := models.Result{QnID:1, UserName:"jerin", Score:10, Answer: "print 'Hello World!'", Code:"python", Filename:"hello.py", Status:true}
	_ = res1.Save()
	c.JSON(200, gin.H{"context": "successfully Loaded!" })
}



func Solution(c *gin.Context) {
	var json Results
	if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": err.Error() , "success": false})
	} else {
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println(json.qn_id)
		fmt.Println(json.script)
		fmt.Println(json.filename)
		fmt.Println(json.lang)
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
			var qn models.Question
			qn, _ = models.GetQuestionWithID(json.qn_id)
			fmt.Println(qn)
			// execution.Compile(json.filename, json.script,json.lang, input, output )
			qn1 := models.Result{QnID:json.qn_id, UserName:"", Answer:json.script, Score:0, Status:false,Filename: json.filename, Code: json.lang}
			_ = qn1.Save()

			c.JSON(200, gin.H{"context": "successfully Loaded!" })
	}
}
