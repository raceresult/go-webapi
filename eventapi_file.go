package webapi

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	model "github.com/raceresult/go-model"
)

// File contains all api endpoints regarding the event file
type File struct {
	api *EventAPI
}

// newFile creates a new File api endpoint group
func newFile(api *EventAPI) *File {
	return &File{
		api: api,
	}
}

// Activate activates participants and returns the number of activated records
func (q *File) Activate(bib int, filter string, maxActivations int) (int, error) {
	values := urlValues{
		"bib":            bib,
		"filter":         filter,
		"maxActivations": maxActivations,
	}
	bts, err := q.api.Get("file/activate", values)
	if err != nil {
		return 0, err
	}

	var dest int
	if err := json.Unmarshal(bts, &dest); err != nil {
		return 0, err
	}
	return dest, nil
}

// NotActivated returns the number of participants which are not activated
func (q *File) NotActivated(filter string) (int, error) {
	values := urlValues{
		"filter": filter,
	}
	bts, err := q.api.Get("file/notactivated", values)
	if err != nil {
		return 0, err
	}

	var dest int
	if err := json.Unmarshal(bts, &dest); err != nil {
		return 0, err
	}
	return dest, nil
}

// SESVersion returns the version of the Sports Event Server
func (q *File) SESVersion() (*model.Version, error) {
	bts, err := q.api.Get("file/sesversion", nil)
	if err != nil {
		return nil, err
	}

	var dest model.Version
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// CheckExpression parses an expression
func (q *File) CheckExpression(expressions string, returnTree bool) (string, error) {
	values := urlValues{
		"expressions": expressions,
		"returnTree":  returnTree,
	}
	bts, err := q.api.Get("file/checkexpression", values)
	if err != nil {
		return "", err
	}
	return string(bts), nil
}

// GetFile returns a copy of the entire event file
func (q *File) GetFile() ([]byte, error) {
	bts, err := q.api.Get("file/getfile", nil)
	if err != nil {
		return nil, err
	}
	return bts, nil
}

// ModJobID returns the modjobid of the file
func (q *File) ModJobID() (int, error) {
	bts, err := q.api.Get("file/modjobid", nil)
	if err != nil {
		return 0, err
	}
	var dest int
	if err := json.Unmarshal(bts, &dest); err != nil {
		return 0, err
	}
	return dest, nil
}

// ModJobIDs returns the normal modjobid and the settings modjobid of the file
func (q *File) ModJobIDs() (int, int, error) {
	bts, err := q.api.Get("file/modjobids", nil)
	if err != nil {
		return 0, 0, err
	}
	arr := strings.Split(string(bts), ";")
	if len(arr) != 2 {
		return 0, 0, errors.New("response invalid")
	}
	mid, _ := strconv.Atoi(arr[0])
	midSettings, _ := strconv.Atoi(arr[1])
	return mid, midSettings, nil
}

// Filename returns the filename of the event file
func (q *File) Filename() (string, error) {
	bts, err := q.api.Get("file/filename", nil)
	if err != nil {
		return "", err
	}
	return string(bts), nil
}

// Owner returns the user ID of the owner of the event.
// Available for online server only.
func (q *File) Owner() (int, error) {
	bts, err := q.api.Get("file/owner", nil)
	if err != nil {
		return 0, err
	}
	var dest int
	if err := json.Unmarshal(bts, &dest); err != nil {
		return 0, err
	}
	return dest, nil
}

// IsOwner returns true if the user owns the event.
// Available for online server only.
func (q *File) IsOwner() (bool, error) {
	bts, err := q.api.Get("file/isowner", nil)
	if err != nil {
		return false, err
	}
	var dest bool
	err = json.Unmarshal(bts, &dest)
	return dest, err
}

// Rights returns the user rights code this user has for this event.
// Available for online server only.
func (q *File) Rights() (string, error) {
	bts, err := q.api.Get("file/rights", nil)
	if err != nil {
		return "", err
	}
	return string(bts), nil
}
