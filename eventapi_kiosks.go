package webapi

import (
	"encoding/json"

	"github.com/raceresult/go-model/kiosk"
)

// Kiosks contains all api endpoints regarding kiosks
type Kiosks struct {
	api *EventAPI
}

// newKiosks creates a new Kiosks api endpoint group
func newKiosks(api *EventAPI) *Kiosks {
	return &Kiosks{
		api: api,
	}
}

// Names returns the names of all kiosks
func (q *Kiosks) Names() ([]string, error) {
	bts, err := q.api.get("kiosks/names", nil)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}

// Get returns a kiosk
func (q *Kiosks) Get(name string) (*kiosk.Kiosk, error) {
	values := urlValues{
		"name": name,
	}
	bts, err := q.api.get("kiosks/get", values)
	if err != nil {
		return nil, err
	}
	var dest kiosk.Kiosk
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Save saves a kiosk
func (q *Kiosks) Save(item *kiosk.Kiosk) error {
	_, err := q.api.post("kiosks/save", nil, item)
	return err
}

// Delete deletes a kiosk
func (q *Kiosks) Delete(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("kiosks/delete", values)
	return err
}

// Copy creates a copy of a kiosk
func (q *Kiosks) Copy(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("kiosks/copy", values)
	return err
}

// Rename renames a kiosk
func (q *Kiosks) Rename(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("kiosks/rename", values)
	return err
}

// New creates a kiosk
func (q *Kiosks) New(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("kiosks/new", values)
	return err
}
