package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	PROJECT_CATEGORY_ARTICLE = iota
	PROJECT_CATEGORY_PICTURE
	PROJECT_CATEGORY_MUSIC
	PROJECT_CATEGORY_MOVIE
)

type Project struct {
	Base `bson:",inline"`

	Category int    `bson:"category" json:"category"`
	Title    string `bson:"title" json:"title"`
	Asset    string `bson:"asset" json:"asset"`
	Intro    string `bson:"intro" json:"intro"`
	Content  string `bson:"content,omitempty" json:"content"`
	Media    string `bson:"media" json:"media"`

	IsDaily bool `bson:"is_daily" json:"-"`

	Author *User `bson:"author" json:"author"`
}

func CreateProject(p *Project) error {
	sess := MgoSess()
	defer sess.Close()

	p.Id = bson.NewObjectId()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	var endPos int = 20
	if len(p.Content) < 20 {
		endPos = len(p.Content)
	}
	p.Intro = p.Content[:endPos]
	return sess.DB(DBNAME).C(C_PROJECT_NAME).Insert(p)
}

func GetPorjectsSelected() ([]*Project, error) {
	sess := MgoSess()
	defer sess.Close()

	var result []*Project = make([]*Project, 0)
	err := sess.DB(DBNAME).C(C_PROJECT_NAME).Find(nil).Sort("-created_at").Limit(10).Select(bson.M{"content": 0}).All(&result)

	return result, err
}

func GetProjectDaily() (*Project, error) {
	sess := MgoSess()
	defer sess.Close()

	result := new(Project)
	err := sess.DB(DBNAME).C(C_PROJECT_NAME).Find(bson.M{"is_daily": true}).Select(bson.M{"content": 0}).One(result)

	return result, err
}

func GetProjectById(id string) (*Project, error) {
	sess := MgoSess()
	defer sess.Close()

	result := new(Project)
	err := sess.DB(DBNAME).C(C_PROJECT_NAME).FindId(bson.ObjectIdHex(id)).One(result)

	return result, err
}
