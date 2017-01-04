package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	Base
	TopicId    bson.ObjectId  `bson:"topic_id" json:"topic_id"`
	Content    string         `bson:"content" json:"content"`
	Comments   []*PostComment `json:"comments"`
	LikeNum    int            `json:"like_num"`
	CommentNum int            `json:"comment_num"`

	Author *User `bson:"user" json:"user"`
}

type PostLike struct {
	Base
	UserId int `json:"user_id`
	PostId int `json:"post_id`
}

type PostComment struct {
	Base
	PostId  int    `json:"post_id"`
	Content string `json:"content"`
}

func (p *Post) Validate() error {
	return nil
}

func CreatePost(p *Post) error {
	sess := MgoSess()
	defer sess.Close()

	p.Id = bson.NewObjectId()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	return sess.DB(DBNAME).C(C_POST_NAME).Insert(p)
}

func GetPostsByTopic(topicId bson.ObjectId) ([]*Post, error) {
	sess := MgoSess()
	defer sess.Close()

	var result []*Post = make([]*Post, 0)
	err := sess.DB(DBNAME).C(C_POST_NAME).Find(bson.M{"topic_id": topicId}).Sort("-created_at").Limit(10).All(&result)

	return result, err
}
