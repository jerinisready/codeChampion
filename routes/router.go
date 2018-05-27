package routes

import (
	"time"
	"go-webapp/config"
	"go-webapp/handle"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)
//InitRouter Initialise router
func InitRouter() *gin.Engine {
	route := gin.New()

	route.Use(cors.New(cors.Config{
					AllowOrigins:     []string{"*"},
					AllowMethods:     []string{"POST", "GET"},
					AllowHeaders:     []string{"Origin", "Authentication", "Content-Type"},
					ExposeHeaders:    []string{"Content-Length"},
					AllowCredentials: true,
					AllowOriginFunc: func(origin string) bool {
						return origin == "*"
					},
					MaxAge: 12 * time.Hour,
				}))
		defConfig := cors.DefaultConfig()
		defConfig.AllowOrigins = []string{"*"}
		// config.AllowOrigins == []string{"http://google.com", "http://facebook.com"}

		route.Use(cors.New(defConfig))

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if config.GetEnv().DEBUG {
		route.Use(gin.Logger()) // Used in development mode, console print request records
	}
	route.Use(handle.Errors()) // Error handling
	registerAPIRouter(route)
	route.LoadHTMLGlob("templates/*")
	route.Static("/assets", "./assets")
    registerWebAppRouter(route)


	return route
}
