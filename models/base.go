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

	C_PROJECT_NAME = "projects"
	C_ALBUM_NAME   = "albums"
)
