package webapi

// Information contains all api endpoints regarding information
type Information struct {
	api *EventAPI
}

// newParticipant creates a new Participant api endpoint group
func newInformation(api *EventAPI) *Information {
	return &Information{
		api: api,
	}
}

// FrequentNames returns frequent first names have the given prefix
func (q *Information) FrequentNames(prefix string, maxNo int) ([]string, error) {
	values := urlValues{
		"prefix": prefix,
		"maxNo":  maxNo,
	}
	bts, err := q.api.Get("information/frequentnames", values)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}

// GetSex returns gender of the given first name
func (q *Information) GetSex(name string) (string, error) {
	values := urlValues{
		"name": name,
	}
	bts, err := q.api.Get("information/getsex", values)
	if err != nil {
		return "", err
	}
	return string(bts), nil
}

// AddFirstName adds a name to the database of first names
func (q *Information) AddFirstName(name string, sex string) error {
	values := urlValues{
		"name": name,
		"sex":  sex,
	}
	_, err := q.api.Get("information/addfirstname", values)
	return err
}
