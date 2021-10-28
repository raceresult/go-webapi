package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// TimingPointRules contains all api endpoints regarding timing point rules
type TimingPointRules struct {
	api *EventAPI
}

// newTimingPointRules creates a new TimingPointRules api endpoint group
func newTimingPointRules(api *EventAPI) *TimingPointRules {
	return &TimingPointRules{
		api: api,
	}
}

// Get returns all timing point rules
func (q *TimingPointRules) Get() ([]model.TimingPointRule, error) {
	bts, err := q.api.get("timingpointrules/get", nil)
	if err != nil {
		return nil, err
	}

	var dest []model.TimingPointRule
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// GetOne returns one timing point rules
func (q *TimingPointRules) GetOne(id int) (*model.TimingPointRule, error) {
	values := urlValues{
		"id": id,
	}
	bts, err := q.api.get("timingpointrules/get", values)
	if err != nil {
		return nil, err
	}

	var dest model.TimingPointRule
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Delete deletes a timing point rule
func (q *TimingPointRules) Delete(id int) error {
	values := urlValues{
		"id": id,
	}
	_, err := q.api.get("timingpointrules/delete", values)
	return err
}

// Save saves timing point rule and returns the IDs
func (q *TimingPointRules) Save(items []model.TimingPointRule) ([]int, error) {
	bts, err := q.api.post("timingpointrules/save", nil, items)
	if err != nil {
		return nil, err
	}
	return parseJsonIntArr(bts)
}
