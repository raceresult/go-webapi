package webapi

import (
	"encoding/json"

	"github.com/raceresult/go-model/certificateset"
)

// CertificateSets contains all api endpoints regarding certificate sets
type CertificateSets struct {
	api *EventAPI
}

// newCertificateSets creates a new CertificateSets api endpoint group
func newCertificateSets(api *EventAPI) *CertificateSets {
	return &CertificateSets{
		api: api,
	}
}

// Names returns the names of all certificate sets
func (q *CertificateSets) Names() ([]string, error) {
	bts, err := q.api.get("certificatesets/names", nil)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}

// Get returns a certificate sets
func (q *CertificateSets) Get(name string) (*certificateset.CertificateSet, error) {
	values := urlValues{
		"name": name,
	}
	bts, err := q.api.get("certificatesets/get", values)
	if err != nil {
		return nil, err
	}
	var dest certificateset.CertificateSet
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Save saves a certificate sets
func (q *CertificateSets) Save(item *certificateset.CertificateSet) error {
	_, err := q.api.post("certificatesets/save", nil, item)
	return err
}

// Delete deletes a certificate sets
func (q *CertificateSets) Delete(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("certificatesets/delete", values)
	return err
}

// Copy creates a copy of a certificate sets
func (q *CertificateSets) Copy(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("certificatesets/copy", values)
	return err
}

// Rename renames a certificate sets
func (q *CertificateSets) Rename(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("certificatesets/rename", values)
	return err
}

// New creates a certificate sets
func (q *CertificateSets) New(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("certificatesets/new", values)
	return err
}

// Create returns a certificate set as pdf
func (q *CertificateSets) Create(name string, contests []int, filter string, lang string) ([]byte, error) {
	values := urlValues{
		"name":    name,
		"contest": intSliceToString(contests),
		"filter":  filter,
		"lang":    lang,
	}
	return q.api.get("certificatesets/create", values)
}

// Count returns the count of participants contained in a certificate set
func (q *CertificateSets) Count(name string, contests []int) (int, error) {
	values := urlValues{
		"name":    name,
		"contest": intSliceToString(contests),
	}
	bts, err := q.api.get("certificatesets/count", values)
	if err != nil {
		return 0, err
	}

	var dest int
	if err := json.Unmarshal(bts, &dest); err != nil {
		return 0, err
	}
	return dest, nil
}
