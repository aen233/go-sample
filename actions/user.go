package actions

import (
	"github.com/gin-gonic/gin"
	"go-sample/models"
	"math"
	"net/http"
	"strconv"
)

// Index index
func UserIndex(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	size, _ := strconv.Atoi(c.Query("size"))

	data := models.GetUsers(page, size)

	meta := make(map[string]interface{})
	total, _ := models.GetUserTotal()
	meta["total"] = total
	meta["current_page"] = page
	meta["per_page"] = size
	meta["last_page"] = math.Ceil(float64(total / size))

	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"data": data,
		"meta": meta,
	})
}
