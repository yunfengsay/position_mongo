package apis

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"position_mongo/models"
	"position_mongo/tools"
	"strconv"
)

const (
	APPID  = "wx4c26683a35d3fab2"
	SECRET = "28397145e103571cfd04867d849501bf"
)

func AddUserApi(c *gin.Context) {
	pwd := c.Request.FormValue("pwd")
	nUser := new(models.User)
	nUser.UserName = c.Request.FormValue("user_name")
	nUser.NickName = c.Request.FormValue("nick_name")
	nUser.Age, _ = strconv.Atoi(c.Request.FormValue("age"))
	nUser.Pwd = tools.CreateHashWithSalt(c.Request.FormValue("pwd"))
	nUser.Email = c.Request.FormValue("email")
	nUser.Gender, _ = strconv.Atoi(c.Request.FormValue("gender"))
	nUser.Summary = c.Request.FormValue("summary")
	nUser.Phone = c.Request.FormValue("phone")
	nUser.AvatarUrl = c.Request.FormValue("avatarurl")
	nUser.Pwd = pwd
	err := models.AddUser(nUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
		})
	}
}

func WXLogin(c *gin.Context) {

}

func DeleteUserApi(c *gin.Context) {
	id := c.Request.FormValue("id")
	err := models.DeleteUser(bson.ObjectIdHex(id))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
		})
	}
}
