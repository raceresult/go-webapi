package webapi

import (
	"encoding/json"
	model "github.com/raceresult/go-model"
	"github.com/raceresult/go-model/datetime"
)

// AgeGroups contains all api endpoints regarding age groups
type AgeGroups struct {
	api *EventAPI
}

// newAgeGroups creates a new AgeGroups api endpoint group
func newAgeGroups(api *EventAPI) *AgeGroups {
	return &AgeGroups{
		api: api,
	}
}

// PDF returns a PDF with all age groups
func (q *AgeGroups) PDF() ([]byte, error) {
	return q.api.get("agegroups/pdf", nil)
}

// Get returns age groups matching the given filters
func (q *AgeGroups) Get(contest int, set int, name string) ([]model.AgeGroup, error) {
	values := urlValues{
		"contest": contest,
		"set":     set,
		"name":    name,
	}
	bts, err := q.api.get("agegroups/get", values)
	if err != nil {
		return nil, err
	}

	var dest []model.AgeGroup
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Delete deletes age groups
func (q *AgeGroups) Delete(id int, contest int, set int) error {
	values := urlValues{
		"id":      id,
		"contest": contest,
		"set":     set,
	}
	_, err := q.api.get("agegroups/delete", values)
	return err
}

// Save saves age groups and returns the age group IDs
func (q *AgeGroups) Save(items []model.AgeGroup) ([]int, error) {
	bts, err := q.api.post("agegroups/save", nil, items)
	if err != nil {
		return nil, err
	}
	return parseJsonIntArr(bts)
}

// Generate generates new age groups from templates
func (q *AgeGroups) Generate(mode string, contest int, set int, ageBase bool, date datetime.DateTime, lang string) (
	[]model.AgeGroup, error) {

	values := urlValues{
		"mode":    mode,
		"contest": contest,
		"set":     set,
		"ageBase": ageBase,
		"date":    date,
		"lang":    lang,
	}
	bts, err := q.api.get("agegroups/generate", values)
	if err != nil {
		return nil, err
	}

	var dest []model.AgeGroup
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Reassign reassigns age groups
func (q *AgeGroups) Reassign(contest int, identifier Identifier, set int, addOnly bool) error {
	values := urlValues{
		"contest":       contest,
		identifier.Name: identifier.Value,
		"set":           set,
		"addOnly":       addOnly,
	}
	_, err := q.api.get("agegroups/reassign", values)
	return err
}
