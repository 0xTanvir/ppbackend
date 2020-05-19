package events

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

// Event is the main data model for contest for our db
type Event struct {
	Description string        `json:"description"`
	Location    string        `json:"location"`
	Start       string        `json:"start"`
	End         string        `json:"end"`
	Summary     string        `json:"summary"`
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	tokFile := "./etc/token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	log.Info("Go to the following link in your browser then type the ",
		"authorization code: \n", authURL)
	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	log.Infof("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	json.NewEncoder(f).Encode(token)
}

// GetUpcomingEvents fetch data from google calendar and
// return a events list.
func GetUpcomingEvents() (*[]Event, error) {
	b, err := ioutil.ReadFile("./etc/client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved client_secret.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	srv, err := calendar.New(getClient(config))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	t := time.Now().Format(time.RFC3339)
	events, err := srv.Events.List("8a4ko50nq55ma5smhuhdp5rpmaek45lg@import.calendar.google.com").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).TimeZone("GMT+6:00").MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}

	var Events []Event

	log.Info("Upcoming events:")
	if len(events.Items) == 0 {
		return nil, errors.New("no upcoming events found")
	} else {
		for _, item := range events.Items {

			sDate := item.Start.DateTime
			if sDate == "" {
				sDate = item.Start.Date
			}

			event := Event{
				Description: item.Description,
				Location:    item.Location,
				Start:       sDate,
				End:         item.End.DateTime,
				Summary:     item.Summary,
			}
			Events = append(Events, event)

			//fmt.Printf("%v (%v)\n", item.Summary, item.Description)
			//fmt.Printf("------------------------")
			//fmt.Printf("%+v\n", item)
			//fmt.Printf("(%v)\n", end)
		}
		return &Events, nil
	}
}
