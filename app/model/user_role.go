package model

type UserRole struct {
	UserId	uint	`json:"user_id"`
	RoleId	int		`json:"role_id"`
}

func (u UserRole) TableName() string {
	return "user_role"
}
