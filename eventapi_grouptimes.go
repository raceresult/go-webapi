package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// GroupTimes contains all api endpoints regarding group times (start times+ finish time limit)
type GroupTimes struct {
	api *EventAPI
}

// newGroupTimes creates a new GroupTimes api endpoint group
func newGroupTimes(api *EventAPI) *GroupTimes {
	return &GroupTimes{
		api: api,
	}
}

// Get returns all contests
func (q *GroupTimes) Get(ttype string) (*model.GroupTimes, error) {
	values := urlValues{
		"type": ttype,
	}
	bts, err := q.api.get("grouptimes/get", values)
	if err != nil {
		return nil, err
	}

	var dest model.GroupTimes
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// Save saves a contest and returns the ID
func (q *GroupTimes) Save(ttype string, item model.GroupTimes) error {
	values := urlValues{
		"type": ttype,
	}
	_, err := q.api.post("grouptimes/save", values, item)
	return err
}
