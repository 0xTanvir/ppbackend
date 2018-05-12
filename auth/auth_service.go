package auth

import (
	"gopkg.in/mgo.v2/bson"
)

// Service to manage userID/accountID auth
type Service struct {
	userID    bson.ObjectId
	LoggedIn bool
	name string
}

// SetUserID sets the current user's id on service
func (r *Service) SetUserID(id string) bool {
	if !bson.IsObjectIdHex(id) {
		return false
	}
	r.userID = bson.ObjectIdHex(id)
	return true
}

// GetUserID gets the current user's id
func (r *Service) GetUserID() bson.ObjectId {
	return r.userID
}

// SetLoggedIn sets the current user's logedin status
func (r *Service) SetLoggedIn(status bool) bool {
	r.LoggedIn = status
	return true
}

// GetLoggedIn gets the current user's id
func (r *Service) GetLoggedIn() bool {
	return r.LoggedIn
}

// SetName sets the current user's role on service
func (r *Service) SetName(name string) bool {
	r.name = name
	return true
}

// GetName gets the current account role
func (r *Service) GetName() string {
	return r.name
}