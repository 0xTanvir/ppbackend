package users

import (
	"gopkg.in/mgo.v2/bson"
)

// UserInfo without id
type UserInfo struct {
	FirstName string `form:"firstname" json:"firstname" binding:"required"`
	LastName  string `form:"lastname" json:"lastname" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required,email"`
	Password  string `form:"password" json:"password,omitempty"`
	ConatctNo string `form:"contactno" json:"contactno"`
	Vjudge    string `form:"vjudge" json:"vjudge" binding:"required"`
	Uva       string `form:"uva" json:"uva"`
	Cf        string `form:"cf" json:"cf"`
	Cc        string `form:"cc" json:"cc"`
	Timus     string `form:"timus" json:"timus"`
	Tc        string `form:"tc" json:"tc"`
	Urij      string `form:"urij" json:"urij"`
	Lo        string `form:"lo" json:"lo"`
}

// User domain model
type User struct {
	ID       bson.ObjectId `bson:"_id" json:"id"`
	UserInfo `bson:",inline"`
}
