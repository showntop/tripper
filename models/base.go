package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type BaseModel struct {
	Id        bson.ObjectId `bson:"_id,omitempty" json:"id" jsonapi:"primary,id"` // 必须要设置bson:"_id" 不然mgo不会认为是主键
	CreatedAt time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at" json:"-"`
	DeletedAt *time.Time    `bson:"deleted_at" json:"-"`
}

func (m BaseModel) NewRecord() bool {
	return m.Id == ""
}

func (m BaseModel) IsDeleted() bool {
	return m.DeletedAt != nil
}
