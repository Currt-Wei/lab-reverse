package model

type ContestReviewLog struct {
	Id			int		`json:"id"`
	SignupId	int		`json:"signup_id"`
	Score		int		`json:"score"`
	Comment		string `json:"comment"`
	CreatedAt	Time	`json:"created_at" gorm:"autoUpdateTime"`
	Reviewer	string	`json:"reviewer"`
}

func (c ContestReviewLog) TableName() string {
	return "contest_review_log"
}