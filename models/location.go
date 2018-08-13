package models

import (
	"fmt"
	"math"
	"position_mongo/db"
	. "position_mongo/tools"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type GeoJson struct {
	Type        string    `json:"-"`
	Coordinates []float64 `json:"coordinates"`
}
type Location struct {
	Id         bson.ObjectId `bson:"_id" json:"id"`
	LType      []int         `bson:"l_type" json:"l_type"`
	CreateAt   time.Time     `bson:"create_at" json:"create_at"`
	UpdatedAt  time.Time     `bson:"update_at" json:"update_at"`
	DeleteAt   time.Time     `bson:"delete_at" json:"delete_at"`
	Imgs       []string      `bson:"imgs" json:"imgs"`
	Location   GeoJson       `bson:"location" json:"location"`
	Content    string        `bson:"content" json:"content"`
	User       bson.ObjectId `bson:"user" json:"user"`
	UserObj    User          `bson:"user_obj" json:"user_obj"`
	IsDelete   int           `bson:"is_delete" json:"is_delete"`
	ViewNum    int64         `bson:"viewd_num" json:"viewd_num"`
	LikedNum   int64         `bson:"liked_num" json:"liked_num"`
	CommentNum int64         `bson:"comment_num" json:"connent_num"`
	Liked      []string      `bson:"liked" json:"liked"`
}

type LocationAction struct {
	AddLocation    func(l *Location) (err error)
	UpdateLocation func(l *Location) (err error)
	DeleteLocation func(l *Location) (err error)
	GetLocation    func(id bson.ObjectId) (l Location, err error)
	GetLocations   func(id bson.ObjectId) (l Location, err error)
	NeerLocation   func()
}

func AddLocation(l *Location) (err error) {
	l.CreateAt = time.Now()
	l.UpdatedAt = time.Now()
	l.LikedNum = 0
	l.ViewNum = 0
	l.CommentNum = 0
	l.Liked = []string{}
	user_obj := User{}
	db.User.Find(bson.M{"_id": l.User}).One(&user_obj)
	user_obj.OpenId = ""
	l.UserObj = user_obj
	err = db.Location.Insert(l)
	PanicError(err)
	return
}

// 接受 mongodb 返回的结果
type AnyLocations struct {
	Ok      int                    `json:"ok"`
	Results []interface{}          `json:"results"`
	Status  map[string]interface{} `json:"status"`
}

func GetPageById(id string) (location interface{}, err error) {
	db.Location.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&location)
	fmt.Println(location)
	return
}

func GetNextPageWithLastId(userid string, size int, lng float64, lat float64, distance int, id ...string) (locations []interface{}, err error) {
	var db_results AnyLocations
	if len(id) != 0 {
		err = db.DB.Run(bson.D{
			{"geoNear", "locations"},
			{"near", []float64{lng, lat}},
			{"spherical", true},
			{"maxDistance", distance},
			{"distanceMultiplier", 6371}, // spherical 的值为true 结果如果是 km 这个要设为 6371
			{"limit", size},
			//{"query", bson.D{{"_id", bson.D{{"$gt", bson.ObjectIdHex(id[0])}}}}},
			{"query", bson.M{
				"_id":       bson.M{"$gt": bson.ObjectIdHex(id[0])},
				"is_delete": 0,
			}},
		}, &db_results)

	} else {
		err = db.DB.Run(bson.D{
			{"geoNear", "locations"},
			{"near", []float64{lng, lat}},
			{"spherical", true},
			{"maxDistance", distance},
			{"distanceMultiplier", 6371},
			{"limit", size},
			{"query", bson.M{
				"is_delete": 0,
			}},
		}, &db_results)
	}

	if err == nil {
		for _, v := range db_results.Results {
			s, _ := v.(bson.M)
			obj, _ := s["obj"].(bson.M)
			//s["dis"] = fmt.Sprintf("%.2f", s["dis"])
			s["dis"] = math.Trunc(s["dis"].(float64)*1e2+0.5) * 1e-2 //保留两位小数
			obj["dis"] = s["dis"]
			obj["id"] = obj["_id"]
			delete(s, "dis")
			obj["is_liked"] = false
			if userid != "not_in" {
				for _, m := range obj["liked"].([]interface{}) {
					if m == userid {
						obj["is_liked"] = true
					}
				}
			}
			//v = obj
			locations = append(locations, obj)
		}
		if locations == nil {
			fmt.Println("🈳️ ", locations)
		}
	} else {
		fmt.Println(err)
	}
	return
}

func DeleteLocation(location_id, user_id string) (err error) {
	err = db.Location.Update(bson.M{"_id": bson.ObjectIdHex(location_id), "user": bson.ObjectIdHex(user_id)}, bson.M{"$set": bson.M{"is_delete": 1}})
	return
}
func init() {

}
