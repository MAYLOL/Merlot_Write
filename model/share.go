package model

type Share struct {
	//分享的时间
	CreatedAt int64 `json:"created_at"`
	//分享文章的人
	UserID    string `json:"user_id"`
	//分享的文章
	Article  Article `json:"article"`
	//分享成功没成功
	IsSuccess   bool `json:"is_success"`
	//分享的状态
	Status      int64 `json:"status"`
}

const(
	//Sharing 分享中
	Sharing = 1
	//SharedSuccess 分享成功
	SharedSuccess = 2
	//SharedFail 分享失败
	SharedFail = 3


)
