package paginationhelper

func Pagination(params map[string]interface{}) (page int64, pageSize int64) {
	pageParams := params["Page"]
	page = 1
	if pageParams == nil || pageParams.(int64) == 0 {
	} else {
		page = pageParams.(int64)
	}
	pageSizeParams := params["PageSize"]
	pageSize = 10
	if pageSizeParams == nil || pageSizeParams.(int64) == 0 {
	} else {
		pageSize = pageSizeParams.(int64)
	}
	return page, pageSize
}
