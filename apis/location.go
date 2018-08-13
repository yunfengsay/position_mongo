package apis

import (
	"fmt"
	"net/http"
	"position_mongo/models"
	"position_mongo/tools"

	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
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
	token := c.Request.Header.Get("token")
	user_id := models.GetUserIdByToken(token)
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
	location.User = bson.ObjectIdHex(user_id)
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

func GetPageByIdApi(c *gin.Context) {
	id := c.Param("id")
	location, err := models.GetPageById(id)
	tools.PanicError(err)
	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"location": location,
	})
}

func GetLocationsApi(c *gin.Context) {
	data := &GetLocationApiForm{}
	token := c.Request.Header.Get("token")
	fmt.Println(token)
	user_id := models.GetUserIdByToken(token)

	e := c.BindJSON(data)
	if e != nil {
		fmt.Println(e)
		c.AbortWithError(400, e)
		return
	}
	if user_id == "" {
		user_id = "not_in"
	}
	locations, err := models.GetNextPageWithLastId(user_id, 10, data.Point[0], data.Point[1], data.R)

	for _, location := range locations {
		l := location.(bson.M)
		if l["user"].(bson.ObjectId).Hex() == user_id {
			l["is_self"] = true
			continue
		}
		fmt.Println(l["user"].(bson.ObjectId).Hex(), user_id)
		l["is_self"] = false
	}

	tools.PanicError(err)
	c.JSON(http.StatusOK, gin.H{
		"code":      0,
		"locations": locations,
	})
}
func DeleteLocation(c *gin.Context) {
	token := c.Request.Header.Get("token")
	user_id := models.GetUserIdByToken(token)
	location_id := c.Param("id")
	err := models.DeleteLocation(location_id, user_id)
	if err != nil {
		c.AbortWithError(400, err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
	})
}
