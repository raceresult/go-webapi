package webapi

import (
	"encoding/json"
	"strconv"
	"strings"

	model "github.com/raceresult/go-model"
)

// Splits contains all api endpoints regarding splits
type Splits struct {
	api *EventAPI
}

// newSplits creates a new Splits api endpoint group
func newSplits(api *EventAPI) *Splits {
	return &Splits{
		api: api,
	}
}

// Get returns the splits of one or all contests
func (q *Splits) Get(contest int) ([]model.Split, error) {
	values := urlValues{
		"contest": contest,
	}
	bts, err := q.api.get("splits/get", values)
	if err != nil {
		return nil, err
	}

	var dest []model.Split
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// GetOne returns the split with the given ID
func (q *Splits) GetOne(id int) (*model.Split, error) {
	values := urlValues{
		"id": id,
	}
	bts, err := q.api.get("splits/get", values)
	if err != nil {
		return nil, err
	}

	var dest model.Split
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Delete deletes splits
func (q *Splits) Delete(ids []int) error {
	sids := make([]string, 0, len(ids))
	for _, id := range ids {
		sids = append(sids, strconv.Itoa(id))
	}
	values := urlValues{
		"id": strings.Join(sids, ","),
	}
	_, err := q.api.get("splits/delete", values)
	return err
}

// Save saves split and returns the IDs
func (q *Splits) Save(items []model.Split) ([]int, error) {
	bts, err := q.api.post("splits/save", nil, items)
	if err != nil {
		return nil, err
	}
	return parseJsonIntArr(bts)
}
