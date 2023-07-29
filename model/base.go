// @Author yangyi 2023/7/26 20:37:00
package model

//type Paginations struct {
//	Current    int   `json:"current" form:"page"`      // 页码
//	PageSize   int   `json:"pageSize" form:"pageSize"` // 每页大小
//	TotalCount int64 `json:"totalCount"`               //总数
//}

type Pagination struct {
	Current    uint64 `json:"current" form:"page"`      // 页码
	PageSize   uint64 `json:"pageSize" form:"pageSize"` // 每页大小
	TotalCount uint64 `json:"totalCount"`               //总数
}
