package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Topic struct {
	Base `bson:",inline"`

	Title   string `bson:"title" json:"title"`
	Content string `bson:"content,omitempty" json:"content"`
	Posts   []Post

	Author *User `bson:"author" json:"author"`
}

func (p *Topic) Validate() error {

	return nil
}

func CreateTopic(p *Topic) error {
	sess := MgoSess()
	defer sess.Close()

	p.Id = bson.NewObjectId()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return sess.DB(DBNAME).C(C_TOPIC_NAME).Insert(p)
}

func GetTopics() ([]*Topic, error) {
	sess := MgoSess()
	defer sess.Close()

	var result []*Topic = make([]*Topic, 0)
	err := sess.DB(DBNAME).C(C_TOPIC_NAME).Find(nil).Sort("-created_at").Limit(10).Select(bson.M{"content": 0}).All(&result)

	return result, err
}

func GetTopicById(id string) (*Topic, error) {
	sess := MgoSess()
	defer sess.Close()

	result := new(Topic)
	err := sess.DB(DBNAME).C(C_TOPIC_NAME).FindId(bson.ObjectIdHex(id)).One(result)
	if err != nil {
		return nil, err
	}
	posts := make([]Post, 0)
	sess.DB(DBNAME).C(C_POST_NAME).Find(bson.M{"topic_id": bson.ObjectIdHex(id)}).All(&posts)
	if err != nil {
		return nil, err
	}
	result.Posts = posts
	return result, nil
}
