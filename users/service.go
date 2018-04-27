package users

import (
	"github.com/0xTanvir/pp/db"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const userCollection = "users"

// Service all logic functionality of User
type Service struct {
	DB *db.DB
}

// Create new user
func (r *Service) Create(userInfo *UserInfo) (*bson.ObjectId, error) {

	session := r.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(userCollection)

	user := User{ID: bson.NewObjectId(),
		UserInfo: *userInfo,
	}
	
	return &user.ID, collection.Insert(user)
}
