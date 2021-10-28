package webapi

import (
	"encoding/json"

	"github.com/raceresult/go-model/emailtemplate"
	"github.com/raceresult/go-model/label"
)

// EmailTemplates contains all api endpoints regarding email templates
type EmailTemplates struct {
	api *EventAPI
}

// newEmailTemplates creates a new EmailTemplates api endpoint group
func newEmailTemplates(api *EventAPI) *EmailTemplates {
	return &EmailTemplates{
		api: api,
	}
}

// Names returns the names of all email templates
func (q *EmailTemplates) Names() ([]string, error) {
	bts, err := q.api.get("emailtemplates/names", nil)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}

// Get returns an email template
func (q *EmailTemplates) Get(name string) (*emailtemplate.EmailTemplate, error) {
	values := urlValues{
		"name": name,
	}
	bts, err := q.api.get("emailtemplates/get", values)
	if err != nil {
		return nil, err
	}
	var dest emailtemplate.EmailTemplate
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Save saves an email template
func (q *EmailTemplates) Save(item *label.Label) error {
	_, err := q.api.post("emailtemplates/save", nil, item)
	return err
}

// Delete deletes an email template
func (q *EmailTemplates) Delete(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("emailtemplates/delete", values)
	return err
}

// Copy creates a copy of an email template
func (q *EmailTemplates) Copy(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("emailtemplates/copy", values)
	return err
}

// Rename renames an email template
func (q *EmailTemplates) Rename(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("emailtemplates/rename", values)
	return err
}

// New creates an email template
func (q *EmailTemplates) New(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("emailtemplates/new", values)
	return err
}

// Preview generates previews of the email
func (q *EmailTemplates) Preview(name string, filter string, lang string) ([]emailtemplate.Preview, error) {
	values := urlValues{
		"name":   name,
		"filter": filter,
		"lang":   lang,
	}
	bts, err := q.api.get("emailtemplates/preview", values)
	if err != nil {
		return nil, err
	}
	var dest []emailtemplate.Preview
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// SendPreview sends a pre-generated preview
func (q *EmailTemplates) SendPreview(name string, lang string, preview *emailtemplate.Preview) error {
	values := urlValues{
		"name": name,
		"lang": lang,
	}
	_, err := q.api.post("emailtemplates/sendpreview", values, preview)
	return err
}

// Send generates the preview and directly sends them
func (q *EmailTemplates) Send(name string, filter string, lang string) error {
	values := urlValues{
		"name":   name,
		"filter": filter,
		"lang":   lang,
	}
	_, err := q.api.get("emailtemplates/send", values)
	return err
}
