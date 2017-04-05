package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Album struct {
	Base     `bson:",inline"`
	Owner    *User     `bson:"owner" json:"owner"`
	Name     string    `bson:"name" json:"name"`
	Projects []Project `bson:"-" json:"projects"`
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
	err := sess.DB(DBNAME).C(C_ALBUM_NAME).Find(query).Limit(10).All(&result)

	return result, err
}

func GetAlbumById(id bson.ObjectId) (*Album, error) {
	sess := MgoSess()
	defer sess.Close()

	var result Album
	err := sess.DB(DBNAME).C(C_ALBUM_NAME).FindId(id).One(&result)
	err = sess.DB(DBNAME).C(C_PROJECT_NAME).Find(bson.M{"album.id": id.Hex()}).All(&result.Projects)

	return &result, err
}
