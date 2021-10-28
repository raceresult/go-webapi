package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"

	"github.com/raceresult/go-model/decimal"
)

// RawData contains all api endpoints regarding raw data entries
type RawData struct {
	api *EventAPI
}

// newRawData creates a new RawData api endpoint group
func newRawData(api *EventAPI) *RawData {
	return &RawData{
		api: api,
	}
}

/*
// Get returns raw data entries
func (q *RawData) Get(bib int) ([]model.RawData, error) {
	bts, err := q.api.get("rawdata/get", nil)
	if err != nil {
		return nil, err
	}

	var dest []model.RawData
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}*/

// ExcelExport returns raw data entries matching the given filters as csv file
func (q *RawData) ExcelExport(bib int, lang string) ([]byte, error) {
	values := urlValues{
		"bib":  bib,
		"lang": lang,
	}
	return q.api.get("rawdata/excelexport", values)
}

// SetInvalid sets a raw data entry to valid or invalid
func (q *RawData) SetInvalid(id int, invalid bool) error {
	values := urlValues{
		"id":      id,
		"invalid": invalid,
	}
	_, err := q.api.get("rawdata/setinvalid", values)
	return err
}

// DeleteID deletes the raw data entry with the given ID
func (q *RawData) DeleteID(id int) error {
	values := urlValues{
		"id": id,
	}
	_, err := q.api.get("rawdata/deleteid", values)
	return err
}

// Delete deletes the raw data entries matching the given filters
func (q *RawData) Delete(timingPoint string, contest int, bib int, filter string,
	minTime decimal.Decimal, maxtime decimal.Decimal, filterInfo string) error {

	values := urlValues{
		"timingPoint": timingPoint,
		"contest":     contest,
		"bib":         bib,
		"filter":      filter,
		"minTime":     minTime,
		"maxtime":     maxtime,
		"filterInfo":  filterInfo,
	}
	_, err := q.api.get("rawdata/delete", values)
	return err
}

// SetInvalidBatch sets multiple raw data entries to valid or invalid
func (q *RawData) SetInvalidBatch(invalid bool, timingPoint string, timeFrom, timeTo decimal.Decimal, filter string) error {
	values := urlValues{
		"invalid":     invalid,
		"timingPoint": timingPoint,
		"timeFrom":    timeFrom,
		"timeTo":      timeTo,
		"filter":      filter,
	}
	_, err := q.api.get("rawdata/setinvalidbatch", values)
	return err
}

// Count counts raw data entries matching the given filters
func (q *RawData) Count(timingPoint string, bib int, contest int, filter string, minTime, maxTime decimal.Decimal) (int, error) {
	values := urlValues{
		"timingPoint": timingPoint,
		"bib":         bib,
		"contest":     contest,
		"filter":      filter,
		"minTime":     minTime,
		"maxTime":     maxTime,
	}
	bts, err := q.api.get("rawdata/count", values)
	if err != nil {
		return 0, err
	}

	var count int
	err = json.Unmarshal(bts, &count)
	return count, err
}

// Copy copies raw data from one participant to another
func (q *RawData) Copy(bibFrom, bibTo int) error {
	values := urlValues{
		"bibFrom": bibFrom,
		"bibTo":   bibTo,
	}
	_, err := q.api.get("rawdata/copy", values)
	return err
}

// DistinctValues returns list of unique values existing in raw data
func (q *RawData) DistinctValues() (*model.RawDataDistinctValues, error) {
	bts, err := q.api.get("rawdata/distinctvalues", nil)
	if err != nil {
		return nil, err
	}

	var dest model.RawDataDistinctValues
	err = json.Unmarshal(bts, &dest)
	return &dest, err
}

// AddManual adds a raw data entry
func (q *RawData) AddManual(timingPoint string, bib int, time decimal.Decimal, addT0 bool) error {
	values := urlValues{
		"timingPoint": timingPoint,
		"bib":         bib,
		"time":        time,
		"addT0":       addT0,
	}
	_, err := q.api.get("rawdata/addmanual", values)
	return err
}
