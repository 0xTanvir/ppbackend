package blog

import (
	"github.com/0xTanvir/pp/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const contestCollection = "blog"

// Service all logic functionality of blog
type Service struct {
	DB *db.DB
}

// Create a new contest
func (s *Service) Create(post Post) (*bson.ObjectId, error) {
	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	collection := session.DB("").C(contestCollection)
	post.ID = bson.NewObjectId()

	return &post.ID, collection.Insert(post)
}
