package contest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/0xTanvir/pp/db"
	"github.com/0xTanvir/pp/helpers/events"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const scheme = "https"
const contestCollection = "contests"
const eventsCollection = "events"
const source = "hr"

// Service all logic functionality of contest
type Service struct {
	DB *db.DB
}

// GetEvents fetch events from mongo
func (s *Service) GetEvents() ([]events.Event, error) {

	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(eventsCollection)

	// Get the collection from database
	var ed EventData
	err := collection.Find(bson.M{"source": "hr"}).One(&ed)
	if err != nil {
		return nil, err
	}

	expiresAt := ed.UpdatedAt.Add(time.Hour * 6).Sub(time.Now())
	fmt.Println(expiresAt)

	// If events get 6 hour old then update the events from cloud calendar
	if expiresAt < 0 {
		fmt.Println("Then update the db from hr")
		evnts, err := s.updateEvents()
		if err != nil {
			return nil, err
		}
		return evnts, nil
	}

	// Now return the list of events
	return ed.Events, nil
}

// updateEvents update event from cloud calendar
func (s *Service) updateEvents() ([]events.Event, error) {
	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(eventsCollection)

	evnts, err := events.GetUpcomingEvents()
	if err != nil {
		return nil, err
	}

	// Update eventes in database
	err = collection.Update(bson.M{"source": "hr"}, bson.M{
		"$set": bson.M{
			"updatedat": time.Now(),
			"events":    *evnts,
		},
	})

	if err != nil {
		return nil, err
	}

	return *evnts, nil
}

// exist check existence of that query.
func (s *Service) exist(query bson.M) bool {
	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	collection := session.DB("").C(contestCollection)
	err := collection.Find(query).Select(nil).One(nil)
	return err == nil
}

// IsVIDExist checks is that vjudge id already exist in our database
func (s *Service) IsVIDExist(vid string) bool {
	return s.exist(bson.M{"vid": vid})
}

// Create a new contest
func (s *Service) Create(ctstInfo CtstInfo) (*bson.ObjectId, error) {
	// Convert string to integer for check
	id, err := strconv.Atoi(ctstInfo.VID)
	if err != nil {
		return nil, err
	}

	// Get contest info from vjudge with vjudge response.
	ctstResponse, err := s.getCtstInfo(ctstInfo.VID)
	if ctstResponse.ID != id {
		return nil, err
	}

	submissions := s.refineSubmissions(ctstResponse)

	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("").C(contestCollection)

	contestData := Ctst{Name: ctstResponse.Title,
		ID:          bson.NewObjectId(),
		Begin:       ctstResponse.Begin,
		Length:      ctstResponse.Length,
		Submissions: *submissions,
		CtstInfo:    ctstInfo,
	}

	return &contestData.ID, collection.Insert(contestData)
}

// Get an Contest by id from the database
func (s *Service) Get(id string) (*Ctst, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("invalid id")
	}

	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	collection := session.DB("").C(contestCollection)
	var contest Ctst

	err := collection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&contest)
	if err != nil {
		return nil, err
	}
	return &contest, err
}

// Update an contest by id
func (s *Service) Update(id string, ctst *Ctst) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("invalid id")
	}

	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	// Todo when an vid get updates then it will again scrap from vjudge
	collection := session.DB("").C(contestCollection)
	return collection.Update(bson.M{
		"_id": bson.ObjectIdHex(id)},
		bson.M{
			"$set": bson.M{
				"remarks":     ctst.Remarks,
				"password":    ctst.Password,
				"vid":         ctst.VID,
				"name":        ctst.Name,
				"begin":       ctst.Begin,
				"length":      ctst.Length,
				"submissions": ctst.Submissions,
			},
		})

}

// Remove a contest by id
func (s *Service) Remove(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("invalid id")
	}

	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	collection := session.DB("").C(contestCollection)
	return collection.Remove(bson.M{"_id": bson.ObjectIdHex(id)})
}

// GetUpcomingContest gets the Upcoming Contest
func (s *Service) GetUpcomingContest() ([]*Ctst, error) {
	var filter = bson.M{}

	// {key:{$gt:value}}
	filter["begin"] = bson.M{"$gt": time.Now().Unix()}

	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	var results []*Ctst
	collection := session.DB("").C(contestCollection)
	err := collection.Find(filter).All(&results)
	if err != nil {
		return nil, err
	}

	return results, err
}

// Find searches the database for a list of products
func (s *Service) Find(qf QueryFilter) ([]*Ctst, error) {
	var filter = bson.M{}

	// Perform full text search on collection
	if len(qf.Query) > 0 {
		filter["q"] = qf.Query
	}
	// Todo add more filter depend on search

	// Calculate how many documents we need to skip
	skip := (qf.Page - 1) * qf.PageSize

	return s.find(filter, skip, qf.PageSize)
}

// find searches the database for one or more products
func (s *Service) find(filter bson.M, skip, limit int) ([]*Ctst, error) {

	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	var results []*Ctst
	collection := session.DB("").C(contestCollection)
	err := collection.Find(filter).Skip(skip).Limit(limit).All(&results)
	if err != nil {
		return nil, err
	}

	return results, err
}

// getCtstInfo gets the contest info from vjudge.
func (s *Service) getCtstInfo(vjudge_id string) (*CtstResponse, error) {
	var u url.URL
	u.Scheme = scheme
	u.Host = viper.GetString("judge.host")
	u.Path = "/user/login"

	form := url.Values{}
	form.Add("username", viper.GetString("judge.username"))
	form.Add("password", viper.GetString("judge.password"))

	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookieJar}

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	u.Path = "/contest/rank/single/" + vjudge_id

	req, err = http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var ctstResponse CtstResponse
	err = json.Unmarshal(body, &ctstResponse)
	if err != nil {
		return nil, err
	}
	return &ctstResponse, nil
}

// refineSubmissions take raw Contest info from vjudge and
// refine the submission and make our submission struct data
func (s *Service) refineSubmissions(vjudgeContest *CtstResponse) *[]Submission {
	var submissions []Submission
	// Iterate each submission in vjudge Contest response
	for _, vjudgeSubmission := range vjudgeContest.Submissions {
		// Check submission time with contest length
		// vjudgeContest.Length is millisecond and vjudgeSubmission[3] is second
		if vjudgeContest.Length >= vjudgeSubmission[3]*1000 {
			submission := Submission{
				Username: vjudgeContest.Participants[strconv.Itoa(vjudgeSubmission[0])][0],
				Problem:  vjudgeSubmission[1],
				Status:   vjudgeSubmission[2],
				Time:     vjudgeSubmission[3],
			}
			submissions = append(submissions, submission)
		}
	}
	return &submissions
}

// IsExist check if an organization id exists
func (s *Service) IsExist(id string) bool {
	if !bson.IsObjectIdHex(id) {
		return false
	}

	return s.exist(bson.M{"_id": bson.ObjectIdHex(id)})
}
