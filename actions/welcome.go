package actions

import (
	"net/http"
	"sample/pkg"

	"github.com/gin-gonic/gin"
)

// Index index
func Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "page/welcome.html", gin.H{
		"data":    "Main website",
		"session": pkg.GetUserSession(c),
	})
}
