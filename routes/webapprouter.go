package routes

import (
	c "go-webapp/controller"

	"github.com/gin-gonic/gin"
)

func registerWebAppRouter(router *gin.Engine) {

	webapp := router.Group("/")
	webapp.GET("", c.HomePage)

}


