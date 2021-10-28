package webapi

// Dependencies contains all api endpoints regarding dependencies
type Dependencies struct {
	api *EventAPI
}

// newDependencies creates a new Dependencies api endpoint group
func newDependencies(api *EventAPI) *Dependencies {
	return &Dependencies{
		api: api,
	}
}

// Show returns a string containing the dependency tree
func (q *Dependencies) Show() (string, error) {
	bts, err := q.api.Get("dependencies/show", nil)
	if err != nil {
		return "", err
	}
	return string(bts), nil
}

// CircularReferences returns a string containing the list of circular references
func (q *Dependencies) CircularReferences() (string, error) {
	bts, err := q.api.Get("dependencies/circularreferences", nil)
	if err != nil {
		return "", err
	}
	return string(bts), nil
}
