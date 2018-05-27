package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

)

var db  *gorm.DB


//User ...User Model
type CodeUser struct {
	gorm.Model
	Username  string	`gorm:"type:varchar(100);unique_index"`
	Password  string	`gorm:"type:varchar(100)"`
}
func (CodeUser) TableName() string {
	return "codeuser"
}


func (cu CodeUser) Save() (error) {
	err := db.Save(&cu).Error
		return err
}

// func (b Backend) FetchUser(condition interface{}) (user models.User, e error) {
// 	var model models.User
// 	err := b.db.Where(condition).First(&model).Error
// 	return model, err
// }

	// var user CodeUser
func CodeUserGet(condition interface{}) (user CodeUser, e error) {
	db = Model
	// var user CodeUser
	fmt.Println("Going to db")
	// err := db.Where(condition).First(&user).Error
	log.WithFields(log.Fields{
		"db": db,
	}).Info("Backend")


	if err := db.Where(condition).First(&user).Error; err != nil {
		fmt.Println("Error Occured	")
		return user, err
    // error handling...
	}else{
		fmt.Println("Error Not Occured	")
		return user, err
	}


	// log.WithFields(log.Fields{
	// 	"error": err,
	// }).Info("Backend")

	// if err != nil{
	// 	fmt.Println("Error is There...............")
		// 	fmt.Println(err)
	// }else{
	// 	fmt.Println("Error is Not There...............")
	// 	fmt.Println(user)

	// return user, err
}


//Refer https://github.com/demo-apps/go-gin-app
//https://github.com/gothinkster/golang-gin-realworld-example-app/blob/master/users/models.go
//https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
//https://github.com/gin-contrib
//https://github.com/vsouza/go-gin-boilerplate/tree/master/middlewares\
//https://github.com/george518/PPGo_Api_Demo_Gin/tree/master/routers
//https://github.com/gothinkster/golang-gin-realworld-example-app
// swagger & user https://github.com/EDDYCJY/go-gin-example
//https://github.com/night-codes/summer
//https://github.com/szuecs/gin-gomonitor
//https://github.com/sbecker/gin-api-demo
//https://github.com/acrosson/gingorm/tree/master/db
//https://github.com/aviddiviner/gin-limit
//https://github.com/Luncher/go-rest
//https://github.com/nightlegend/apigateway
//https://github.com/rageix/ginAuth
