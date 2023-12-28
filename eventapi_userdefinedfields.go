package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// UserDefinedFields contains all api endpoints regarding user defined fields
type UserDefinedFields struct {
	api *EventAPI
}

// newUserDefinedFields creates a new UserDefinedFields api endpoint group
func newUserDefinedFields(api *EventAPI) *UserDefinedFields {
	return &UserDefinedFields{
		api: api,
	}
}

// Get returns all user defined fields
func (q *UserDefinedFields) Get() ([]model.UserDefinedField, error) {
	bts, err := q.api.get("userdefinedfields/get", nil)
	if err != nil {
		return nil, err
	}

	var dest []model.UserDefinedField
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Set overwrites all user defined fields
func (q *UserDefinedFields) Set(items []model.UserDefinedField) error {
	_, err := q.api.post("userdefinedfields/set", nil, items)
	return err
}
