package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// TimingPoints contains all api endpoints regarding timing points
type TimingPoints struct {
	api *EventAPI
}

// newTimingPoints creates a new TimingPoints api endpoint group
func newTimingPoints(api *EventAPI) *TimingPoints {
	return &TimingPoints{
		api: api,
	}
}

// Get returns all timing points
func (q *TimingPoints) Get() ([]model.TimingPoint, error) {
	bts, err := q.api.get("timingpoints/get", nil)
	if err != nil {
		return nil, err
	}

	var dest []model.TimingPoint
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// GetOne returns one timing points
func (q *TimingPoints) GetOne(name string) (*model.TimingPoint, error) {
	values := urlValues{
		"name": name,
	}
	bts, err := q.api.get("timingpoints/get", values)
	if err != nil {
		return nil, err
	}

	var dest model.TimingPoint
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Delete deletes one or all timing points
func (q *TimingPoints) Delete(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("timingpoints/delete", values)
	return err
}

// Save saves a timing point
func (q *TimingPoints) Save(item model.TimingPoint, oldName string) error {
	values := urlValues{
		"oldName": oldName,
	}
	_, err := q.api.post("timingpoints/save", values, item)
	return err
}
