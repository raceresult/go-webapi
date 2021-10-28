package webapi

import (
	"encoding/json"

	"github.com/raceresult/go-model/label"
)

// Labels contains all api endpoints regarding labels
type Labels struct {
	api *EventAPI
}

// newLabels creates a new Labels api endpoint group
func newLabels(api *EventAPI) *Labels {
	return &Labels{
		api: api,
	}
}

// Names returns the names of all labels
func (q *Labels) Names() ([]string, error) {
	bts, err := q.api.get("labels/names", nil)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}

// Get returns a label
func (q *Labels) Get(name string) (*label.Label, error) {
	values := urlValues{
		"name": name,
	}
	bts, err := q.api.get("labels/get", values)
	if err != nil {
		return nil, err
	}
	var dest label.Label
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Save saves a label
func (q *Labels) Save(item *label.Label) error {
	_, err := q.api.post("labels/save", nil, item)
	return err
}

// Delete deletes a label
func (q *Labels) Delete(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("labels/delete", values)
	return err
}

// Copy creates a copy of a label
func (q *Labels) Copy(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("labels/copy", values)
	return err
}

// Rename renames a label
func (q *Labels) Rename(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("labels/rename", values)
	return err
}

// New creates a label
func (q *Labels) New(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("labels/new", values)
	return err
}

// Create returns labels as pdf
func (q *Labels) Create(name string, contests []int, startX, startY int, lang string) ([]byte, error) {
	values := urlValues{
		"name":    name,
		"contest": intSliceToString(contests),
		"startX":  startX,
		"startY":  startY,
		"lang":    lang,
	}
	return q.api.get("labels/create", values)
}
