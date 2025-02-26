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
func (q *Times) ExcelExport(identifier Identifier, result int, lang string) ([]byte, error) {
	values := urlValues{
		identifier.Name: identifier.Value,
		"result":        result,
		"lang":          lang,
	}
	return q.api.get("times/excelexport", values)
}

// Delete deletes times matching the given filters
func (q *Times) Delete(identifier Identifier, contest, result int, filter string, filterInfo string) error {
	values := urlValues{
		identifier.Name: identifier.Value,
		"contest":       contest,
		"result":        result,
		"filter":        filter,
		"filterInfo":    filterInfo,
	}
	_, err := q.api.get("times/delete", values)
	return err
}

// Swap swaps the times of two participants
func (q *Times) Swap(identifier1, identifier2 Identifier) error {
	values := urlValues{
		identifier1.Name + "1": identifier1.Value,
		identifier2.Name + "2": identifier2.Value,
	}
	_, err := q.api.get("times/swap", values)
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
	_, err := q.api.get("times/singlestart", values)
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
	_, err := q.api.get("times/randomtimes", values)
	return err
}

// Copy copies times from one participant to another
func (q *Times) Copy(from, to Identifier, overwriteExisting bool) error {
	values := urlValues{
		from.Name + "From":  from.Value,
		from.Name + "To":    to.Value,
		"overwriteExisting": overwriteExisting,
	}
	_, err := q.api.get("times/copy", values)
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
	_, err := q.api.get("times/interpolate", values)
	return err
}

// Get returns times matching the given filters
func (q *Times) Get(identifier Identifier, result int) ([]model.Time, error) {
	values := urlValues{
		identifier.Name: identifier.Value,
		"result":        result,
	}
	bts, err := q.api.get("times/get", values)
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
func (q *Times) Count(identifier Identifier, contest int, result int, filter string) (int, error) {
	values := urlValues{
		identifier.Name: identifier.Value,
		"contest":       contest,
		"result":        result,
		"filter":        filter,
	}
	bts, err := q.api.get("times/count", values)
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
	bts, err := q.api.post("times/add", values, passings)
	if err != nil {
		return nil, err
	}

	var dest []model.TimesAddResponseItem
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}
