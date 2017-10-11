package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddLocationApi(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
