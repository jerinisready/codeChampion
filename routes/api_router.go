package routes

import (
	c "go-webapp/controller"
	// "go-webapp/controller/auth"
	"go-webapp/middleware/session"
	// "go-webapp/module/server"
	"github.com/gin-gonic/gin"
)

func registerAPIRouter(router *gin.Engine) {
	router.Use(session.SessionMiddleWare())
	// api := router.Group("/api")
	// api.GET("/index", c.IndexApi)
	//
	// v1 := router.Group("api/v1")
	// {
	// 	v1.POST("/register", auth.Register)
	// 	v1.POST("/login", auth.UserLogin)
	// }
	//
  // // DEBUGGER
	// router.GET("/version", server.Version)
	api := router.Group("/")
	api.POST("/i-am-me/", c.LoginAPI)
	api.GET("/question-set/", c.QuestionSet)
	api.GET("/scoreboard/", c.Scoreboard)
	api.GET("/fixture/", c.Fixture)
	api.POST("/add-question/", c.AddQuestionAPI)
	api.POST("/compile/", c.Solution)
	// api.GET("/scoreboard/", c.Scoreboard)
}



// debugger := router.Group("/api/debug")
// {
// 	//TODO Session Must Admin
// 	debugger.GET("/pprof/", debug.IndexHandler())
// 	debugger.GET("/pprof/heap", debug.HeapHandler())
// 	debugger.GET("/pprof/goroutine", debug.GoroutineHandler())
// 	debugger.GET("/pprof/block", debug.BlockHandler())
// 	debugger.GET("/pprof/threadcreate", debug.ThreadCreateHandler())
// 	debugger.GET("/pprof/cmdline", debug.CmdlineHandler())
// 	debugger.GET("/pprof/profile", debug.ProfileHandler())
// 	debugger.GET("/pprof/symbol", debug.SymbolHandler())
// 	debugger.POST("/pprof/symbol", debug.SymbolHandler())
// 	debugger.GET("/pprof/trace", debug.TraceHandler())
// }
