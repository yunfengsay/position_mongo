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
	router.POST("/user/add", apis.AddUserApi)
	//router.GET("/addlike", apis.AddLike)
	return router
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	MongoSession := db.MongoSession
	fmt.Println("ok")
	defer MongoSession.Close()
	router := initRouter()
	router.Run(":8002")
}
