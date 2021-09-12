package models

type TeamsCreateRequest struct {
	Name    string   `json:"name"`
	Type    int      `json:"type"` //1 - private
	Members []string `json:"members,omitempty"`
	Room    struct {
		ReadOnly bool `json:"readOnly"`
	} `json:"room,omitempty"`
}
type TeamsAddMembersRequest struct {
	TeamID  string `json:"teamId"`
	Members []struct {
		UserID string   `json:"userId"`
		Roles  []string `json:"roles"`
	} `json:"members"`
}
