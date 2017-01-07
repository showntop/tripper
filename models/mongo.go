package models

import (
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
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
	log.Infoln("...................mongo..................")
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
	log.Infof("addr: %s", Config.Database["addrs"].(string))
	log.Infoln("...................mongo..................")
	session = sess
	session.SetMode(mgo.Monotonic, true)
	createIndexes()
}

func createIndexes() {
	//status 稀疏索引
	//stime etime索引

	//creative created time索引
}
