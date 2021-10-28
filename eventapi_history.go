package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
	"github.com/raceresult/go-model/vbdate"
)

// History contains all api endpoints regarding history entries
type History struct {
	api *EventAPI
}

// newParticipant creates a new Participant api endpoint group
func newHistory(api *EventAPI) *History {
	return &History{
		api: api,
	}
}

// Get returns history entries matching the given filters
func (q *History) Get(bib int) ([]model.History, error) {
	values := urlValues{
		"bib": bib,
	}
	bts, err := q.api.get("history/get", values)
	if err != nil {
		return nil, err
	}

	var dest []model.History
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// ExcelExport returns history entries matching the given filters as csv file
func (q *History) ExcelExport(bib int, lang string) ([]byte, error) {
	values := urlValues{
		"bib":  bib,
		"lang": lang,
	}
	return q.api.get("history/excelexport", values)
}

// Delete deletes history entries matching the given filters
func (q *History) Delete(bib int, contest int, field string, dateForm, dateTo vbdate.VBDate, filter string) error {
	values := urlValues{
		"bib":      bib,
		"contest":  contest,
		"field":    field,
		"dateForm": dateForm,
		"dateTo":   dateTo,
		"filter":   filter,
	}
	_, err := q.api.get("history/delete", values)
	return err
}

// Count counts history entries matching the given filters
func (q *History) Count(bib int, contest int, field string, dateForm, dateTo vbdate.VBDate, filter string) (int, error) {
	values := urlValues{
		"bib":      bib,
		"contest":  contest,
		"field":    field,
		"dateForm": dateForm,
		"dateTo":   dateTo,
		"filter":   filter,
	}
	bts, err := q.api.get("history/count", values)
	if err != nil {
		return 0, err
	}

	var count int
	err = json.Unmarshal(bts, &count)
	return count, err
}
