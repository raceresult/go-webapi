package webapi

import (
	"encoding/json"
	"strings"

	"github.com/raceresult/go-model/variant"
)

// Settings contains all api endpoints regarding settings
type Settings struct {
	api *EventAPI
}

// newSettings creates a new Settings api endpoint group
func newSettings(api *EventAPI) *Settings {
	return &Settings{
		api: api,
	}
}

// Get returns setting values
func (q *Settings) Get(names ...string) (variant.VariantMap, error) {
	values := urlValues{}
	switch len(names) {
	case 0:
		return variant.VariantMap{}, nil
	case 1:
		values["name"] = names[0]
	default:
		values["names"] = strings.Join(names, ",")
	}
	bts, err := q.api.Get("settings/getsettings", values)
	if err != nil {
		return nil, err
	}

	var dest variant.VariantMap
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// GetValue returns a single setting value
func (q *Settings) GetValue(name string) (variant.Variant, error) {
	vv, err := q.Get(name)
	if err != nil {
		return nil, err
	}
	return vv[name], nil
}

// Save saves setting values
func (q *Settings) Save(values variant.VariantMap) error {
	_, err := q.api.Post("settings/savesettings", nil, values)
	return err
}

// SaveValue saves a single setting value
func (q *Settings) SaveValue(name string, value variant.Variant) error {
	return q.Save(variant.VariantMap{
		name: value,
	})
}

// Delete deletes a single setting which can be linked to a contest/result
func (q *Settings) Delete(name string, contest int, result int) error {
	values := urlValues{
		"name":    name,
		"contest": contest,
		"result":  result,
	}
	_, err := q.api.Get("settings/delete", values)
	return err
}

// NamesByPrefix returns the names of settings matching the given prefix
func (q *Settings) NamesByPrefix(prefix string) ([]string, error) {
	values := urlValues{
		"prefix": prefix,
	}
	bts, err := q.api.Get("settings/settingnamesbyprefix", values)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}
