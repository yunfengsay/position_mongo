package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "position_mongo/tools"
)

func GetQiniuTokenApi(c *gin.Context) {
	token := GetQiniuToken()
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"token": token,
	})
}
