package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"urlLike/db"
	. "urlLike/tools"
)

type User struct {
	Id        bson.ObjectId `bson:"_id"`
	CreateAt  time.Time     `bson:"create_at"`
	UpdatedAt time.Time     `bson:"update_at"`
	DeleteAt  time.Time     `bson:"delete_at"`
	NickName  string        `bson:"nick_name"`
	UserName  string        `bson:"user_name"`
	Age       int           `bson:"age"`
	Pwd       string        `bson:"pwd"`
	Email     string        `bson:"email"`
	Gender    int           `bson:"gender"`
	Summary   string        `bson:"summary"`
	Phone     string        `bson:"phone"`
	IsDelete  bool          `bson:"is_delete"`
	OpenId    string        `bson:"open_id"`
	AvatarUrl string        `bson:"avatar_url"`
}

func AddUser(user *User) {
	user.Id = bson.NewObjectId()
	user.CreateAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsDelete = false
	err := db.Users.Insert(user)
	PanicError(err)
}
