package model

// ContestSignup 参赛的人/队伍
type ContestSignup struct {
	Id 			int `json:"id"`
	ContestId 	int `json:"contest_id" form:"contest_id"`
	TargetID 	int	`json:"target_id" form:"target_id"`	// 用户的id(注意不是学号)或者是团队的id
	TargetType 	string `json:"target_type" form:"target_type"`
	WorkLink 	string `json:"work_link" form:"work_link"`
	Score 		int	`json:"score"`
	Comment		string	`json:"comment"`
	Status 		string	`json:"status"`

	Contest		*Contest	`json:"contest"`
	User		*User		`json:"user" gorm:"foreignKey:TargetID"`
	Team		*Team		`json:"team" gorm:"foreignKey:TargetID"`
	ReviewLogs 	[]ContestReviewLog	`json:"review_log" gorm:"foreignKey:SignupId"`
}

func (c ContestSignup) TableName() string {
	return "contest_signup"
}