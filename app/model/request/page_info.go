package request

type PageInfo struct {
	Limit int	`form:"limit,default=10"` // 每页多少条
	Page int	`form:"page,default=1"`   // 第几页
}
