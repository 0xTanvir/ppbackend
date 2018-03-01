package contest

import "gopkg.in/mgo.v2/bson"

type CtstInfo struct {
	Remarks  string `json:"remarks"`
	Password string `json:"password"`
	VID      string `json:"vid" binding:"required"`
}

type Submission struct {
	Username string `json:"username"`
	Problem  int    `json:"problem"`
	Status   int    `json:"status"`
	Time     int    `json:"time"`
}

type Ctst struct {
	CtstInfo    `bson:",inline"`
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `json:"title"`
	Begin       int64         `json:"begin"`
	Length      int           `json:"length"`
	Submissions []Submission  `json:"submissions"`
}

type CtstResponse struct {
	Submissions  [][]int             `json:"submissions"`
	Participants map[string][]string `json:"participants"`
	Title        string              `json:"title"`
	Begin        int64               `json:"begin"`
	Length       int                 `json:"length"` // millisecond
	ID           int                 `json:"id"`
}
