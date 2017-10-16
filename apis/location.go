package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"position_mongo/models"
	"position_mongo/tools"
)

type AddLocationApiForm struct {
	Imgs    []string  `binding:"required"`
	Point   []float64 `binding:"required"`
	Content string
	L_type  []int `binding:"required" json:"l_type"`
}

func AddLocationApi(c *gin.Context) {
	location := new(models.Location)

	data := new(AddLocationApiForm)
	e := c.BindJSON(data)
	if e != nil {
		c.AbortWithError(400, e)
		return
	}
	fmt.Println(data.Point)
	//c.String(200, fmt.Sprintf("%#v", data))
	location.Content = data.Content
	location.LType = data.L_type
	location.Imgs = data.Imgs
	lng_lat := &models.GeoJson{
		Type:        "Point",
		Coordinates: data.Point,
	}
	location.Location = *lng_lat

	location.Id = bson.NewObjectId()
	location.User = bson.NewObjectId()

	err := models.AddLocation(location)
	tools.PanicError(err)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}

type GetLocationApiForm struct {
	Point  []float64 `binding:"required" json:"point"`
	L_type []int     `binding:"required" json:"l_type"`
	R      int       `binding:"required" json:"r"`
}

func GetLocationsApi(c *gin.Context) {
	data := &GetLocationApiForm{}
	e := c.BindJSON(data)
	if e != nil {
		fmt.Println(e)
		c.AbortWithError(400, e)
		return
	}
	fmt.Println(data)
	locations, err := models.GetNextPageWithLastId(10, data.Point[0], data.Point[1], data.R)
	tools.PanicError(err)
	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"locations": locations,
	})
}
