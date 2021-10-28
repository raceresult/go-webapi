package webapi

// Table contains all api endpoints regarding table
type Table struct {
	api *EventAPI
}

// newTable creates a new Table api endpoint group
func newTable(api *EventAPI) *Table {
	return &Table{
		api: api,
	}
}
