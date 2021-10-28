package webapi

import "encoding/json"

// Synchronization contains all api endpoints regarding synchronization
type Synchronization struct {
	api *EventAPI
}

// newSynchronization creates a new Synchronization api endpoint group
func newSynchronization(api *EventAPI) *Synchronization {
	return &Synchronization{
		api: api,
	}
}

// IsCheckedOut returns true if the file has status check out.
// Available for online server only.
func (q *Synchronization) IsCheckedOut() (bool, error) {
	bts, err := q.api.Get("synchronization/isCheckedOut", nil)
	if err != nil {
		return false, err
	}

	var dest bool
	if err := json.Unmarshal(bts, &dest); err != nil {
		return false, err
	}
	return dest, nil
}

// SetCheckedIn sets the status of the event to checkedIn.
// Available for online server only. Use to set lost file back to checked in.
func (q *Synchronization) SetCheckedIn() error {
	_, err := q.api.Get("synchronization/setCheckedIn", nil)
	return err
}
