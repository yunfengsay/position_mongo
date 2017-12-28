package main

import (
	"fmt"
	"position_mongo/conf"
	"position_mongo/db"

	"position_mongo/apis"
	"position_mongo/models"

	"github.com/gin-gonic/gin"
)

func AuthNeedLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("session")
		fmt.Println(token)
		if token == "" {
			c.AbortWithStatus(400)
		}
		id := models.GetUserIdByToken(token)
		c.Set("userid", id)
		c.Next()
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", apis.IndexApi)
	user := router.Group("/user")
	user.POST("/add", apis.AddUserApi)
	user.POST("/wx/login", apis.WXLogin)
	location := router.Group("/location")

	location.POST("/add_location", AuthNeedLogin(), apis.AddLocationApi)
	location.POST("/get_locations", apis.GetLocationsApi)
	location.GET("/get_location/:id", apis.GetPageByIdApi)

	router.GET("/get_upload_token", AuthNeedLogin(), apis.GetQiniuTokenApi)
	router.POST("/like/update", AuthNeedLogin(), apis.UpdateLike)
	router.DELETE("/user/delete", AuthNeedLogin(), apis.DeleteUserApi)
	return router
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	MongoSession := db.MongoSession
	defer MongoSession.Close()
	router := initRouter()
	fmt.Println("😄")
	err := router.Run(conf.ConfigContext.ServerPort)
	fmt.Println(err)
}
