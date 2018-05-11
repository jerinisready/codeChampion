package handle

import (
	"fmt"
	"go-webapp/config"
	"io"
	"os"
	"runtime/debug"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	defaultWriter io.Writer
)

// Errors : Handler to handle errors
func Errors() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {

				file, openErr := os.OpenFile(config.GetEnv().ERROR_LOG_PATH, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
				if openErr == nil {
					if config.GetEnv().DEBUG {
						defaultWriter = io.MultiWriter(file, os.Stdout)
					} else {
						defaultWriter = io.MultiWriter(file)
					}

					fmt.Fprintf(defaultWriter, "%s", "\n")
					//fmt.Fprintf(defaultWriter, "%s %3d %s", red, "Error Msg: ", reset)
					fmt.Fprintf(defaultWriter, "%s", "["+time.Now().Format("2006-01-02 15:04:05")+"] app.ERROR: ")
					fmt.Fprintf(defaultWriter, "%s", err)
					fmt.Fprintf(defaultWriter, "%s", "\nStack trace:\n")
					fmt.Fprintf(defaultWriter, "%s", debug.Stack())
					fmt.Fprintf(defaultWriter, "%s", "\n")
				}

				c.JSON(200, gin.H{
					"code": 10500,
					"msg":  err,
				})
			}
		}()
		c.Next()
	}
}
