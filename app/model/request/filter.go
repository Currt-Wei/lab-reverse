package request

// UserFilter 查找用户的筛选信息
type UserFilter struct {
	Account string	`form:"account"`
	Name	string	`form:"name"`
	College	string	`form:"college"`
	Degree	string	`form:"degree"`
	Grade	string	`form:"grade"`
	Identity string	`form:"identity"`
}

// TeamFilter 查找队伍的筛选信息
type TeamFilter struct {
	Leader	string	`form:"leader"`	// 队长
	Name	string	`form:"name"`	// 队名
	Status	string	`form:"status"`	// 状态
}

// ContestFilter 比赛的筛选条件
type ContestFilter struct {
	Title	string	`form:"title"`
	Status	string	`form:"status"`
	Attribute	string	`form:"attribute"`
	Manager 	string	`form:"manager"`
}