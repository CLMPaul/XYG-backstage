package proto

type WorkTypeRequest struct {
	SearchText  string `form:"searchText"`
	CurrentPage int    `form:"currentPage"`
	PageSize    int    `form:"pageSize"`
}
