package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// RawDataRules contains all api endpoints regarding raw data rules
type RawDataRules struct {
	api *EventAPI
}

// newRawDataRules creates a new RawDataRules api endpoint group
func newRawDataRules(api *EventAPI) *RawDataRules {
	return &RawDataRules{
		api: api,
	}
}

// Get returns one or all raw data rules
func (q *RawDataRules) Get(id int, resultID int) ([]model.RawDataRule, error) {
	values := urlValues{
		"id":       id,
		"resultID": resultID,
	}
	bts, err := q.api.Get("rawdatarules/get", values)
	if err != nil {
		return nil, err
	}

	var dest []model.RawDataRule
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Delete deletes a raw data rule
func (q *RawDataRules) Delete(id int) error {
	values := urlValues{
		"id": id,
	}
	_, err := q.api.Get("rawdatarules/delete", values)
	return err
}

// Save saves raw data rule and returns the IDs
func (q *RawDataRules) Save(items []model.RawDataRule) ([]int, error) {
	bts, err := q.api.Post("rawdatarules/save", nil, items)
	if err != nil {
		return nil, err
	}
	return parseJsonIntArr(bts)
}
