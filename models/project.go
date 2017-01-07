package models

import (
	"fmt"
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

	Album struct {
		Id   string `bson:"id" json:"id"`
		Name string `bson:"name" json:"name"`
	} `bson:"album" json:"album"`

	Author *User `bson:"author" json:"author"`

	Comments []ProjectComment `bson:"-" json:"comments"`
}

type ProjectComment struct {
	Base      `bson:",inline"`
	ProjectId string `bson:"project_id" json:"project_id"`
	Content   string `bson:"content" json:"content"`
	Commentor *User  `bson:"commentor" json:"commentor"`
}

type ProjectLike struct {
	Base      `bson:",inline"`
	ProjectId string `bson:"project_id" json:"project_id"`
	Liker     *User  `bson:"liker" json:"liker"`
}

func (p *Project) Validate() error {
	if len(p.Title) <= 0 {
		return fmt.Errorf("project validate: %s", "title length < 1")
	}
	if len(p.Content) < 50 {
		return fmt.Errorf("project validate: %s", "content length < 50")
	}
	return nil
}

func CreateProject(p *Project) error {
	sess := MgoSess()
	defer sess.Close()

	p.Id = bson.NewObjectId()
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	var endPos int = 20
	if len([]rune(p.Content)) < 20 {
		endPos = len([]rune(p.Content))
	}
	p.Intro = string([]rune(p.Content)[:endPos])
	return sess.DB(DBNAME).C(C_PROJECT_NAME).Insert(p)
}

func CreateProjectComment(pc *ProjectComment) error {
	sess := MgoSess()
	defer sess.Close()

	pc.Id = bson.NewObjectId()
	pc.CreatedAt = time.Now()
	pc.UpdatedAt = time.Now()

	return sess.DB(DBNAME).C(C_PROJECT_COMMENT_NAME).Insert(pc)
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
	if err != nil {
		return result, err
	}
	err = sess.DB(DBNAME).C(C_PROJECT_COMMENT_NAME).Find(bson.M{"project_id": id}).Limit(20).All(&result.Comments)

	return result, err
}

func CreateProjectLike(pl *ProjectLike) error {
	sess := MgoSess()
	defer sess.Close()

	if !pl.Id.Valid() {
		pl.Id = bson.NewObjectId()
	}

	_, err := sess.DB(DBNAME).C(C_PROJECT_LIKE_NAME).Upsert(bson.M{"project_id": pl.ProjectId, "liker.id": pl.Liker.Id}, bson.M{"$set": pl})
	if err != nil {
		return err
	}

	return err
}

func DeleteProjectLike(pl *ProjectLike) error {
	sess := MgoSess()
	defer sess.Close()

	err := sess.DB(DBNAME).C(C_PROJECT_LIKE_NAME).Remove(bson.M{"project_id": pl.ProjectId, "liker.id": pl.Liker.Id})
	if err != nil {
		return err
	}

	return err
}
