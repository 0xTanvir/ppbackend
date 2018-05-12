package blog

import (
	"github.com/0xTanvir/pp/auth"
	"github.com/0xTanvir/pp/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

const blogCollection = "blog"

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

	collection := session.DB("").C(blogCollection)
	post.ID = bson.NewObjectId()
	post.AID = s.Auth.GetUserID().Hex()
	post.HID = post.ID.Hex()

	return &post.ID, collection.Insert(post)
}


// GetEachBlog gets Each Blog
func (s *Service) GetEachBlog(id string) (*Post, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("invalid id")
	}
	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(blogCollection)
	var post Post
	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&post)
	if err != nil {
		return nil, err
	}
	return &post, err
}

// GetMyContest is list of contest for user
func (s *Service) GetBlog() ([]*Post, error) {

	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	collection := session.DB("").C(blogCollection)
	var posts []*Post

	err := collection.Find(nil).Limit(3).All(&posts)
	if err != nil {
		return nil, err
	}
	return posts, err
}

// GetMyContest is list of contest for user
func (s *Service) GetMyBlog() ([]*Post, error) {

	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	collection := session.DB("").C(blogCollection)
	var posts []*Post

	err := collection.Find(bson.M{"aid": s.Auth.GetUserID().Hex()}).Limit(3).All(&posts)
	if err != nil {
		return nil, err
	}
	return posts, err
}
