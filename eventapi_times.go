package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"

	"github.com/raceresult/go-model/decimal"
)

// Times contains all api endpoints regarding times
type Times struct {
	api *EventAPI
}

// newTimes creates a new Times api endpoint group
func newTimes(api *EventAPI) *Times {
	return &Times{
		api: api,
	}
}

// ExcelExport returns times matching the given filters as csv file
func (q *Times) ExcelExport(bib int, result int, lang string) ([]byte, error) {
	values := urlValues{
		"bib":    bib,
		"result": result,
		"lang":   lang,
	}
	return q.api.Get("times/excelexport", values)
}

// Delete deletes times matching the given filters
func (q *Times) Delete(bib, contest, result int, filter string, filterInfo string) error {
	values := urlValues{
		"bib":        bib,
		"contest":    contest,
		"result":     result,
		"filter":     filter,
		"filterInfo": filterInfo,
	}
	_, err := q.api.Get("times/delete", values)
	return err
}

// Swap swaps the times of two participants
func (q *Times) Swap(bib1, bib2 int) error {
	values := urlValues{
		"bib1": bib1,
		"bib2": bib2,
	}
	_, err := q.api.Get("times/swap", values)
	return err
}

// SingleStart creates single start times
func (q *Times) SingleStart(result int, contest int, firstTime decimal.Decimal, interval decimal.Decimal, sort string,
	filter string, noHistory bool) error {

	values := urlValues{
		"result":    result,
		"contest":   contest,
		"firstTime": firstTime,
		"interval":  interval,
		"sort":      sort,
		"filter":    filter,
		"noHistory": noHistory,
	}
	_, err := q.api.Get("times/singlestart", values)
	return err
}

// RandomTimes creates random times
func (q *Times) RandomTimes(result int, contest int, minTime decimal.Decimal, maxTime decimal.Decimal, offsetResult int,
	filter string, noHistory bool) error {

	values := urlValues{
		"result":       result,
		"contest":      contest,
		"minTime":      minTime,
		"maxTime":      maxTime,
		"offsetResult": offsetResult,
		"filter":       filter,
		"noHistory":    noHistory,
	}
	_, err := q.api.Get("times/randomtimes", values)
	return err
}

// Copy copies times from one participant to another
func (q *Times) Copy(bibFrom, bibTo int, overwriteExisting bool) error {
	values := urlValues{
		"bibFrom":           bibFrom,
		"bibTo":             bibTo,
		"overwriteExisting": overwriteExisting,
	}
	_, err := q.api.Get("times/copy", values)
	return err
}

// Interpolate interpolates missing times
func (q *Times) Interpolate(destID, helperID int, contest int, helpers int) error {
	values := urlValues{
		"destID":   destID,
		"helperID": helperID,
		"contest":  contest,
		"helpers":  helpers,
	}
	_, err := q.api.Get("times/interpolate", values)
	return err
}

// Get returns times matching the given filters
func (q *Times) Get(bib int, result int) ([]model.Time, error) {
	values := urlValues{
		"bib":    bib,
		"result": result,
	}
	bts, err := q.api.Get("times/get", values)
	if err != nil {
		return nil, err
	}

	var dest []model.Time
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Count counts times matching the given filters
func (q *Times) Count(bib int, contest int, result int, filter string) (int, error) {
	values := urlValues{
		"bib":     bib,
		"contest": contest,
		"result":  result,
		"filter":  filter,
	}
	bts, err := q.api.Get("times/count", values)
	if err != nil {
		return 0, err
	}

	var count int
	err = json.Unmarshal(bts, &count)
	return count, err
}

// Add adds times/passings
func (q *Times) Add(passings []model.Passing, returnFields []string, contestFilter int, ignoreBibToBibAssign bool) (
	[]model.TimesAddResponseItem, error) {

	values := urlValues{
		"returnFields":         returnFields,
		"contestFilter":        contestFilter,
		"ignoreBibToBibAssign": ignoreBibToBibAssign,
	}
	bts, err := q.api.Post("times/add", values, passings)
	if err != nil {
		return nil, err
	}

	var dest []model.TimesAddResponseItem
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}
