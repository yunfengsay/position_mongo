package models

import "gopkg.in/mgo.v2/bson"

type Urls struct {
	Id         bson.ObjectId   `bson:"_id"`
	Title      string          `bson:"title"`
	Summary    string          `bson:"summary"`
	Url        []bson.ObjectId `bson:"url"`
	PersionNum int64           `bson:"persion_num"`
	Score      float32         `bson:"score"`
}
