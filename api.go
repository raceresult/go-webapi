package webapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync/atomic"
	"time"
)

// API is the main object of the web api
type API struct {
	server     string
	secure     bool
	public     *Public
	timeout_ms int32
	userAgent  string
}

// NewAPI creates a new API object
func NewAPI(server string, https bool, userAgent string) *API {
	if userAgent == "" {
		userAgent = "webapi/1.0"
	}
	q := &API{
		server:     server,
		secure:     https,
		timeout_ms: 30000,
		userAgent:  userAgent,
	}
	q.public = newPublic(q)
	return q
}

// EventAPI returns an EventAPI for the given event
func (q *API) EventAPI(eventID string) *EventAPI {
	return NewEventAPI(eventID, q)
}

// Public returns the endpoint group for public servers
func (q *API) Public() *Public {
	return q.public
}

// General returns the endpoint group for general functions
func (q *API) General() *General {
	return newGeneral(q)
}

// SetTimeout sets the timeout for all following requests. Default is 30 seconds.
func (q *API) SetTimeout(timeout time.Duration) {
	atomic.StoreInt32(&q.timeout_ms, int32(timeout.Milliseconds()))
}

// GetTimeout returns the current request timeout
func (q *API) GetTimeout() time.Duration {
	return time.Duration(atomic.LoadInt32(&q.timeout_ms)) * time.Millisecond
}

// SessionID returns the session ID
func (q *API) SessionID() string {
	return q.public.sessionID
}

// get makes a GET request on the server
func (q *API) get(eventID, cmd string, values urlValues) ([]byte, error) {
	req, err := http.NewRequest("GET", q.buildURL(eventID, cmd, values), nil)
	if err != nil {
		return nil, err
	}
	return q.do(req)
}

// post makes a POST request on the server
func (q *API) post(eventID, cmd string, values urlValues, contentType string, data interface{}) ([]byte, error) {
	var reader io.Reader
	if data != nil {
		switch d := data.(type) {
		case []byte:
			reader = bytes.NewReader(d)
		case string:
			reader = strings.NewReader(d)
		default:
			btsData, err := json.Marshal(data)
			if err != nil {
				return nil, err
			}
			reader = bytes.NewReader(btsData)
		}
	}

	// make request
	req, err := http.NewRequest("POST", q.buildURL(eventID, cmd, values), reader)
	if err != nil {
		return nil, err
	}
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	return q.do(req)
}

// do executes a http request
func (q *API) do(req *http.Request) ([]byte, error) {
	// make request
	client := http.Client{
		Timeout: q.GetTimeout(),
	}
	req.Header.Set("Authorization", "Bearer "+q.public.sessionID)
	req.Header.Set("User-Agent", q.userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// read response
	defer func() { _ = resp.Body.Close() }()
	bts, err := ioutil.ReadAll(resp.Body)

	// check http code
	if resp.StatusCode != 200 {
		return bts, errors.New(string(bts))
	}

	// return without error
	return bts, err
}

// buildURL compiles the url for any request
func (q *API) buildURL(eventID, cmd string, values urlValues) string {
	var sb strings.Builder
	if q.secure {
		sb.WriteString("https://")
	} else {
		sb.WriteString("http://")
	}
	sb.WriteString(q.server)
	if eventID != "" {
		sb.WriteString("/_")
		sb.WriteString(eventID)
	}
	sb.WriteString("/api/")
	sb.WriteString(cmd)
	if len(values) != 0 {
		sb.WriteString("?")
		sb.WriteString(values.URLEncode())
	}
	return sb.String()
}
