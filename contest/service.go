package contest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/0xTanvir/pp/db"

	"github.com/spf13/viper"
)

const scheme = "https"

// Service all logic functionality of Account
type Service struct {
	DB *db.DB
}

func (s *Service) Create(vjudge_id string) (*CtstResponse, error) {
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
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
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

	return &ctstResponse, err
}
