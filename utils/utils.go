package utils

// Page 分页
func Page(Limit, Page int64) (limit, offset int64) {
	if Limit > 0 {
		limit = Limit
	} else {
		limit = 10
	}
	if Page > 1 {
		offset = (Page - 1) * limit
	} else {
		offset = -1
	}
	return limit, offset
}

// Sort 排序
// 默认 created_at desc
func Sort(Sort string) (sort string) {
	if Sort != "" {
		sort = Sort
	} else {
		sort = "created_at desc"
	}
	return sort
}
