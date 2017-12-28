package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"position_mongo/models"
	"position_mongo/tools"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	APPID  = "wx4c26683a35d3fab2"
	SECRET = "28397145e103571cfd04867d849501bf"
)

type WXUserLoginForm struct {
	Code     string      `binding:"required" json:"code"`
	UserInfo interface{} `binding:"required" json:"user_info"`
}
type OpenID struct {
	Openid string
}

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
	err, _ := models.AddUser(nUser)
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

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	fmt.Println(json.NewDecoder(r.Body).Decode(target))
	return json.NewDecoder(r.Body).Decode(target)
}
func WXLogin(c *gin.Context) {
	var token string
	wx_login := &WXUserLoginForm{}
	err := c.BindJSON(wx_login)
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=" + APPID + "&secret=" + SECRET + "&js_code=" + wx_login.Code + "&grant_type=authorization_code"
	respId := OpenID{}
	getJson(url, &respId)
	userid := models.FindUserByOpenid(respId.Openid)
	if userid != "" {
		token, err = models.AddOrUpdate(userid, respId.Openid)
	} else {
		nUser := new(models.User)
		user_info := wx_login.UserInfo.(map[string]interface{})
		nUser.Gender = int(user_info["gender"].(float64))
		nUser.NickName = user_info["nickName"].(string)
		nUser.AvatarUrl = user_info["avatarUrl"].(string)
		nUser.OpenId = respId.Openid
		_, id := models.AddUser(nUser)
		token, err = models.AddOrUpdate(id, respId.Openid)
	}
	tools.PanicError(err)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"token": token,
	})
}

func DeleteUserApi(c *gin.Context) {
	id := c.Request.FormValue("id")
	err := models.DeleteUser(id)
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
