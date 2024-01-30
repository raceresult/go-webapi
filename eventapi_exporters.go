package webapi

import (
	"encoding/json"
	"errors"

	model "github.com/raceresult/go-model"
)

// Exporters contains all api endpoints regarding exporters
type Exporters struct {
	api *EventAPI
}

// newExporters creates a new Exporters api endpoint group
func newExporters(api *EventAPI) *Exporters {
	return &Exporters{
		api: api,
	}
}

// Get returns all exporters
func (q *Exporters) Get() ([]model.Exporter, error) {
	bts, err := q.api.get("exporters/get", nil)
	if err != nil {
		return nil, err
	}

	var dest []model.Exporter
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// GetOne returns the contest with the given ID
func (q *Exporters) GetOne(id int) (*model.Exporter, error) {
	values := urlValues{
		"id": id,
	}
	bts, err := q.api.get("exporters/get", values)
	if err != nil {
		return nil, err
	}

	var dest []model.Exporter
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("exporter not found")
	}
	return &dest[0], nil
}

// Delete deletes exporters
func (q *Exporters) Delete(id int) error {
	values := urlValues{
		"id": id,
	}
	_, err := q.api.get("exporters/delete", values)
	return err
}

// Save saves a exporter and returns the ID
func (q *Exporters) Save(item model.Exporter) (int, error) {
	bts, err := q.api.post("exporters/save", nil, item)
	if err != nil {
		return 0, err
	}

	var id int
	if err := json.Unmarshal(bts, &id); err != nil {
		return 0, err
	}
	return id, nil
}
