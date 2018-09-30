package models

//PagingData ..
func PagingData(pageIndex int, rowsIndex int, pageCount int, total int, rows interface{}) interface{} {
	return map[string]interface{}{
		"pageIndex": pageIndex, //第几页  页数从1开始
		"rowsIndex": rowsIndex, //每页行数
		"pageCount": pageCount, //总页数
		"total":     total,     //总记录条数
		"rows":      rows,      //当前页的数据集合
	}
}
