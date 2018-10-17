package model

type Like struct {
	//点赞创建的时间
	CreatedAt int64 `json:"created_at"`
	//点赞更新的时间
	UpdatedAt int64 `json:"updated_at"`
	//取消点赞的时间
	DeletedAt int64 `json:"deleted_at"`
	//点赞的人
	UserID    string `json:"user_id"`
	//点赞的文章
	ArticleID string  `json:"article_id"`
	//点赞的日记
	DiaryID string `json:"diary_id"`
	//点赞的主题
	TopicID string `json:"topic_id"`
	//是否点赞
	HasLiked bool  `json:"has_liked"`
	//点赞用户的昵称
	UserName string `json:"user_name"`
	//点赞用户的照片
	Portrait string `json:"portrait"`
	//点赞的评论
	CommentID string `json:"comment_id"`
	//点赞的回复
	ReplyID string `json:"reply_id"`
	//点赞的回答的ID
	AnswerID string `json:"answer_id"`
}


