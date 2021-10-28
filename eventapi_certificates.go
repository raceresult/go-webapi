package webapi

// Certificates contains all api endpoints regarding certificates
type Certificates struct {
	api *EventAPI
}

// newCertificates creates a new Certificates api endpoint group
func newCertificates(api *EventAPI) *Certificates {
	return &Certificates{
		api: api,
	}
}

// Names returns the names of all certificates
func (q *Certificates) Names() ([]string, error) {
	bts, err := q.api.get("certificates/names", nil)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}

// Delete deletes a certificate
func (q *Certificates) Delete(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("certificates/delete", values)
	return err
}

// Copy creates a copy of a certificate
func (q *Certificates) Copy(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("certificates/copy", values)
	return err
}

// Rename renames a certificate
func (q *Certificates) Rename(name, newName string) error {
	values := urlValues{
		"name":    name,
		"newName": newName,
	}
	_, err := q.api.get("certificates/rename", values)
	return err
}

// New creates a certificate
func (q *Certificates) New(name string) error {
	values := urlValues{
		"name": name,
	}
	_, err := q.api.get("certificates/new", values)
	return err
}

// Thumbnail returns a thumbnail of a certificate
func (q *Certificates) Thumbnail(name string, maxWidth, maxHeight int) ([]byte, error) {
	values := urlValues{
		"name":      name,
		"maxWidth":  maxWidth,
		"maxHeight": maxHeight,
	}
	return q.api.get("certificates/thumbnail", values)
}

// PreviewJPG returns a preview of a certificate as jpg
func (q *Certificates) PreviewJPG(name string, page int, dpi int, lang string) ([]byte, error) {
	values := urlValues{
		"name": name,
		"page": page,
		"dpi":  dpi,
		"lang": lang,
	}
	return q.api.get("certificates/previewJPG", values)
}

// CreatePDF returns a certificate as pdf
func (q *Certificates) CreatePDF(name string, page int, bib int, lang string) ([]byte, error) {
	values := urlValues{
		"name":   name,
		"page":   page,
		"bib":    bib,
		"lang":   lang,
		"format": "pdf",
	}
	return q.api.get("certificates/create", values)
}

// CreateJPG returns a certificate as jpg
func (q *Certificates) CreateJPG(name string, page int, bib int, dpi int, lang string) ([]byte, error) {
	values := urlValues{
		"name":   name,
		"page":   page,
		"bib":    bib,
		"dpi":    dpi,
		"lang":   lang,
		"format": "jpg",
	}
	return q.api.get("certificates/create", values)
}
