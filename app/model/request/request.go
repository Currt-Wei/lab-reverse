package request

type JudgeAdd struct {
	ContestId	int		`json:"contest_id" form:"contest_id"`
	Judges		[]uint	`json:"judges" form:"judges"`
}

type AuthorityRequest struct {
	RoleId			int                	`json:"role_id" form:"role_id"`
	AuthorityInfo	[]AuthorityInfo 	`json:"authority_info" form:"authority_info"`
}

// AuthorityInfo 返回casbin表中的path和method数据
type AuthorityInfo struct {
	Path   string `json:"path"`   // 路径
	Method string `json:"method"` // 方法
}

type UserRoles struct {
	UserId	uint	`json:"user_id"`
	RoleIds	[]int	`json:"role_ids"`
}

// MenuInfo 前端与后端交互的menu的信息
type MenuInfo struct {
	Path 		string	`json:"path"`
	Description string	`json:"description"`
}

// MenuRequest 提交的菜单的结构体
type MenuRequest struct {
	RoleId		int         `json:"role_id" form:"role_id"`
	MenuInfo	[]MenuInfo 	`json:"menu_info" form:"menu_info"`
}