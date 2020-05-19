package contest

import (
	"time"

	"github.com/0xTanvir/pp/helpers/events"

	"gopkg.in/mgo.v2/bson"
)

// CtstInfo contains string parameters of a contest
type CtstInfo struct {
	Remarks  string `json:"remarks"`
	Password string `json:"password"`
	VID      string `json:"vid" binding:"required"`
}

// Submission is the struct for contest submission
type Submission struct {
	Username string `json:"username"`
	Problem  int    `json:"problem"`
	Status   int    `json:"status"`
	Time     int    `json:"time"`
}

// Ctst is the main data model for contest for our db
type Ctst struct {
	CtstInfo    `bson:",inline"`
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `json:"title"`
	Begin       int64         `json:"begin"`
	Length      int           `json:"length"`
	AID         string        `bson:"aid" json:"aid"`
	Submissions []Submission  `json:"submissions"`
}

// CtstResponse contains vjudge response from vjudge
type CtstResponse struct {
	Submissions  [][]int             `json:"submissions"`
	Participants map[string][]string `json:"participants"`
	Title        string              `json:"title"`
	Begin        int64               `json:"begin"`
	Length       int                 `json:"length"` // millisecond
	ID           int                 `json:"id"`
}

// QueryFilter contains the query string parameters of a search request
type QueryFilter struct {
	Query    string `form:"q"`
	Page     int    `form:"page"`
	PageSize int    `form:"size"`
}

// EventData contains upcomming event
type EventData struct {
	Source    string         `json:"source"`
	UpdatedAt time.Time      `json:"updatedat"`
	Events    []events.Event `json:"events"`
}
