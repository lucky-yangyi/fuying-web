// @Author yangyi 2023/7/26 20:27:00
package utils

func GetPageOffset(page, pageSize uint64) uint64 {
	//页码数  1
	//每页数量 10
	var result uint64
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
