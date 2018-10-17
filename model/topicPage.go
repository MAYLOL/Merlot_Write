package model

type TopicPage struct {
	//每页条数
	PageSize int64 `json:"page_size"`
	//当前是第几页
	OffSet int64 `json:"off_set"`
	//总条数
	Total int64 `json:"total"`
	//总页数
	TotalPageSize int64 `json:"total_page_size"`
	//是否是首页
	IsTop bool `json:"is_top"`
	//是否是尾页
	IsBottom bool `json:"is_bottom"`
}