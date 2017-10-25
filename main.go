package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"position_mongo/apis"
	"position_mongo/db"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "clinet_request")
		c.Next()
		fmt.Println("after middleware")
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()
	router.Use(MiddleWare())
	router.GET("/", apis.IndexApi)

	user := router.Group("/user")
	user.POST("/add", apis.AddUserApi)
	user.POST("/wx/login", apis.WXLogin)

	location := router.Group("/location")

	location.POST("/add_location", apis.AddLocationApi)
	location.POST("/get_locations", apis.GetLocationsApi)
	location.GET("/get_location/:id", apis.GetPageByIdApi)

	router.GET("/get_upload_token", apis.GetQiniuTokenApi)
	router.POST("/like/update", apis.UpdateLike)
	router.DELETE("/user/delete", apis.DeleteUserApi)

	return router
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	MongoSession := db.MongoSession
	defer MongoSession.Close()
	router := initRouter()
	fmt.Println("ok")

	router.Run(":8002")
}
