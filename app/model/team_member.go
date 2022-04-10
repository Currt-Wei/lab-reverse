package model

type TeamMember struct {
	Id 		int `json:"id"`
	Account string `json:"account"`
	TeamId 	int	`json:"team_id"`
	Name 	string `json:"name"`
	Email 	string `json:"email"`
	Acked 	string	`json:"acked" gorm:"default:no"`

	Team	*Team	`json:"team"`
}

func (t TeamMember) TableName() string {
	return "team_member"
}