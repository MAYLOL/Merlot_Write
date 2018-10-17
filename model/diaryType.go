package model

type DiaryType struct {
	//日记分类id
	DiaryTypeID  string `json:"diary_type_id"`
	//日记分类父ID
	ParentID string `json:"parent_id"`
	//日记分类状态：1：非叶子节点 0：叶子节点
	DiaryTypeStatus uint `json:"diary_type_status"`
	//日记分类是否上线：1：上线、0：不上线（即用户看不到）
	DiaryTypeOnlineStatus uint `json:"diary_type_online_status"`
	//日记分类名称
	DiaryTypeName string `json:"diary_type_name"`
	//创建时间
	CreateTime int64  `json:"create_time"`
	//修改时间
	UpdateTime  int64 `json:"update_time"`


}
