package model

type User struct {
	//用户的ID
	UserID  string    `json:"user_id"`
	//用户的昵称
	NickName   string `json:"nick_name"`
	//用户的头像
	Portrait   []byte  `json:"portrait"`
	//用户的等级
	Level    int64   `json:"level"`
	//用户的权重
	Power    int64   `json:"power"`
    //用户的性别
    Gen     int64  `json:"gen"`
}
