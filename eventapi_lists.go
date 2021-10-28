package webapi

import (
	"encoding/json"

	"github.com/raceresult/go-model/list"
)

// Lists contains all api endpoints regarding lists
type Lists struct {
	api *EventAPI
}

// newLists creates a new Lists api endpoint group
func newLists(api *EventAPI) *Lists {
	return &Lists{
		api: api,
	}
}

// Names returns the names of all lists
func (q *Lists) Names() ([]string, error) {
	bts, err := q.api.get("lists/names", nil)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}

// Delete deletes a list
func (q *Lists) Delete(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("lists/delete", values)
	return err
}

// Copy creates a copy of a list
func (q *Lists) Copy(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("lists/copy", values)
	return err
}

// Rename renames a list
func (q *Lists) Rename(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("lists/rename", values)
	return err
}

// New creates a new list
func (q *Lists) New(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("lists/new", values)
	return err
}

// Get returns the settings of a list
func (q *Lists) Get(name string, noTranslate bool, lang string) (*list.List, error) {
	values := urlValues{
		"name":        name,
		"noTranslate": noTranslate,
		"lang":        lang,
	}
	bts, err := q.api.get("lists/get", values)
	if err != nil {
		return nil, err
	}
	var dest *list.List
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Save saves a list
func (q *Lists) Save(item *list.List) error {
	_, err := q.api.post("lists/save", nil, item)
	return err
}

// ParticipantsNotActivated returns the number of participants in the list which are not activated
func (q *Lists) ParticipantsNotActivated(name string, contests []int, onlyWithUnderscores bool) (int, error) {
	values := urlValues{
		"name":                name,
		"contest":             intSliceToString(contests),
		"onlyWithUnderscores": onlyWithUnderscores,
	}
	bts, err := q.api.get("lists/participantsnotactivated", values)
	if err != nil {
		return 0, err
	}

	var dest int
	if err := json.Unmarshal(bts, &dest); err != nil {
		return 0, err
	}
	return dest, nil
}
