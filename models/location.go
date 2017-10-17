package models

import (
	"gopkg.in/mgo.v2/bson"
	"position_mongo/db"
	. "position_mongo/tools"
	"time"
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
	IsDelete   int           `bson:"is_delete" json:"is_delete"`
	ViewNum    int64         `bson:"viewd_num" json:"viewd_num"`
	LikedNum   int64         `bson:"liked_num" json:"liked_num"`
	CommentNum int64         `bson:"comment_num" json:"connent_num"`
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
	err = db.Location.Insert(l)
	PanicError(err)
	return
}

//type AnyLocations struct {
//	Dis      float64 `json:"dis"`
//	Location Location
//}
type AnyLocations struct {
	Ok      int                    `json:"ok"`
	Results []interface{}          `json:"results"`
	Status  map[string]interface{} `json:"status"`
}

func GetNextPageWithLastId(size int, lng float64, lat float64, distance int, id ...bson.ObjectId) (locations AnyLocations, err error) {
	//err = db.Location.Find(bson.M{
	//	"location": bson.M{
	//		"$geoNear": bson.M{
	//			"$geometry": bson.M{
	//				"type":        "Point",
	//				"coordinates": []float64{lng, lat},
	//			},
	//			"$maxDistance": distance,
	//		},
	//	},
	//}).All(&locations)
	//type A interface {
	//}
	//a := new(A)
	err = db.DB.Run(bson.D{
		{"geoNear", "locations"},
		{"near", []float64{lng, lat}},
		{"spherical", true},
	}, &locations)
	return
}

func init() {

}
