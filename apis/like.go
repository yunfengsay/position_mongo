package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateLike(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
