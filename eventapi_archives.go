package webapi

import (
	"encoding/json"

	"github.com/raceresult/go-model/archives"
)

// Archives contains all api endpoints regarding archives
type Archives struct {
	api *EventAPI
}

// newArchives creates a new Archives api endpoint group
func newArchives(api *EventAPI) *Archives {
	return &Archives{
		api: api,
	}
}

// CreateNewRegNo creates a new registration number
func (q *Archives) CreateNewRegNo() (int, error) {
	bts, err := q.api.get("archives/createnewregno", nil)
	if err != nil {
		return 0, err
	}

	var dest int
	if err := json.Unmarshal(bts, &dest); err != nil {
		return 0, err
	}
	return dest, nil
}

// GetMatches returns matching archive entries
func (q *Archives) GetMatches(prefix string, maxNumber int) ([]archives.Match, error) {
	values := urlValues{
		"prefix":    prefix,
		"maxNumber": maxNumber,
	}
	bts, err := q.api.get("archives/getmatches", values)
	if err != nil {
		return nil, err
	}

	var dest []archives.Match
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// GetEntry returns a single archives entry
func (q *Archives) GetEntry(id int, regNo string) (*archives.Participant, error) {
	values := urlValues{
		"id":    id,
		"regNo": regNo,
	}
	bts, err := q.api.get("archives/getentry", values)
	if err != nil {
		return nil, err
	}

	var dest archives.Participant
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return &dest, nil
}

// GetParticipations returns the participations for the given bib
func (q *Archives) GetParticipations(bib int) ([]archives.ParticipationExt, error) {
	values := urlValues{
		"bib": bib,
	}
	bts, err := q.api.get("archives/getparticipations", values)
	if err != nil {
		return nil, err
	}

	var dest []archives.ParticipationExt
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// Download returns the entire archives file
func (q *Archives) Download() ([]byte, error) {
	return q.api.get("archives/download", nil)
}

// Remove deletes the entire archive from the event
func (q *Archives) Remove() error {
	_, err := q.api.get("archives/remove", nil)
	return err
}

// Create creates a new archives file
func (q *Archives) Create() error {
	_, err := q.api.get("archives/create", nil)
	return err
}

// Write writes the data of the current event into the archive
func (q *Archives) Write() error {
	_, err := q.api.get("archives/write", nil)
	return err
}

// Import imports an archive into the event file
func (q *Archives) Import(bts []byte) error {
	_, err := q.api.post("archives/import", nil, bts)
	return err
}
