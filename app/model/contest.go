package model

// Contest 比赛表
type Contest struct {
	Id int `json:"id"`
	Title	string	`gorm:"title" json:"title"`
	Brief	string	`gorm:"brief" json:"brief"`
	Award	string	`gorm:"award" json:"award"`
	Condition string	`gorm:"condition" json:"condition"`
	Sponsor   string	`gorm:"sponsor" json:"sponsor"`
	BeginSignup	Time	`gorm:"begin_signup" json:"begin_signup"`
	EndSignup	Time	`gorm:"end_signup" json:"end_signup"`
	BeginSubmit	Time	`gorm:"begin_submit" json:"begin_submit"`
	EndSubmit	Time	`gorm:"end_submit" json:"end_submit"`
	Attribute 	string	`gorm:"attribute" json:"attribute"`
	MinNum	int	`gorm:"min_num;default:1" json:"min_num"`
	MaxNum	int	`gorm:"max_num;default:1" json:"max_num"`
	Manager	string	`gorm:"manager" json:"manager"`
	CreatedAt Time	`gorm:"type:timestamp;autoCreatTime" json:"created_at" swaggerignore:"true"`
	UpdatedAt Time `gorm:"type:timestamp;autoUpdateTime" json:"updated_at" swaggerignore:"true"`

	Signups []ContestSignup	`json:"signups"`
	Judges	[]User	`json:"judges" gorm:"many2many:contest_judge;foreignKey:Id;joinForeignKey:ContestId;References:Id;JoinReferences:TeacherId"`

}

func (c Contest) TableName() string {
	return "contest"
}