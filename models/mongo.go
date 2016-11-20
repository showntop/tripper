package models

import (
	"strings"
	"time"

	"gopkg.in/mgo.v2"

	. "github.com/showntop/tripper/config"
)

var session *mgo.Session

func Close(session *mgo.Session) {
	session.Close()
}

func MgoSess() *mgo.Session {
	return session.Copy()
}

func init() {
	// url := beego.AppConfig.String("mongodb::url")

	sess, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    strings.Split(Config.Database["addrs"].(string), "|"),
		Timeout:  15 * time.Second,
		Database: Config.Database["dbname"].(string),
		Username: Config.Database["user"].(string),
		Password: Config.Database["password"].(string),
	})
	if err != nil {
		panic(err)
	}

	session = sess
	session.SetMode(mgo.Monotonic, true)
	createIndexes()
}

func createIndexes() {

	//status 稀疏索引
	//stime etime索引

	//creative created time索引
}
