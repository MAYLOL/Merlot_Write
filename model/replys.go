package model

type Reply struct {
	//创建时间
	CreatedAt int64 `json:"created_at"`
	//回复更新的时间
	UpdatedAt  int64 `json:"updated_at"`
	//回复的ID
    ReplyID string `json:"reply_id"`
	//回复的评论ID
	CommentID string   `json:"comment_id"`
	//回复的回答的ID
	AnswerID string `json:"answer_id"`
	//回复的评论属于那个日记
	DiaryID string `json:"diary_id"`
	//回复者ID
	ReplierID string   `json:"replier_id"`
	//可见性
	Visible   bool   `json:"visible"`
	//portrait
	ReplierPortrait string `json:"replier_portrait"`
	//点赞数
	LikeNum int64 `json:"like_num"`
	//回复者昵称
	ReplierNickName string  `json:"replier_nick_name"`
	//回复的内容
	Content string `json:"content"`
	//回复的主题
	TopicID   string `json:"topic_id"`

}