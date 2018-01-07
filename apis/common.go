package apis

import (
	"github.com/gin-gonic/gin"
	"position_mongo/models"
	//"position_mongo/db"
	"gopkg.in/mgo.v2/bson"
	//"time"
	"net/http"
)

type AddCommentType struct{
	LocationId bson.ObjectId `bson: "location_id"` // location的id
	To       bson.ObjectId `bson: "to"` // location 的 user id
	From     bson.ObjectId `bson: "from"` // 评论者的id
	UserAvater string `bson: "user_avater"` // 评论者的头像
	UserName string  `bson: "user_name"` // 评论者的昵称
}

func IndexApi(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"message": 1})
}

func AddComment(c *gin.Context){
	comment := &models.Comments{}
	//postData := &AddCommentType{}
	models.AddComments(*comment)
}
