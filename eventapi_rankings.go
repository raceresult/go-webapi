package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// Rankings contains all api endpoints regarding rankings
type Rankings struct {
	api *EventAPI
}

// newRankings creates a new Rankings api endpoint group
func newRankings(api *EventAPI) *Rankings {
	return &Rankings{
		api: api,
	}
}

// Get returns all rankings
func (q *Rankings) Get() ([]model.Ranking, error) {
	bts, err := q.api.Get("ranks/get", nil)
	if err != nil {
		return nil, err
	}

	var dest []model.Ranking
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// GetOne returns the ranking with the given ID
func (q *Rankings) GetOne(id int) (*model.Ranking, error) {
	values := urlValues{
		"id": id,
	}
	bts, err := q.api.Get("ranks/get", values)
	if err != nil {
		return nil, err
	}

	var dest model.Ranking
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Delete deletes ranks
func (q *Rankings) Delete(id int) error {
	values := urlValues{
		"id": id,
	}
	_, err := q.api.Get("ranks/delete", values)
	return err
}

// Save saves ranks and returns the rank IDs
func (q *Rankings) Save(items []model.Ranking) ([]int, error) {
	bts, err := q.api.Post("ranks/save", nil, items)
	if err != nil {
		return nil, err
	}
	return parseJsonIntArr(bts)
}
