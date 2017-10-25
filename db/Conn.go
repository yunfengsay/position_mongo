package db

import (
	"gopkg.in/mgo.v2"
)

const (
	MONGO_URL = "123.206.116.94:27017"
)

var Location *mgo.Collection
var User *mgo.Collection
var Like *mgo.Collection
var Comment *mgo.Collection
var Session *mgo.Collection

var MongoSession *mgo.Session
var DB *mgo.Database

func init() {
	MongoSession, _ = mgo.Dial(MONGO_URL)
	//切换到数据库
	DB = MongoSession.DB("position")
	//切换到collection
	User = DB.C("users")
	Location = DB.C("locations")
	//Location.EnsureIndex(mgo.Index{Name: "location", Key: []string{"$2dsphere:location"}})
	Like = DB.C("like")
	Session = DB.C("session")
	Comment = DB.C("comments")
	//Urls = db.C("urls")
}
