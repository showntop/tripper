package stores

import (
	"regexp"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/showntop/tripper/errors"
	"github.com/showntop/tripper/models"
)

var (
	// collection = "users"
	indexes = []mgo.Index{
		mgo.Index{
			Key:        []string{"mobile"},
			Unique:     true,
			DropDups:   true,
			Background: true, // See notes.
			Sparse:     true,
		},
	}
)

type UserStore struct {
	Collection *mgo.Collection
	// *Store
}

func NewUserStore(store *Store) *UserStore {
	// Entry implementation
	collection := store.Master.C("users")

	for _, v := range indexes {
		err := collection.EnsureIndex(v)
		if err != nil {
			panic(err)
		}
	}
	return &UserStore{collection}
}

func (us *UserStore) Save(user *models.User) errors.SunError {
	user.PreSave()
	err := us.Collection.Insert(user)
	if err != nil {
		return errors.NewDBError(err.Error())
	}
	user.PostSave()
	return nil
}

func (us *UserStore) FindByAny(fieldValue string) *models.User {
	if bson.IsObjectIdHex(fieldValue) {
		return us.Find(fieldValue)
	}

	if strings.Contains(fieldValue, "@") {
		return us.FindByEmail(fieldValue)
	}

	if ok, _ := regexp.Match(`^1\d{10}$`, []byte(fieldValue)); ok {
		return us.FindByMobile(fieldValue)
	}

	// username
	return us.FindByUsername(fieldValue)
}

// 通过ID得到用户
func (us *UserStore) Find(userId string) *models.User {
	user := &models.User{}
	us.Collection.FindId(bson.ObjectIdHex(userId)).One(&user)
	return user
}

// 通过email得到用户
func (us *UserStore) FindByEmail(email string) *models.User {
	user := &models.User{}
	us.Collection.Find(bson.M{"email": email}).One(&user)
	return user
}

// 通过username得到用户
func (us *UserStore) FindByUsername(username string) *models.User {
	user := &models.User{}
	username = strings.ToLower(username)
	us.Collection.Find(bson.M{"username": username}).One(&user)
	return user
}

// 通过username得到用户
func (us *UserStore) FindByMobile(mobile string) *models.User {
	user := &models.User{}
	us.Collection.Find(bson.M{"mobile": mobile}).One(&user)
	return user
}
