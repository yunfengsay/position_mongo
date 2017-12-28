package db

import (
	"position_mongo/conf"

	"gopkg.in/mgo.v2"
)

var Location *mgo.Collection
var User *mgo.Collection
var Like *mgo.Collection
var Comment *mgo.Collection
var Session *mgo.Collection

var MongoSession *mgo.Session
var DB *mgo.Database

func init() {
	diaInfo := &mgo.DialInfo{
		Addrs:    []string{conf.ConfigContext.DBUrl},
		Username: conf.ConfigContext.DBUser,
		Password: conf.ConfigContext.DBPwd,
	}
	MongoSession, _ = mgo.DialWithInfo(diaInfo)
	// fmt.Println(conf.ConfigContext.DBUrl)
	// MongoSession, _ = mgo.Dial(conf.ConfigContext.DBUrl)
	// if err != nil {
	// 	fmt.Println("😢 数据库连接失败 ", err)
	// }
	// 选择数据库
	DB = MongoSession.DB(conf.ConfigContext.DBName)

	User = DB.C("users")
	Location = DB.C("locations")
	//Location.EnsureIndex(mgo.Index{Name: "location", Key: []string{"$2dsphere:location"}})
	Like = DB.C("like")
	Session = DB.C("session")
	Comment = DB.C("comments")

}

// 要在 mongodb中切换到 position 运行
//db.locations.ensureIndex({"location":"2dsphere"})
