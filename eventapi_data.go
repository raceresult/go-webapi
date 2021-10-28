package webapi

import (
	"encoding/json"

	"github.com/raceresult/go-model/variant"
)

// Data contains all api endpoints regarding data
type Data struct {
	api *EventAPI
}

// newData creates a new Data api endpoint group
func newData(api *EventAPI) *Data {
	return &Data{
		api: api,
	}
}

// Count returns the count of records matching the given filters
func (q *Data) Count(filter string) (int, error) {
	values := urlValues{
		"filter": filter,
	}
	bts, err := q.api.get("data/count", values)
	if err != nil {
		return 0, err
	}
	var dest int
	if err := json.Unmarshal(bts, &dest); err != nil {
		return 0, err
	}
	return dest, nil
}

// List returns arbitrary records
func (q *Data) List(fields []string, filter string, sort []string, limitFrom, limitTo int, groups []string,
	multiplierField string, selectorResult string) ([][]variant.Variant, error) {

	values := urlValues{
		"fields":          fields,
		"filter":          filter,
		"sort":            sort,
		"limitFrom":       limitFrom,
		"limitTo":         limitTo,
		"groups":          groups,
		"multiplierField": multiplierField,
		"selectorResult":  selectorResult,
		"listFormat":      "JSON",
	}
	bts, err := q.api.get("data/list", values)
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

// Transformation creates min/max/sum/count/avg statistic
func (q *Data) Transformation(colField string, rowFields []string, filter string, field string, mode int,
	sortByValue bool) ([][]variant.Variant, error) {

	values := urlValues{
		"colField":    colField,
		"rowFields":   rowFields,
		"filter":      filter,
		"field":       field,
		"mode":        mode,
		"sortByValue": sortByValue,
	}
	bts, err := q.api.get("data/transformation", values)
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
