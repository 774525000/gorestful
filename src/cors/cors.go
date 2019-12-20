package cors

import "github.com/gin-gonic/gin"

var CorsHeaders = map[string]string{
	"Access-Control-Allow-Origin": "*",
	"Access-Control-Allow-Headers": "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token",
	"Access-Control-Allow-Methods": "POST, GET, OPTIONS",
	"Access-Control-Expose-Headers": "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type",
	"Access-Control-Allow-Credentials": "true",
}

func HandlerCors() gin.HandlerFunc {
	return func(c *gin.Context) {
		for key, value := range CorsHeaders{
			c.Header(key, value)
		}
	}
}