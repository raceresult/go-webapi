package webapi

import (
	"encoding/json"

	"github.com/raceresult/go-model/variant"

	"github.com/raceresult/go-model/statistic"
)

// Statistics contains all api endpoints regarding statistics
type Statistics struct {
	api *EventAPI
}

// newStatistics creates a new Statistics api endpoint group
func newStatistics(api *EventAPI) *Statistics {
	return &Statistics{
		api: api,
	}
}

// Names returns the names of all statistics
func (q *Statistics) Names() ([]string, error) {
	bts, err := q.api.Get("statistics/names", nil)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}

// Get returns a statistics
func (q *Statistics) Get(name string) (*statistic.Statistics, error) {
	values := urlValues{
		"name": name,
	}
	bts, err := q.api.Get("statistics/get", values)
	if err != nil {
		return nil, err
	}
	var dest statistic.Statistics
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Save saves a statistics
func (q *Statistics) Save(item *statistic.Statistics) error {
	_, err := q.api.Post("statistics/save", nil, item)
	return err
}

// Delete deletes a statistics
func (q *Statistics) Delete(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.Get("statistics/delete", values)
	return err
}

// Copy creates a copy of a statistics
func (q *Statistics) Copy(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.Get("statistics/copy", values)
	return err
}

// Rename renames a statistics
func (q *Statistics) Rename(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.Get("statistics/rename", values)
	return err
}

// New creates a statistics
func (q *Statistics) New(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.Get("statistics/new", values)
	return err
}

// Create returns a statistics
func (q *Statistics) Create(name string, format string, contests []int) ([]byte, error) {
	values := urlValues{
		"name":    name,
		"format":  format,
		"contest": intSliceToString(contests),
	}
	return q.api.Get("statistics/create", values)
}

// Statistics creates arbitrary statistics
func (q *Statistics) Statistics(row, col, filter, field string, aggregation statistic.Aggregation) ([][]variant.Variant, error) {
	values := urlValues{
		"row":         row,
		"col":         col,
		"filter":      filter,
		"field":       field,
		"aggregation": int(aggregation),
	}
	bts, err := q.api.Get("statistics/statistics", values)
	if err != nil {
		return nil, err
	}

	// parse json
	var arr [][]interface{}
	if err := json.Unmarshal(bts, &arr); err != nil {
		return nil, err
	}

	// convert to variant
	dest := make([][]variant.Variant, 0, len(arr))
	for _, col := range arr {
		vl := make([]variant.Variant, len(col))
		for j, x := range col {
			vl[j] = variant.ToVariant(x)
		}
		dest = append(dest, vl)
	}
	return dest, nil
}
