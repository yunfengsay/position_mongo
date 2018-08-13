package models

import (
	"fmt"
	"position_mongo/db"
	"position_mongo/tools"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Session struct {
	Token    string        `json:"token"`
	Id       bson.ObjectId `json:"id"`
	Openid   string        `json:"openid"`
	Expire   time.Time     `json:"expire"`
	CreateAt time.Time     `json:"create_at"`
}

func AddOrUpdate(id string, openid string) (token string, err error) {
	session := &Session{}
	curent_time := time.Now()
	count, _ := db.Session.Find(bson.M{"id": bson.ObjectIdHex(id)}).Count()
	token = tools.CreateHashWithSalt(id)
	if count == 0 {
		session.CreateAt = curent_time
		session.Id = bson.ObjectIdHex(id)
		if openid != "none" {
			session.Openid = openid
		}
		session.Token = token
		err = db.Session.Insert(session)
	} else {
		err = db.Session.Update(bson.M{"id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"token": token}})
	}
	return
}

func FindUserByOpenid(openid string) (id string) {
	var result []User
	fmt.Println(openid)
	err := db.User.Find(bson.M{"openid": openid}).All(&result)
	if len(result) != 0 {
		return result[0].Id.Hex()
	}
	tools.PanicError(err)
	return ""
}

func GetUserIdByToken(token string) (id string) {
	result := Session{}

	db.Session.Find(bson.M{"token": token}).One(&result)
	id = result.Id.Hex()
	return
}
