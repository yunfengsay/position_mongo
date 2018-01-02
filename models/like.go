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
	Location bson.ObjectId `bson: "location"` // location的id
	To       bson.ObjectId `bson: "to"` // location 的 user id
	From     bson.ObjectId `bson: "from"` // 点击者的id
}

type LikeAction struct {
	AddLike    func(l *Like) (err error)
	DeleteLike func(id bson.ObjectId) (err error)
}

func AddOrDeleteLike(location_id, to_id, from_id,like_type string) (err error) {
	l := Like{}
	if like_type == "like"{
		l.CreateAt = time.Now()
		l.Location = bson.ObjectIdHex(location_id)
		l.To = bson.ObjectIdHex(to_id)
		l.From = bson.ObjectIdHex(from_id)
		l.Id = bson.NewObjectId()
		go db.Like.Insert(l)
		err =db.Location.Update(bson.M{"_id": bson.ObjectIdHex(location_id)},bson.M{"$inc": bson.M{"liked_num" :1},"$addToSet":bson.M{"liked":from_id}})
		PanicError(err)

	} else {
		go db.Like.Remove(bson.M{"from":bson.ObjectIdHex(from_id),"location":bson.ObjectIdHex(location_id)})
		db.Location.Update(bson.M{"_id": bson.ObjectIdHex(location_id)}, bson.M{"$inc": bson.M{"liked_num" :-1},"$pull":bson.M{"liked":from_id}})
	}
	PanicError(err)
	return
}
