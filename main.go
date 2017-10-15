package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"position_mongo/apis"
	"position_mongo/db"
)

func initRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", apis.IndexApi)
	router.GET("/get_upload_token", apis.GetQiniuTokenApi)
	router.POST("/user/add", apis.AddUserApi)
	router.POST("/location/add_location", apis.AddLocationApi)
	router.POST("/location/get_locations", apis.GetLocationsApi)
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
