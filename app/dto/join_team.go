package dto

import "lab-reverse/app/model"

type JoinTeam struct {
	TeamId	int			`json:"team_id"`
	Acked	string		`json:"acked"`
	Team	model.Team	`json:"team"`
}

func ToJoinTeam(members []model.TeamMember) []JoinTeam {
	var res []JoinTeam
	for _, val := range members {
		res = append(res, JoinTeam{
			TeamId: val.TeamId,
			Acked: val.Acked,
			Team: *val.Team,
		})
	}
	return res
}
