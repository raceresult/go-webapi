package webapi

import (
	"encoding/json"
	"errors"
	"net/url"
	"time"

	model "github.com/raceresult/go-model"

	"github.com/raceresult/go-model/vbdate"

	"golang.org/x/oauth2"
)

// Public contains all api endpoints regarding functions on public servers
type Public struct {
	api       *API
	sessionID string
}

// newPublic creates a new Public api endpoint group
func newPublic(api *API) *Public {
	return &Public{
		api:       api,
		sessionID: "0",
	}
}

// Login creates a new session
func (q *Public) Login(user, pw string) error {
	values := url.Values{}
	values.Set("user", user)
	values.Set("pw", pw)
	resp, err := q.api.post("", "public/login", nil, "application/x-www-form-urlencoded", values.Encode())
	if err != nil {
		return err
	}
	q.sessionID = string(resp)
	return nil
}

// Logout terminates the session
func (q *Public) Logout() error {
	if q.sessionID == "" {
		return errors.New("not logged in")
	}
	_, err := q.api.get("", "public/logout", nil)
	return err
}

type EventListItem struct {
	ID            string
	UserID        int
	UserName      string
	CheckedOut    bool
	Participants  int
	NotActivated  int
	EventName     string
	EventDate     vbdate.VBDate
	EventDate2    vbdate.VBDate
	EventLocation string
	EventCountry  int
}

// EventList returns a list of events
func (q *Public) EventList(year int, filter string) ([]EventListItem, error) {
	values := urlValues{
		"year":        year,
		"filter":      filter,
		"addsettings": "EventName,EventDate,EventDate2,EventLocation,EventCountry",
	}
	bts, err := q.api.get("", "public/eventlist", values)
	if err != nil {
		return nil, err
	}

	var dest []EventListItem
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// CreateEvent creates a new event (online server only) and returns the new eventID
func (q *Public) CreateEvent(eventName string, eventDate time.Time, eventCountry int, copyOf int, templateID int,
	mode int, laps int) (*EventAPI, error) {

	values := urlValues{
		"eventName":    eventName,
		"eventDate":    eventDate,
		"eventCountry": eventCountry,
		"copyOf":       copyOf,
		"templateID":   templateID,
		"mode":         mode,
		"laps":         laps,
	}

	resp, err := q.api.get("", "public/createevent", values)
	if err != nil {
		return nil, err
	}
	return NewEventAPI(string(resp), q.api), nil
}

// DeleteEvent deletes an event, use with care!
func (q *Public) DeleteEvent(eventID string) error {
	values := urlValues{
		"eventID": eventID,
	}
	_, err := q.api.get("", "public/deleteevent", values)
	return err
}

// TokenFromSession returns an auth token for other rr services
func (q *Public) TokenFromSession() (*oauth2.Token, error) {
	bts, err := q.api.get("", "public/tokenfromsession", nil)
	if err != nil {
		return nil, err
	}

	var dest *oauth2.Token
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// UserInfo returns ID + name of current user
func (q *Public) UserInfo() (*model.UserInfo, error) {
	bts, err := q.api.get("", "public/userinfo", nil)
	if err != nil {
		return nil, err
	}

	var dest model.UserInfo
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// UserRightsGet retrieves the list of users with access rights
func (q *Public) UserRightsGet(eventID string) ([]model.UserRight, error) {
	values := urlValues{
		"eventID": eventID,
	}
	bts, err := q.api.get("", "userrights/get", values)
	if err != nil {
		return nil, err
	}

	var dest []model.UserRight
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// UserRightsSave saves users rights
func (q *Public) UserRightsSave(eventID string, user string, rights string) error {
	values := urlValues{
		"eventID": eventID,
		"user":    user,
		"rights":  rights,
	}
	_, err := q.api.get("", "userrights/save", values)
	return err
}

// UserRightsDelete deletes users rights
func (q *Public) UserRightsDelete(eventID string, userID int) error {
	values := urlValues{
		"eventID": eventID,
		"userID":  userID,
	}
	_, err := q.api.get("", "userrights/delete", values)
	return err
}

// SessionID returns the sessionID after a login
func (q *Public) SessionID() string {
	return q.sessionID
}
