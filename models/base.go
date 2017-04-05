package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Base struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"-"`
}

const (
	DBNAME = "tripper"

	C_CATEGORY_NAME         = "categories"
	C_PROJECT_NAME         = "projects"
	C_PROJECT_COMMENT_NAME = "project_comments"
	C_PROJECT_LIKE_NAME    = "project_likes"
	C_ALBUM_NAME           = "albums"
	C_TOPIC_NAME           = "topics"
	C_POST_NAME            = "posts"
)
