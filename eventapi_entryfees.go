package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// EntryFees contains all api endpoints regarding entry fees
type EntryFees struct {
	api *EventAPI
}

// newEntryFees creates a new EntryFees api endpoint group
func newEntryFees(api *EventAPI) *EntryFees {
	return &EntryFees{
		api: api,
	}
}

// PDF returns a PDF with all entry fees
func (q *EntryFees) PDF() ([]byte, error) {
	return q.api.Get("entryfees/pdf", nil)
}

// Get returns entry fees matching the given filters
func (q *EntryFees) Get(contest int, id int) ([]model.EntryFee, error) {
	values := urlValues{
		"contest": contest,
		"id":      id,
	}
	bts, err := q.api.Get("entryfees/get", values)
	if err != nil {
		return nil, err
	}

	var dest []model.EntryFee
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Delete deletes entry fees
func (q *EntryFees) Delete(id int) error {
	values := urlValues{
		"id": id,
	}
	_, err := q.api.Get("entryfees/delete", values)
	return err
}

// Save saves entry fees and returns the entry fee IDs
func (q *EntryFees) Save(items []model.EntryFee) ([]int, error) {
	bts, err := q.api.Post("entryfees/save", nil, items)
	if err != nil {
		return nil, err
	}
	return parseJsonIntArr(bts)
}
