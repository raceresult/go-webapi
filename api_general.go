package webapi

// General contains all api endpoints regarding general functions
type General struct {
	api *API
}

// newGeneral creates a new General api endpoint group
func newGeneral(api *API) *General {
	return &General{
		api: api,
	}
}

// Fonts returns a list of supported fonts
func (q *General) Fonts() ([]string, error) {
	bts, err := q.api.get("", "fonts", nil)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}

// AppVersion returns the webserver version
func (q *General) AppVersion() (string, error) {
	bts, err := q.api.get("", "appversion", nil)
	if err != nil {
		return "", err
	}
	return string(bts), nil
}

// Translate translates fields/expressions
func (q *General) Translate(items []string, fromEnglish bool, lang string) ([]string, error) {
	values := urlValues{
		"fromEnglish": fromEnglish,
		"lang":        lang,
	}
	bts, err := q.api.post("", "translate2", values, "", items)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}

// GetLangItem returns a translation text item
func (q *General) GetLangItem(name, lang string) (string, error) {
	values := urlValues{
		"name": name,
		"lang": lang,
	}
	bts, err := q.api.get("", "getlangitem", values)
	if err != nil {
		return "", err
	}
	return string(bts), nil
}
