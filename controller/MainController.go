package controllers

import (
	"errors"
	"fmt"
	"go-webapp/models"
	"go-webapp/compileunit"
	"go-webapp/common"
	"go-webapp/middleware/header"
	"github.com/gin-gonic/gin"
	"go-webapp/middleware/authenticate"
	log "github.com/sirupsen/logrus"

)

type Login struct {
	Username     string `form:"username" json:"username" binding:"required"`
	Password 		 string `form:"password" json:"password" binding:"required"`
}


// Login API
func LoginAPI(c *gin.Context) {

	var json Login
	fmt.Println("")
	header.Secure(c)
	header.Options(c)
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
	header.Secure(c)
	header.Options(c)

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
	c.JSON(200, gin.H{"questionset": qns[:3],
		 "error": nil })
	}
}


// Scoreboard API
func Scoreboard(c *gin.Context) {
	header.Secure(c)
	header.Options(c)

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

 // out := execution.Complier("hello.py", "print 'Hello World!'", "python", "", "Hello World!")
	out := ""

	// var scoretable []models.Scores
		c.JSON(200, gin.H{"output":out})
	// ISSUE % EL
}


type AddQuestion struct {
	Title     string `form:"Title" json:"Title" binding:"required"`
	Description     string `form:"Description" json:"Description" binding:"required"`
	Score     string `form:"Score" json:"Score" `
	Bonus     string `form:"Bonus" json:"Bonus"`
	Input1     string `form:"Input1" json:"Input1" `
	Output1    string `form:"Output1" json:"Output1" binding:"required"`
	Input2     string `form:"Input2" json:"Input2" `
	Output2    string `form:"Output2" json:"Output2" `
}


func AddQuestionAPI(c *gin.Context) {
	header.Secure(c)
	header.Options(c)

	var json AddQuestion
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": err.Error() , "success": false})
		return
	} else {
			var qn models.Question
			qn.Title = json.Title
			qn.Description = json.Description
			qn.Score = 10
			qn.Bonus = 5
			qn.Input1 = json.Input1
			qn.Output1 = json.Output1
			qn.Input2 = json.Input2
			qn.Output2 = json.Output2
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


type Results struct {
	Username string `form:"username" json:"username" binding:"required"`
	QnId    int	`form:"qn_id" json:"qn_id" binding:"required"`
	Script	 string	`form:"script" json:"script" binding:"required"`
	Filename string	`form:"filename" json:"filename" binding:"required"`
	Lang		 string	`form:"lang" json:"lang" binding:"required"`
}



func Solution(c *gin.Context) {
	header.Secure(c)
	header.Options(c)

	var json Results
	if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(400, gin.H{"error": err.Error() , "success": false})
	} else {
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		fmt.Println(json)
		fmt.Println(json.QnId)
		fmt.Println(json.Script)
		fmt.Println(json.Filename)
		fmt.Println(json.Lang)
		fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
		var qn models.Question
		qn, _ = models.GetQuestionWithID(json.QnId)
		fmt.Println(qn)
		out := execution.Complier(json.Filename, json.Script, json.Lang, qn.Input1, qn.Output1, qn.Input2, qn.Output2, )
		// out := execution.Complier("hello.py", "print 'Hello World!'", "python", "", "Hello World!", "-", "-" )
		// var scoretable []models.Scores
		var score int
		score = 0
		if out == "" {
				score = qn.Score
				if models.BonusEligible(json.QnId){
						score = score + qn.Bonus
				}else{

				}
		}
		var message string
		success := bool(out =="")
		if success {
			message = "You Have completed this Level! Go For Next Program"
		}else{
			message = "Please Correct Errors and try again"
		}
		c.JSON(200, gin.H{"success": success, "message":message, "output":out, "errors": ""})
		// execution.Compile(json.filename, json.script,json.lang, input, output )
		qn1 := models.Result{QnID:json.QnId, UserName:json.Username, Answer:json.Script, Score:score, Status:false,Filename: json.Filename, Code: json.Lang}
		_ = qn1.Save()
	}
}
