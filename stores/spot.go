package stores

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/showntop/tripper/models"
)

type SpotStore struct {
	Collection *mgo.Collection
	// *Store
}

func NewSpotStore(store *Store) *SpotStore {
	// Entry implementation
	collection := store.Master.C("spots")

	return &SpotStore{collection}
}

func (ss *SpotStore) Save(spot *models.Spot) error {
	err := ss.Collection.Insert(spot)
	if err != nil {
		return NewStoreError(err.Error())
	}
	return nil
}

func (ss *SpotStore) List(spots *[]*models.Spot) error {
	err := ss.Collection.Find(bson.M{}).All(spots)
	if err != nil {
		return err
	}
	return nil
}
