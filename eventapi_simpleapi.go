package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// SimpleAPI contains all api endpoints regarding SimpleAPI
type SimpleAPI struct {
	api *EventAPI
}

// newSimpleAPI creates a new SimpleAPI api endpoint group
func newSimpleAPI(api *EventAPI) *SimpleAPI {
	return &SimpleAPI{
		api: api,
	}
}

// Get returns one simple api entry
func (q *SimpleAPI) Get() ([]model.SimpleAPIItem, error) {
	bts, err := q.api.get("simpleapi/get", nil)
	if err != nil {
		return nil, err
	}

	var dest []model.SimpleAPIItem
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Delete a SimpleAPI entry
func (q *SimpleAPI) Delete(key string) error {
	values := urlValues{
		"key": key,
	}
	_, err := q.api.get("simpleapi/delete", values)
	return err
}

// Save saves SimpleAPI entries
func (q *SimpleAPI) Save(items []model.SimpleAPIItem) error {
	_, err := q.api.post("simpleapi/save", nil, items)
	return err
}

// SaveAll saves SimpleAPI entries and replaces all existing entries
func (q *SimpleAPI) SaveAll(items []model.SimpleAPIItem) error {
	_, err := q.api.post("simpleapi/saveall", nil, items)
	return err
}
