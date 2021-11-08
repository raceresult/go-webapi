package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// WebHooks contains all api endpoints regarding webhooks
type WebHooks struct {
	api *EventAPI
}

// newWebHooks creates a new WebHooks api endpoint group
func newWebHooks(api *EventAPI) *WebHooks {
	return &WebHooks{
		api: api,
	}
}

// Get returns all webhooks
func (q *WebHooks) Get() ([]model.WebHook, error) {
	bts, err := q.api.get("webhooks/get", nil)
	if err != nil {
		return nil, err
	}

	var dest []model.WebHook
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Delete deletes webhooks
func (q *WebHooks) Delete(id int) error {
	values := urlValues{
		"id": id,
	}
	_, err := q.api.get("webhooks/delete", values)
	return err
}

// Save saves webhooks and returns the IDs
func (q *WebHooks) Save(items []model.WebHook) ([]int, error) {
	bts, err := q.api.post("webhooks/save", nil, items)
	if err != nil {
		return nil, err
	}
	return parseJsonIntArr(bts)
}
