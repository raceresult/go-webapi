package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// Backup contains all api endpoints regarding backup
type Backup struct {
	api *EventAPI
}

// newBackup creates a new Backup api endpoint group
func newBackup(api *EventAPI) *Backup {
	return &Backup{
		api: api,
	}
}

// Active returns true if backup is running
func (q *Backup) Active() (bool, error) {
	bts, err := q.api.get("backup/active", nil)
	if err != nil {
		return false, err
	}

	var b bool
	err = json.Unmarshal(bts, &b)
	return b, err
}

// Start starts the backup process
func (q *Backup) Start(hostname, filename string) error {
	values := urlValues{
		"hostname": hostname,
		"filename": filename,
	}
	_, err := q.api.get("backup/start", values)
	return err
}

// Restart restarts the backup process with previous settings
func (q *Backup) Restart() error {
	_, err := q.api.get("backup/restart", nil)
	return err
}

// Stop stops the backup process
func (q *Backup) Stop() error {
	_, err := q.api.get("backup/stop", nil)
	return err
}

// Info returns statistics about the backup process
func (q *Backup) Info() (model.ForwardingInfo, error) {
	bts, err := q.api.get("backup/info", nil)
	if err != nil {
		return model.ForwardingInfo{}, err
	}

	var dest model.ForwardingInfo
	err = json.Unmarshal(bts, &dest)
	return dest, err
}
