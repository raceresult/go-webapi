package webapi

import (
	"encoding/json"

	"github.com/raceresult/go-model/decimal"
)

// OverwriteValues contains all api endpoints regarding timing point rules
type OverwriteValues struct {
	api *EventAPI
}

// newOverwriteValues creates a new OverwriteValues api endpoint group
func newOverwriteValues(api *EventAPI) *OverwriteValues {
	return &OverwriteValues{
		api: api,
	}
}

// Count returns the number of overwrite values matching the given filters
func (q *OverwriteValues) Count(identifier Identifier, result, contest int, filter string) (int, error) {
	values := urlValues{
		identifier.Name: identifier.Value,
		"result":        result,
		"contest":       contest,
		"filter":        filter,
	}
	bts, err := q.api.get("overwritevalues/count", values)
	if err != nil {
		return 0, err
	}

	var count int
	err = json.Unmarshal(bts, &count)
	return count, err
}

// Delete deletes overwrite values
func (q *OverwriteValues) Delete(identifier Identifier, result, contest int, filter string) error {
	values := urlValues{
		identifier.Name: identifier.Value,
		"result":        result,
		"contest":       contest,
		"filter":        filter,
	}
	_, err := q.api.get("overwritevalues/delete", values)
	return err
}

// Save saves an overwrite value
func (q *OverwriteValues) Save(identifier Identifier, result int, value decimal.Decimal) error {
	values := urlValues{
		identifier.Name: identifier.Value,
		"result":        result,
		"value":         value,
	}
	_, err := q.api.get("overwritevalues/save", values)
	return err
}
