package apis

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"urlLike/models"
)

func AddUserApi(c *gin.Context) {
	pwd := c.Request.FormValue("pwd")
	username := c.Request.FormValue("username")
	nUser := new(models.User)
	nUser.Id = bson.NewObjectId()
	nUser.Username = username
	nUser.Pwd = pwd
	models.AddUser(nUser)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
