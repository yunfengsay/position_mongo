package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"position_mongo/tools"
	"fmt"
	"position_mongo/models"
)

type UpdateLikePostBody struct {
	Target string
	Type   string
	TargetUser string `json:"target_user"`
}

func UpdateLike(c *gin.Context) {
	data := &UpdateLikePostBody{}
	e := c.BindJSON(&data)
	fmt.Println(data)
	tools.PanicError(e)
	user,_:= c.Get("userid")
	models.AddOrDeleteLike(data.Target,data.TargetUser,user.(string),data.Type)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
