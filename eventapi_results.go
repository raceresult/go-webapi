package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// Results contains all api endpoints regarding results
type Results struct {
	api *EventAPI
}

// newResults creates a new Results api endpoint group
func newResults(api *EventAPI) *Results {
	return &Results{
		api: api,
	}
}

// Get returns results matching the given filters
func (q *Results) Get(name string, onlyFormulas, onlyNoFormulas bool) ([]model.Result, error) {
	values := urlValues{
		"name":           name,
		"onlyFormulas":   onlyFormulas,
		"onlyNoFormulas": onlyNoFormulas,
	}
	bts, err := q.api.Get("results/get", values)
	if err != nil {
		return nil, err
	}

	var dest []model.Result
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// GetOne returns results with the given ID
func (q *Results) GetOne(id int) (*model.Result, error) {
	values := urlValues{
		"id": id,
	}
	bts, err := q.api.Get("results/get", values)
	if err != nil {
		return nil, err
	}

	var dest model.Result
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Delete deletes results
func (q *Results) Delete(id int) error {
	values := urlValues{
		"id": id,
	}
	_, err := q.api.Get("results/delete", values)
	return err
}

// Save saves results and returns the result IDs
func (q *Results) Save(items []model.Result) error { // TODO: oldID
	_, err := q.api.Post("results/save", nil, items)
	return err
}
