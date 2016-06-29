package stores

import (
	"gopkg.in/mgo.v2"

	"github.com/showntop/tripper/database"
)

type Store struct {
	Master *mgo.Database

	User *UserStore
	// Relationship *RelationshipStore
	Spot *SpotStore
}

func NewStore() *Store {
	database.InitMongo()

	store := &Store{}
	store.Master = database.Mongodb

	store.User = NewUserStore(store)
	// store.Relationship = &RelationshipStore{store.Master.C("relationships")}
	store.Spot = NewSpotStore(store)

	return store
}
