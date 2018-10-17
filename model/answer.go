package model
import 	pb "merlot_project/merlot_write/service"
type Answer struct{
	//1.回答的ID
	AnswerID           string         `json:"answer_id"`
	//2.创建的时间
	CreateTime         int64          `json:"create_time"`
	//3.更新的时间
	UpdateTime		   int64 		  `json:"update_time"`
	//4.回答的标题
	AnswerTitle 	   string 		  `json:"answer_title"`
	//5.审核状态
	Status			    int64         `json:"status"`
	//6.回答的内容
	Content 			string 		  `json:"content"`
	//12.是否已经上链
	HasUpChain          bool          `json:"has_up_chain"`
	//13.奖励金
	MoneyReward         float64       `json:"money_reward"`
	//14.奖励状态
	MoneyRewardStatus    int64         `json:"money_reward_status"`
	//15.用户的ID
	UserID               string        `json:"user_id"`



	CoverMap             *Image   	   `json:"cover_map"`
	//19.主题的ID
	TopicID              string        `json:"topic_id"`
	//20.主题的内容
	TopicText            string        `json:"topic_text"`
	//21.回答的图片
	Images               []*pb.TopicImage 		`json:"images"`
	//22.回答里的视频
	Videos               []*pb.TopicVideo 		`json:"videos"`
	//Labels               []*la `json
	//XXX_NoUnkeyedLiteral struct{}     `json:"xxx_no_unkeyed_literal"`
	XXX_unrecognized     []byte
	XXX_sizecache        int32
   //标签
   Labels []pb.TopicLabel `json:"labels"`
	AnswerText         string `json:"answer_text"`
}
