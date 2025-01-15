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

// SetInvalidBatch sets multiple raw data entries to valid or invalid
func (q *RawData) SetInvalidBatch(filter string, rdFilter model.RawDataFilter, invalid bool) error {
	btsRDFilter, _ := json.Marshal(rdFilter)
	values := urlValues{
		"filter":   filter,
		"rdFilter": string(btsRDFilter),
		"invalid":  invalid,
	}
	_, err := q.api.get("rawdata/setinvalidbatch", values)
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
func (q *RawData) Delete(bib int, filter string, rdFilter model.RawDataFilter) error {
	btsRDFilter, _ := json.Marshal(rdFilter)
	values := urlValues{
		"bib":      bib,
		"filter":   filter,
		"rdFilter": string(btsRDFilter),
	}
	_, err := q.api.get("rawdata/delete", values)
	return err
}

// Get returns raw data entries
func (q *RawData) Get(bib int, filter string, rdFilter model.RawDataFilter, addFields []string,
	firstRow int, maxRows int, sortBy string) ([]model.RawDataWithAdditionalFields, error) {
	btsRDFilter, _ := json.Marshal(rdFilter)
	values := urlValues{
		"bib":       bib,
		"filter":    filter,
		"rdFilter":  string(btsRDFilter),
		"addFields": addFields,
		"firstRow":  firstRow,
		"maxRows":   maxRows,
		"sortBy":    sortBy,
	}
	bts, err := q.api.get("rawdata/get", values)
	if err != nil {
		return nil, err
	}

	var dest []model.RawDataWithAdditionalFields
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Export returns raw data entries
func (q *RawData) Export(bib int, filter string, rdFilter model.RawDataFilter, fields []string,
	firstRow int, maxRows int, sortBy string) ([][]interface{}, error) {
	btsRDFilter, _ := json.Marshal(rdFilter)
	values := urlValues{
		"bib":      bib,
		"filter":   filter,
		"rdFilter": string(btsRDFilter),
		"fields":   fields,
		"firstRow": firstRow,
		"maxRows":  maxRows,
		"sortBy":   sortBy,
	}
	bts, err := q.api.get("rawdata/export", values)
	if err != nil {
		return nil, err
	}

	var dest [][]interface{}
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Count counts raw data entries matching the given filters
func (q *RawData) Count(bib int, filter string, rdFilter model.RawDataFilter) (int, error) {
	btsRDFilter, _ := json.Marshal(rdFilter)
	values := urlValues{
		"bib":      bib,
		"filter":   filter,
		"rdFilter": string(btsRDFilter),
	}
	bts, err := q.api.get("rawdata/count", values)
	if err != nil {
		return 0, err
	}

	var count int
	err = json.Unmarshal(bts, &count)
	return count, err
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

// Copy copies raw data from one participant to another
func (q *RawData) Copy(bibFrom, bibTo int) error {
	values := urlValues{
		"bibFrom": bibFrom,
		"bibTo":   bibTo,
	}
	_, err := q.api.get("rawdata/copy", values)
	return err
}

// Swap swaps raw data between two participant
func (q *RawData) Swap(bib1, bib2 int) error {
	values := urlValues{
		"bib1": bib1,
		"bib2": bib2,
	}
	_, err := q.api.get("rawdata/swap", values)
	return err
}
