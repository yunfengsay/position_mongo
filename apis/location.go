package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"position_mongo/models"
	"position_mongo/tools"
)

type Imgs struct {
	Imgs    []string  `binding:"required"`
	Point   []float64 `binding: "reauired"`
	Content string    `binding: "reauired"`
	L_type  []int     `binding:"required";json:"l_type"`
}

func AddLocationApi(c *gin.Context) {
	location := new(models.Location)

	data := new(Imgs)
	e := c.BindJSON(data)
	if e != nil {
		c.AbortWithError(400, e)
		return
	}
	fmt.Println(data.L_type)
	//c.String(200, fmt.Sprintf("%#v", data))
	location.Content = data.Content
	location.LType = data.L_type
	location.Imgs = data.Imgs
	lng_lat := &models.GeoJson{
		Type:        "point",
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

func GetLocationsApi(c *gin.Context) {

}
