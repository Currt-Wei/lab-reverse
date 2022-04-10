package model

type Team struct {
	Id 				int `json:"id" gorm:"primaryKey"`
	Name 			string `json:"name"`
	Declaration 	string `json:"declaration"`
	Email 			string `json:"email"`
	Number 			int	`json:"number"`
	Leader 			string `json:"leader"`
	AckNumber 		int `json:"ack_number" gorm:"default:1"`
	Status 			string `json:"status" gorm:"default:invalid"`

	Members 		[]TeamMember	`json:"members" gorm:"foreignKey:TeamId"`
	ContestSignups  []ContestSignup	`gorm:"polymorphic:Target;polymorphicValue:team"`
}

func (t Team) TableName() string {
	return "team"
}