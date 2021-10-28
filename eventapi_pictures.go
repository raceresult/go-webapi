package webapi

import (
	"io"
)

// Pictures contains all api endpoints regarding pictures
type Pictures struct {
	api *EventAPI
}

// newPictures creates a new Pictures api endpoint group
func newPictures(api *EventAPI) *Pictures {
	return &Pictures{
		api: api,
	}
}

// Names returns the picture names in the given folder
func (q *Pictures) Names(folder string) ([]string, error) {
	values := urlValues{
		"folder": folder,
	}
	bts, err := q.api.Get("pictures/names", values)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}

// Get returns the picture with the given name
func (q *Pictures) Get(name string) ([]byte, error) {
	values := urlValues{
		"name": name,
	}
	return q.api.Get("pictures/get", values)
}

// Thumbnail returns a thumbnail of the picture with the given name
func (q *Pictures) Thumbnail(name string, maxWidth, maxHeight int) ([]byte, error) {
	values := urlValues{
		"name":      name,
		"maxWidth":  maxWidth,
		"maxHeight": maxHeight,
	}
	return q.api.Get("pictures/thumbnail", values)
}

// Info returns infos about the picture with the given name
func (q *Pictures) Info(name string) (string, error) {
	values := urlValues{
		"name": name,
	}
	bts, err := q.api.Get("pictures/info", values)
	if err != nil {
		return "", err
	}
	return string(bts), nil
}

// Delete deletes the picture with the given name
func (q *Pictures) Delete(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.Get("pictures/delete", values)
	return err
}

// Import saves a picture
func (q *Pictures) Import(folder string, name string, content io.Reader) error {
	values := urlValues{
		"folder": folder,
		"name":   name,
	}
	_, err := q.api.Post("pictures/import", values, content)
	return err
}
