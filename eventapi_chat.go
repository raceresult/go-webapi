package webapi

import (
	"encoding/json"

	model "github.com/raceresult/go-model"
)

// Chat contains all api endpoints regarding chat
type Chat struct {
	api *EventAPI
}

// newChat creates a new Chat api endpoint group
func newChat(api *EventAPI) *Chat {
	return &Chat{
		api: api,
	}
}

// GetMessages returns the chat messages starting at a certain ID
func (q *Chat) GetMessages(minID int) ([]model.ChatMessage, error) {
	values := urlValues{
		"minID": minID,
	}
	bts, err := q.api.get("chat/getmessages", values)
	if err != nil {
		return nil, err
	}

	var dest []model.ChatMessage
	if err := json.Unmarshal(bts, &dest); err != nil {
		return nil, err
	}
	return dest, nil
}

// GetUsers registers a user and returns a list of all users
func (q *Chat) GetUsers(username string) ([]string, error) {
	values := urlValues{
		"username": username,
	}
	bts, err := q.api.get("chat/getusers", values)
	if err != nil {
		return nil, err
	}
	return parseJsonStringArr(bts)
}

// PostMessage posts a new message
func (q *Chat) PostMessage(username, msg string) error {
	values := urlValues{
		"username": username,
	}
	_, err := q.api.post("chat/postmessage", values, []byte(msg))
	return err
}
