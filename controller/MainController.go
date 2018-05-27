package controllers

import (
	"errors"
	"fmt"
	"go-webapp/models"
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
	title     string `form:"title" json:"title" binding:"required"`
	description     string `form:"description" json:"description" binding:"required"`
	score     string `form:"description" json:"description" binding:"required"`
	bonus     string `form:"bonus" json:"bonus" binding:"required"`
	input     string `form:"input" json:"input" binding:"required"`
	output    string `form:"output" json:"output" binding:"required"`
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
	c.JSON(200, gin.H{"scoreboard": err })
	}else{
		c.JSON(200, gin.H{"scoreboard": str })
	// ISSUE % EL
	}
}


