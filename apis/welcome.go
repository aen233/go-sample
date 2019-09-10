package apis

import (
	"net/http"
	"sample/pkg"

	"github.com/gin-gonic/gin"
)

// Index index
func Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"data":    "hello-world",
		"session": pkg.GetUserSession(c),
	})
}
