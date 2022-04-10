package model

// Role 角色
type Role struct {
	Id 				int 		`json:"id" gorm:"primaryKey"`
	Name 			string 		`json:"name"`
	Remark			string		`json:"remark"`
}

func (r Role) TableName() string {
	return "role"
}
