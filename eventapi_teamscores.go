package webapi

import (
	"encoding/json"
	"errors"

	model "github.com/raceresult/go-model"
)

// TeamScores contains all api endpoints regarding team scores
type TeamScores struct {
	api *EventAPI
}

// newTeamScores creates a new TeamScores api endpoint group
func newTeamScores(api *EventAPI) *TeamScores {
	return &TeamScores{
		api: api,
	}
}

// Get returns all team scores
func (q *TeamScores) Get() ([]model.TeamScore, error) {
	bts, err := q.api.Get("teamscores/get", nil)
	if err != nil {
		return nil, err
	}

	var dest []model.TeamScore
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// GetOne returns the team scores with the given ID
func (q *TeamScores) GetOne(id int) (*model.TeamScore, error) {
	values := urlValues{
		"id": id,
	}
	bts, err := q.api.Get("teamscores/get", values)
	if err != nil {
		return nil, err
	}

	var dest []model.TeamScore
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	if len(dest) == 0 {
		return nil, errors.New("team score not found")
	}
	return &dest[0], nil
}

// Delete deletes team scores
func (q *TeamScores) Delete(id int) error {
	values := urlValues{
		"id": id,
	}
	_, err := q.api.Get("teamscores/delete", values)
	return err
}

// Save saves a team score
func (q *TeamScores) Save(item model.TeamScore) error {
	_, err := q.api.Post("teamscores/save", nil, item)
	return err
}
