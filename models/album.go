package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Album struct {
	Base  `bson:",inline"`
	Owner *User  `bson:"owner" json:"owner"`
	Name  string `bson:"name" json:"name"`
}

func CreateAlbum(v *Album) error {
	sess := MgoSess()
	defer sess.Close()

	v.Id = bson.NewObjectId()
	v.CreatedAt = time.Now()
	v.UpdatedAt = time.Now()

	return sess.DB(DBNAME).C(C_ALBUM_NAME).Insert(v)
}

func GetAlbums(query bson.M) ([]*Album, error) {
	sess := MgoSess()
	defer sess.Close()

	var result []*Album = make([]*Album, 0)
	err := sess.DB(DBNAME).C(C_ALBUM_NAME).Find(query).All(&result)

	return result, err
}
