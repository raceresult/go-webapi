package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// Contests contains all api endpoints regarding contests
type Contests struct {
	api *EventAPI
}

// newContests creates a new Contests api endpoint group
func newContests(api *EventAPI) *Contests {
	return &Contests{
		api: api,
	}
}

// PDF returns a PDF with all contests
func (q *Contests) PDF() ([]byte, error) {
	return q.api.get("contests/pdf", nil)
}

// GetOne returns the contest with the given ID
func (q *Contests) GetOne(id int) (*model.Contest, error) {
	values := urlValues{
		"id": id,
	}
	bts, err := q.api.get("contests/get", values)
	if err != nil {
		return nil, err
	}

	var dest model.Contest
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Get returns all contests
func (q *Contests) Get() ([]model.Contest, error) {
	bts, err := q.api.get("contests/get", nil)
	if err != nil {
		return nil, err
	}

	var dest []model.Contest
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Delete deletes a contest
func (q *Contests) Delete(id int) error {
	values := urlValues{
		"id": id,
	}
	_, err := q.api.get("contests/delete", values)
	return err
}

// Save saves a contest and returns the ID
func (q *Contests) Save(item model.Contest, oldID int) (int, error) {
	values := urlValues{
		"oldID": oldID,
	}
	bts, err := q.api.post("contests/save", values, item)
	if err != nil {
		return 0, err
	}

	var id int
	if err := json.Unmarshal(bts, &id); err != nil {
		return 0, err
	}
	return id, nil
}
