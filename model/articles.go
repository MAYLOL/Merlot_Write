package model

import 	pb "merlot_project/merlot_write/service"
type Article struct {
	//文章的ID（服务端）
	ArticleID string `json:"article_id"`
	//文章创建的时间（服务端）
	CreatedAt int64 `json:"created_at"`
	////文章更新的时间
	UpdatedAt int64 `json:"updated_at"`
	//文章的标题（客户端）
	Title string `json:"title"`
	////文章处于审核的哪个状态
	Status int64 `json:"status"`
	//文章的(富文本内容）
	Content string `json:"content"`
	//发表文章的用户
	UserID string `json:"user_id"`
	//发送的内容的文字部分
	ArticleText  string `json:"article_text"`
	//发送的多个图片
	Images []*pb.ArticleImage `json:"images"`
	//发送的视频（待定）
	Videos []*pb.ArticleVideo `json:"videos"`
	//封面图片
	Cover string `json:"cover"`
    //文章标签
	ArticleType ArticleType `json:"article_type"`
    //是否上链
    HasUpChain bool `json:"has_up_chain"`
    //赏金
    MoneyReward float64 `json:"money_reward"`
    //赏金状态
    MoneyRewardStatus uint `json:"money_reward_status"`

    //文章分类集合:目前标签：医美、大健康、护肤、时尚、彩妆、随记
    //文字、图片、视频进行三方接口顾虑检测
    ArticleTypes []ArticleType `json:"article_types"`
    //分页
    ArticlePage ArticlePage `json:"article_page"`
	Labels []*pb.ArticleLabel `json:"labels"`


//备注：文章需要根据（评论数、点赞数、分享数、创建时间）等维度进行计算得出文章进行热门类别集合当中----热度排行
//推荐文章
//1、用户登陆
//一、根据自己内推的文章30%
//二、根据用户所关注的人-从热点文章中过滤出的30%
//三、用户自己大的文章分类标签-过滤除的文章40%

//2、用户没有登陆
//一、我们自己内推的文章30%
//二、热度高的文章30%
//三、客户端本地记录浏览行为40%


}
const (
	//ArticleVerifying 审核中
	ArticleVerifying = 1
	//ArticleVerifySuccess 审核通过
	ArticleVerifySuccess =2
	//ArticleVerifyFail 审核未通过
	ArticleVerifyFail = 3
)

