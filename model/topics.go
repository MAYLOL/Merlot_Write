package model
//接收上传来的话题
import 	pb "merlot_project/merlot_write/service"

type Topic struct {
//主题的ID
TopicID     string `json:"topic_id"`
//主题创建的时间
CreatedAt  int64 `json:"created_at"`
//主题更新的时间
UpdatedAt  int64 `json:"updated_at"`
//主题的标题
Title    string   `json:"title"`
////主题处于审核的哪个状态
Status   int64     `json:"status"`
//主题的富文本的内容
Content  string   `json:"content"`

//发表主题的用户ID
UserID string  `json:"user_id"`

//标签
Labels []pb.TopicLabel `json:"labels"`
//发送的内容的文字部分
TopicText  string `json:"article_text"`
//发送的多个图片
Images []*pb.TopicImage `json:"images"`
//发送的视频（暂时可能不支持视频上传）
//Videos []*pb.TopicVideo `json:"videos"`
//评论数
CommentNum int64  `json:"comment_num"`

//是否上链
HasUpChain bool `json:"has_up_chain"`
//赏金
MoneyReward float64 `json:"money_reward"`
//赏金状态
MoneyRewardStatus uint `json:"money_reward_status"`
//topicType
TopicType TopicType `json:"topic_type"`
//分页
TopicPage TopicPage `json:"topic_page"`
}