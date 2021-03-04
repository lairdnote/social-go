package api

import (
	"github.com/gin-gonic/gin"
)

// 主页
func Index(c *gin.Context) {
	c.String(200, "================   Welcome social Restful API Index Page!      ================")
}
