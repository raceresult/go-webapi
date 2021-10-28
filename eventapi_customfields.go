package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// CustomFields contains all api endpoints regarding custom fields
type CustomFields struct {
	api *EventAPI
}

// newCustomFields creates a new CustomFields api endpoint group
func newCustomFields(api *EventAPI) *CustomFields {
	return &CustomFields{
		api: api,
	}
}

// Get returns custom fields matching the given filters
func (q *CustomFields) Get() ([]model.CustomField, error) {
	bts, err := q.api.Get("fields/get", nil)
	if err != nil {
		return nil, err
	}

	var dest []model.CustomField
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// GetOne returns the custom field with the given ID
func (q *CustomFields) GetOne(id int) (*model.CustomField, error) {
	values := urlValues{
		"id": id,
	}
	bts, err := q.api.Get("fields/get", values)
	if err != nil {
		return nil, err
	}

	var dest model.CustomField
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Delete deletes a custom field
func (q *CustomFields) Delete(id int) error {
	values := urlValues{
		"id": id,
	}
	_, err := q.api.Get("fields/delete", values)
	return err
}

// Save saves custom fields and returns the IDs
func (q *CustomFields) Save(items []model.CustomField) ([]int, error) {
	bts, err := q.api.Post("fields/save", nil, items)
	if err != nil {
		return nil, err
	}
	return parseJsonIntArr(bts)
}
