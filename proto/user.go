package proto

type UserActivityRequest struct {
	UserId      int    `form:"user_id"`     // 用户ID
	Mode        int    `form:"mode"`        // 查询模式，0: 按时间倒叙；1: 按点赞量倒叙
	Keyword     string `form:"keyword"`     // 搜索关键字
	CurrentPage int    `form:"currentPage"` // 当前页码
	PageSize    int    `form:"pageSize"`    // 分页数量
	IsOnline    int    `form:"isOnline"`    // 一级分类
	IsSchool    int    `form:"isSchool"`    // 二级分类
	IsBegin     int    `form:"isBegin"`
}

type UserFollowRequest struct {
	UserId     int64 `json:"user_id,omitempty"`
	OtherId    int64 `json:"other_id,omitempty"`
	ActionType int   `json:"action_type,omitempty"`
}
