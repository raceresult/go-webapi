package webapi

import (
	"strings"

	model "github.com/raceresult/go-model"
)

// ChipFile contains all api endpoints regarding chip file
type ChipFile struct {
	api *EventAPI
}

// newChipFile creates a new ChipFile api endpoint group
func newChipFile(api *EventAPI) *ChipFile {
	return &ChipFile{
		api: api,
	}
}

// Get returns the entire chip file
func (q *ChipFile) Get() ([]model.ChipFileEntry, error) {
	bts, err := q.api.Get("chipfile/get", nil)
	if err != nil {
		return nil, err
	}

	var dest []model.ChipFileEntry
	for _, v := range strings.Split(string(bts), "\r\n") {
		arr := strings.Split(v, ";")
		if len(arr) != 2 {
			continue
		}
		dest = append(dest, model.ChipFileEntry{
			Transponder:    arr[0],
			Identification: arr[1],
		})
	}
	return dest, nil
}

// Save saves a new chip file
func (q *ChipFile) Save(items []model.ChipFileEntry) error {
	// create data
	arr := make([]string, 0, len(items))
	for _, e := range items {
		arr = append(arr, e.Transponder+";"+e.Identification)
	}
	bts := []byte(strings.Join(arr, "\r\n"))

	_, err := q.api.Post("chipfile/save", nil, bts)
	return err
}

// Clear clears the entire chip file
func (q *ChipFile) Clear() error {
	_, err := q.api.Get("chipfile/clear", nil)
	return err
}
