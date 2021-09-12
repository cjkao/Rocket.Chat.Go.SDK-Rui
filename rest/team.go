package rest

import (
	"bytes"
	"encoding/json"
	"log"
	"time"

	"github.com/cjkao/Rocket.Chat.Go.SDK/models"
)

type TeamsCreateResponse struct {
	Status
	Team struct {
		ID        string    `json:"_id"`
		Name      string    `json:"name"`
		Type      int       `json:"type"`
		CreatedAt time.Time `json:"createdAt"`
		CreatedBy struct {
			ID       string `json:"_id"`
			Username string `json:"username"`
		} `json:"createdBy"`
		UpdatedAt time.Time `json:"_updatedAt"`
		RoomID    string    `json:"roomId"`
	} `json:"team"`
}
type GeneralResponse struct {
	Status
}

func (c *Client) TeamsCreate(req *models.TeamsCreateRequest) (*TeamsCreateResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(TeamsCreateResponse)
	err = c.Post("teams.create", bytes.NewBuffer(body), response)
	if err != nil {
		log.Fatal(err)
	}
	return response, nil
}
func (c *Client) TeamsAddMembers(req *models.TeamsAddMembersRequest) (*GeneralResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	response := new(GeneralResponse)
	err = c.Post("teams.create", bytes.NewBuffer(body), response)
	if err != nil {
		log.Fatal(err)
	}
	return response, nil
}
