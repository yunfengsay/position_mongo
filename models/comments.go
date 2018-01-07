package models
import (
	"gopkg.in/mgo.v2/bson"
	"position_mongo/db"
	. "position_mongo/tools"
	"time"
)

type Comments struct {
	Id       bson.ObjectId `bson: "_id"`
	CreateAt time.Time     `bson: "create_at"`
	LocationId bson.ObjectId `bson: "location_id"` // location的id
	To       bson.ObjectId `bson: "to"` // location 的 user id
	From     bson.ObjectId `bson: "from"` // 评论者的id
	UserAvater string `bson: "user_avater"` // 评论者的头像
	UserName string  `bson: "user_name"` // 评论者的昵称
}

func AddComments(comment Comments)(err error){
	comment.CreateAt  = time.Now()
	err = db.Comment.Insert(comment)
	if err !=nil{
		PanicError(err)
	}
	return
}