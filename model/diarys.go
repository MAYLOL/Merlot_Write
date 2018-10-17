package model
import 	pb "merlot_project/merlot_write/service"
//美丽日记
type Diary struct {
	//日记的ID
	DiaryID     string `json:"diary_id"`
	//主题创建的时间
	CreatedAt  int64 `json:"created_at"`
	//主题更新的时间
	UpdatedAt  int64 `json:"updated_at"`
	//主题的标题
	Title    string   `json:"title"`
	////主题处于审核的哪个状态
	Status   int64   `json:"status"`
	//主题的富文本的内容
	Content  string   `json:"content"`

	//发表主题的用户ID
	UserID string `json:"user_id"`
    //diaryType
    DiaryType DiaryType `json:"diary_type"`
	//发送的内容的文字部分
	DiaryText string `json:"diary_text"`
	//发送的多个图片
	Images []*pb.DiaryImage `json:"images"`
	//发送的视频（日记可能暂时不支持发视频）
	//Videos []*pb.DiaryVideo `json:"videos"`
	//是否上链
	HasUpChain bool `json:"has_up_chain"`
	//赏金
	MoneyReward float64 `json:"money_reward"`
	//赏金状态
	MoneyRewardStatus uint `json:"money_reward_status"`
   //分页
     DiaryPage DiaryPage `json:"diary_page"`
    //标签
    Labels []pb.DiaryLabel `json:"labels"`



}