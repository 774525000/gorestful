package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerLogin(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"name": "shuaige",
		"age": 18,
	})
}
