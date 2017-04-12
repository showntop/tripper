package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Feed struct {
	Base `bson:",inline"`

	Title string `bson:"title" json:"title"`
	Asset string `bson:"asset" json:"asset"`
	Intro string `bson:"intro" json:"intro"`
	Media string `bson:"media" json:"media"`

	Album struct {
		Id   string `bson:"id" json:"id"`
		Name string `bson:"name" json:"name"`
	} `bson:"album" json:"album"`

	Author *User `bson:"author" json:"author"`
}

func ListFeeds(skip, limit int) ([]*Project, error) {
	sess := MgoSess()
	defer sess.Close()

	var result []*Project = make([]*Project, 0)
	err := sess.DB(DBNAME).C(C_PROJECT_NAME).Find(nil).Sort("-created_at").Skip(skip).Limit(limit).Select(bson.M{"content": 0}).All(&result)

	return result, err
}
