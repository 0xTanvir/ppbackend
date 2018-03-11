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
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Tittle      string        `json:"tittle"`
	Description string        `json:"description"`
	Date        time.Time     `json:"date"`
	Modified    time.Time     `json:"modified"`
	Comments    []Comment     `json:"comments"`
}
