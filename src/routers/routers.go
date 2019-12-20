package routers

import (
	"github.com/gin-gonic/gin"
	. "gorestful/src/routers/index"
	. "gorestful/src/routers/login"
)

func Routers (router *gin.Engine)  {
	router.GET("/", HandlerIndex)
	router.GET("/api/login", HandlerLogin)
}
