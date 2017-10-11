package models

import (
	"gopkg.in/mgo.v2/bson"
	"position_mongo/db"
	. "position_mongo/tools"
	"time"
)

type Like struct {
	Id       bson.ObjectId `bson: "_id"`
	CreateAt time.Time     `bson: "create_at"`
	Location bson.ObjectId `bson: "location"`
	To       bson.ObjectId `bson: "to"`
	From     bson.ObjectId `bson: "from"`
}

type LikeAction struct {
	AddLike    func(l *Like) (err error)
	DeleteLike func(id bson.ObjectId) (err error)
}

func AddLike(l *Like) (err error) {
	l.CreateAt = time.Now()
	err = db.User.Insert(l)
	PanicError(err)
	return
}

func DeleteLike(id bson.ObjectId) (err error) {
	err = db.Like.Remove(bson.M{"_id": id})
	PanicError(err)
	return
}
