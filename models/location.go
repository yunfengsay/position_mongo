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
	Id         bson.ObjectId `bson: "_id"`
	CreateAt   time.Time     `bson: "create_at"`
	UpdatedAt  time.Time     `bson: "update_at"`
	DeleteAt   time.Time     `bson: "delete_at"`
	Imgs       []string      `bson: "imgs"`
	Location   GeoJson       `bson: "location"`
	Content    string        `bson: "content"`
	User       bson.ObjectId `bson: "user"`
	IsDelete   int           `bson: "is_delete"`
	ViewNum    int64         `bson: "viewd_num"`
	LikedNum   int64         `bson: "liked_num"`
	CommentNum int64         `bson: "comment_num"`
}

type LocationAction struct {
	AddLocation    func(l *Location) (err error)
	UpdateLocation func(l *Location) (err error)
	DeleteLocation func(l *Location) (err error)
	GetLocation    func(id bson.ObjectId) (l Location, err error)
	NeerLocation   func()
}

func AddLocation(l *Location) (err error) {
	l.CreateAt = time.Now()
	l.UpdatedAt = time.Now()
	l.LikedNum = 0
	l.ViewNum = 0
	l.CommentNum = 0
	err = db.User.Insert(l)
	PanicError(err)
	return
}
