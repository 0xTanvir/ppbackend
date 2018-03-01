package contest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"

	"github.com/0xTanvir/pp/db"

	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const scheme = "https"
const contestCollection = "contests"

// Service all logic functionality of Account
type Service struct {
	DB *db.DB
}

// IsVIDExist checks is that vjudge id already exist in our database
func (s *Service) IsVIDExist(vid string) bool {
	session := s.DB.Clone()
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	var ctst Ctst

	collection := session.DB("").C(contestCollection)
	// Todo can be make it efficient without insert or skip that ctst
	err := collection.Find(bson.M{"vid": vid}).One(&ctst)
	return err == nil
}

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
