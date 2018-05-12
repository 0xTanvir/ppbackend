package blog

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Comment contains a post's comment
type Comment struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Description string        `json:"description"`
	CreatedAt   time.Time     `json:"createdat"`
	UpdatedAt   time.Time     `json:"updatedat"`
}

// Post is the main struct for blog post
type Post struct {
	ID          bson.ObjectId `bson:"_id" json:"id" form:"id"`
	Tittle      string        `json:"tittle" form:"tittle"`
	AID         string        `bson:"aid" json:"aid"`
	HID         string        `bson:"hid" json:"hid"`
	Description string        `json:"description" form:"description"`
	Date        time.Time     `json:"date" form:"date"`
	Modified    time.Time     `json:"modified" form:"modified"`
	Comments    []Comment     `json:"comments" form:"comments"`
}
