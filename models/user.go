package models

import (
	"gopkg.in/mgo.v2/bson"
	"position_mongo/db"
	. "position_mongo/tools"
	"time"
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
	IsDelete  int           `bson:"is_delete"`
	OpenId    string        `bson:"open_id"`
	AvatarUrl string        `bson:"avatar_url"`
}

type UserAction struct {
	AddUser    func(user *User) (err error)
	UpdateUser func(user *User) (err error)
	DeleteUser func(id bson.ObjectId) (err error)
	GetUser    func(id bson.ObjectId) (user User, err error)
}

func AddUser(user *User) (err error) {
	user.CreateAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsDelete = 0
	err = db.User.Insert(user)
	PanicError(err)
	return
}

func UpdateUser(user *User) (err error) {
	return
}

func GetUser(id bson.ObjectId) (err error) {
	return
}

func DeleteUser(id bson.ObjectId) (err error) {
	err = db.User.Update(bson.M{"_id": id}, bson.M{"$set": bson.M{"is_delete": 1}})
	return
}
