package index

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandlerIndex(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.html", nil)
}
