package blog

import (
	"github.com/0xTanvir/pp/auth"
	"github.com/0xTanvir/pp/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const contestCollection = "blog"

// Service all logic functionality of blog
type Service struct {
	DB   *db.DB
	Auth *auth.Service
}

// Create a new contest
func (s *Service) Create(post Post) (*bson.ObjectId, error) {
	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	collection := session.DB("").C(contestCollection)
	post.ID = bson.NewObjectId()
	post.AID = s.Auth.GetUserID().Hex()

	return &post.ID, collection.Insert(post)
}

// GetMyContest is list of contest for user
func (s *Service) GetMyBlog() ([]*Post, error) {

	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	collection := session.DB("").C(contestCollection)
	var posts []*Post

	err := collection.Find(bson.M{"aid": s.Auth.GetUserID().Hex()}).All(&posts)
	if err != nil {
		return nil, err
	}
	return posts, err
}
