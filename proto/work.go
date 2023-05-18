package proto

type WorkRequest struct {
	UserId      int64  `form:"user_id"`     // 用户ID
	Mode        int    `form:"mode"`        // 查询模式，0: 按时间倒叙；1: 按点赞量倒叙
	Keywords    string `form:"keywords"`    // 搜索关键字
	State       *int   `form:"state"`       // 作品状态
	CurrentPage int    `form:"currentPage"` // 当前页码
	PageSize    int    `form:"pageSize"`    // 分页数量
	TypeID      *int   `form:"typeID"`      // 类别ID
}

type WorkForm struct {
	Id          int      `json:"id"`
	Introduce   string   `json:"introduce"`
	Title       string   `json:"title"`        // 作品名称
	WorkPicture []string `json:"work_picture"` // 照片集
	TypeID      int64    `json:"typeID"`       // 作品类别
}
