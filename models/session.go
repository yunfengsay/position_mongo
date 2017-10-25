package models

import (
	"crypto/md5"
	"gopkg.in/mgo.v2/bson"
	"position_mongo/db"
	"position_mongo/tools"
	"time"
)

type Session struct {
	Token    string        `json:"token"`
	Id       bson.ObjectId `json:"id"`
	Expire   time.Time     `json:"expire"`
	CreateAt time.Time     `json:"create_at"`
}

func (session *Session) Add(id string) (token string, err error) {
	tools.PanicError(err)
	curent_time := time.Now()
	session.CreateAt = curent_time
	token = tools.CreateHashWithSalt(id)
	err = db.User.Insert(session)
	return
}
