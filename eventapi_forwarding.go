package webapi

import (
	"encoding/json"
)

// Forwarding contains all api endpoints regarding online forwarding
type Forwarding struct {
	api *EventAPI
}

// newParticipant creates a new Participant api endpoint group
func newForwarding(api *EventAPI) *Forwarding {
	return &Forwarding{
		api: api,
	}
}

// ForwardingInfo contains statistics from forwarding
type ForwardingInfo struct {
	BytesSent     int
	BytesReceived int
}

// Active returns true if forwarding is running
func (q *Forwarding) Active() (bool, error) {
	bts, err := q.api.Get("forwarding/active", nil)
	if err != nil {
		return false, err
	}

	var b bool
	err = json.Unmarshal(bts, &b)
	return b, err
}

// Start starts the forwarding process
func (q *Forwarding) Start(hostname, eventid, authToken string) error {
	values := urlValues{
		"hostname":  hostname,
		"eventid":   eventid,
		"authToken": authToken,
	}
	_, err := q.api.Get("forwarding/start", values)
	return err
}

// Restart restarts the forwarding process with previous settings
func (q *Forwarding) Restart() error {
	_, err := q.api.Get("forwarding/restart", nil)
	return err
}

// Stop stops the forwarding process
func (q *Forwarding) Stop() error {
	_, err := q.api.Get("forwarding/stop", nil)
	return err
}

// Info returns statistics about the forwarding process
func (q *Forwarding) Info() (ForwardingInfo, error) {
	bts, err := q.api.Get("forwarding/info", nil)
	if err != nil {
		return ForwardingInfo{}, err
	}

	var dest ForwardingInfo
	err = json.Unmarshal(bts, &dest)
	return dest, err
}
