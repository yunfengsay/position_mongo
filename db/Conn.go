package db

import (
	"gopkg.in/mgo.v2"
)

const (
	MONGO_URL = "127.0.0.1:27017"
)

var Location *mgo.Collection
var User *mgo.Collection
var Like *mgo.Collection
var Comment *mgo.Collection
var MongoSession *mgo.Session

func init() {
	MongoSession, _ = mgo.Dial(MONGO_URL)
	//切换到数据库
	db := MongoSession.DB("position")
	//切换到collection
	User = db.C("users")
	Location = db.C("locations")
	//Location.EnsureIndex(mgo.Index{Name: "location", Key: []string{"$2dsphere:location"}})
	Like = db.C("like")
	Comment = db.C("comments")
	//Urls = db.C("urls")
}
