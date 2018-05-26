package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"

)

func IndexApi(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
// func Handler(c *gin.Context) {
//
// 	c.HTML(http.StatusOK, gin.H{
// 		"message": "pong",
// 	})
// }

type Login struct {
	Username     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type AddQuestion struct {
	title     string `form:"title" json:"title" binding:"required"`
	description     string `form:"description" json:"description" binding:"required"`
	score     string `form:"description" json:"description" binding:"required"`
	bonus     string `form:"bonus" json:"bonus" binding:"required"`
	input     string `form:"input" json:"input" binding:"required"`
	output    string `form:"output" json:"output" binding:"required"`
}

type Result struct {
	qn_id    int	`form:"qn_id" json:"qn_id" binding:"required"`
	script	 string	`form:"script" json:"script" binding:"required"`
	filename string	`form:"filename" json:"filename" binding:"required"`
	lang		 string	`form:"lang" json:"lang" binding:"required"`
}

func LoginAPI(c *gin.Context) {
	var json Login
	if err := c.ShouldBindJSON(&json); err != nil {
		// if json.username == "manu" && json.Password == "123" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "success": false})
		// } else {
			// c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		// }
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "you are about to logged in"})
	}
	status := true
	c.JSON(http.StatusOK, gin.H{
		"success": status,
		"error": "",
	})
}
