package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
	"urlLike/db"
	. "urlLike/tools"
)

type Industry struct {
	Id         bson.ObjectId   `bson:"_id"`
	Title      string          `bson:"title"`
	Summary    string          `bson:"summary"`
	CreateAt   time.Time       `bson:"create_at"`
	Urls       []bson.ObjectId `bson:"urls"`
	PersionNum int64           `bson:"persion_num"`
	Score      float32         `bson:"score"`
}

func AddIndustry(industry *Industry) {
	iIndustry := new(Industry)
	iIndustry.Id = bson.NewObjectId()
	iIndustry.Title = industry.Title
	iIndustry.Summary = industry.Summary
	iIndustry.PersionNum = 0
	iIndustry.CreateAt = time.Now()
	iIndustry.Score = 0
	err := db.Industry.Insert(iIndustry)
	PanicError(err)
}
