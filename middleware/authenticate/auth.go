package authentication

import (
	"github.com/gin-gonic/gin"
)


func GetSession(c *gin.Context) {
   values := c.Request.Header["Authentication"]
   if len(values) > 0 {
      c.Set("username",  values[0])
    }
}

func SetSession(c *gin.Context, username string){
    c.Set("username",  username)
}
