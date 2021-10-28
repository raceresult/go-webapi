package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// BibRanges contains all api endpoints regarding bib ranges
type BibRanges struct {
	api *EventAPI
}

// newBibRanges creates a new BibRanges api endpoint group
func newBibRanges(api *EventAPI) *BibRanges {
	return &BibRanges{
		api: api,
	}
}

// PDF returns a PDF with all bib ranges
func (q *BibRanges) PDF() ([]byte, error) {
	return q.api.Get("bibranges/pdf", nil)
}

// Get returns bib ranges matching the given filters
func (q *BibRanges) Get(contest int, id int) ([]model.BibRange, error) {
	values := urlValues{
		"contest": contest,
		"id":      id,
	}
	bts, err := q.api.Get("bibranges/get", values)
	if err != nil {
		return nil, err
	}

	var dest []model.BibRange
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Delete deletes bib ranges
func (q *BibRanges) Delete(id int) error {
	values := urlValues{
		"id": id,
	}
	_, err := q.api.Get("bibranges/delete", values)
	return err
}

// Save saves bib ranges and returns the bib range IDs
func (q *BibRanges) Save(items []model.BibRange) ([]int, error) {
	bts, err := q.api.Post("bibranges/save", nil, items)
	if err != nil {
		return nil, err
	}
	return parseJsonIntArr(bts)
}
