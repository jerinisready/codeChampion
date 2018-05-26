package main

import (
	"fmt"
	m "go-webapp/models"
)

func main() {
	fmt.Println("Applying migration")
	// m.Model.AutoMigrate(&m.User{})
	fmt.Println("Applying CodeUser migration")
	m.Model.AutoMigrate(&m.CodeUser{})
	fmt.Println("Applying Question migration")
	m.Model.AutoMigrate(&m.Question{})
	fmt.Println("Applying Result migration")
	m.Model.AutoMigrate(&m.Result{})
	fmt.Println("Finished migration")
}
