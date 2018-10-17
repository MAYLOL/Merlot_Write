package model

type TopicType struct {
	//话题分类ID
	TopicTypeID int64 `json:"topic_type_id"`
	//话题分类父ID
	TopicParentID  int64 `json:"topic_parent_id"`
	//话题分类状态：1：非叶子节点 0:叶子节点
	TopicTypeStatus uint `json:"topic_type_status"`
	//话题分类是否上线：1上线 0不上线（用户看不到）
	TopicTypeOnlineStatus uint `json:"topic_type_online_status"`
	//话题分类名称
	TopicTypeName string `json:"topic_type_name"`
	//创建时间
	CreateTime  int64 `json:"create_time"`
	//修改时间
	UpdateTime int64 `json:"update_time"`
}