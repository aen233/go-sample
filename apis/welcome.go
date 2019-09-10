package apis

import (
	"github.com/gin-gonic/gin"
	"go-sample/pkg"
	"net/http"
)

// Index index
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"data":    "hello-world",
		"session": pkg.GetUserSession(c),
	})
}
