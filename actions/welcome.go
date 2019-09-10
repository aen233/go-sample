package actions

import (
	"github.com/gin-gonic/gin"
	"go-sample/pkg"
	"net/http"
)

// Index index
func Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "page/welcome.html", gin.H{
		"data":    "Main website",
		"session": pkg.GetUserSession(c),
	})
}
