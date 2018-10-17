package model
type ArticleType struct {
	//文章分类id 
	 ArticleTypeID string `json:"article_type_id"`
	 //文章分类父id  0:根节点 
	 ParentID string `json:"parent_id"`
	//文章分类状态 1:非叶子节点、0:叶子节点 
	 ArticleTypeStatus uint `json:"article_type_status"`
	//文章分类是否上线 1:上线、0不上线（即用户看不到） 
	 ArticleTypeOnlineStatus uint `json:"article_type_online_status"`
	//文章分类名称 
	 ArticleTypeName string `json:"article_type_name"`
	//创建时间 
	 CreateTime int64   `json:"create_time"`
	//修改时间 
	 UpdateTime int64   `json:"update_time"`
}