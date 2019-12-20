package src

import (
	"fmt"
	"github.com/gin-gonic/gin"
	. "gorestful/src/cors"
	. "gorestful/src/routers"
)


func Run(addr string)  {
	router := gin.Default()
	// 全局处理跨域
	router.Use(HandlerCors())

	// 处理html
	router.LoadHTMLGlob("src/template/*")

	// 路由
	Routers(router)

	// 启动项目
	err := router.Run(addr)
	if err != nil{
		fmt.Println("糟糕，gg了")
	}
}
