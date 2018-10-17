package model

type Comment struct{
    //评论的ID（服务器）
    CommentID    string `json:"comment_id"`
    //评论创建的时间（服务器））
	CreatedAt  int64 `json:"created_at"`
	//评论更新的时间
	UpdatedAt  int64 `json:"updated_at"`
	//评论者ID
	CommenterID  string `json:"commenter_id"`
	//评论的文章
	ArticleID    string `json:"article_id"`
	//评论的日记ID
	DiaryID string `json:"diary_id"`
	//评论的主题ID
	TopicID string `json:"topic_id"`
	//回复数
	ReplyNum int64 `json:"reply_num"`
    //可见行
    Visible   bool   `json:"visible"`
    //portrait
    Portrait  string   `json:"portrait"`
    //点赞数
    LikeNum int64 `json:"like_num"`
    //评论者头像
    CommenterPortrait string `json:"commenter_portrait"`
    //Commenter评论者昵称
    CommenterNickName string  `json:"commenter_nick_name"`
    //评论的内容
    Content string `json:"content"`

}

const(
	//CommentVerifying 审核中
	CommentVerifying = 1
	//CommentVerifySuccess 审核通过
	CommentVerifySuccess =2
	//CommentVerifyFail  审核失败
	CommentVerifyFail = 3
)

