package webapi

import (
	"encoding/json"

	"github.com/raceresult/go-model/registration"
)

// Registrations contains all api endpoints regarding certificate sets
type Registrations struct {
	api *EventAPI
}

// newRegistrations creates a new Registrations api endpoint group
func newRegistrations(api *EventAPI) *Registrations {
	return &Registrations{
		api: api,
	}
}

// Names returns the names of all certificate sets
func (q *Registrations) Names() ([]string, error) {
	bts, err := q.api.get("registrations/names", nil)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}

// Get returns a certificate sets
func (q *Registrations) Get(name string) (*registration.Registration, error) {
	values := urlValues{
		"name": name,
	}
	bts, err := q.api.get("registrations/get", values)
	if err != nil {
		return nil, err
	}
	var dest registration.Registration
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Save saves a certificate sets
func (q *Registrations) Save(item *registration.Registration) error {
	_, err := q.api.post("registrations/save", nil, item)
	return err
}

// Delete deletes a certificate sets
func (q *Registrations) Delete(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("registrations/delete", values)
	return err
}

// Copy creates a copy of a certificate sets
func (q *Registrations) Copy(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("registrations/copy", values)
	return err
}

// Rename renames a certificate sets
func (q *Registrations) Rename(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("registrations/rename", values)
	return err
}

// New creates a certificate sets
func (q *Registrations) New(name string, group bool) error {
	values := urlValues{
		"name":  name,
		"group": group,
	}
	_, err := q.api.get("registrations/new", values)
	return err
}
