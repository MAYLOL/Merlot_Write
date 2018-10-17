package article

import (
	"github.com/kataras/iris"
	"merlot_project/merlot_write/model"
	"fmt"
	"log"
	pb "merlot_project/merlot_write/service"
	"time"
	"github.com/satori/go.uuid"
	"merlot_project/merlot_write/rpc"
	response2 "merlot_project/merlot_write/response"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)
type TopicServiceInterface interface {
	//1.发送主题
	Topic(ctx iris.Context) (interface{},error)
	//2.更新主题
	UpdateTopic(ctx iris.Context)(interface{},error)
	//3.评论（参与）主题
	CommentTopic(ctx iris.Context)(interface{},error)
	//4.回复评论
	ReplyComment(ctx iris.Context)(interface{},error)
	//5.点赞主题
	LikeTopic(ctx iris.Context)(interface{},error)
	//6.点赞评论
	LikeComment(ctx iris.Context)(interface{},error)
	//7.点赞回复
	LikeReply(ctx iris.Context)(interface{},error)
	//8.取消评论点赞
	DislikeComment(ctx iris.Context)(interface{},error)
	//9.取消主题点赞
	DislikeTopic(ctx iris.Context)(interface{},error)
	//10.取消回复点赞
	DislikeReply(ctx iris.Context)(interface{},error)
}

type TopicsService struct {
	topic model.Topic
}

//1.发送主题
func(t *TopicsService)Topic(ctx iris.Context) (interface{},error){
	topic:=new(model.Topic)
	//1.发表主题的时间
	topic.CreatedAt = time.Now().UnixNano()
	//2.主题更新的时间
	topic.UpdatedAt = time.Now().UnixNano()
	//3.主题的ID
	TID,_ := uuid.NewV1()
	topic.TopicID = TID.String()
	//4.发表主题者的ID
	topic.UserID = ctx.FormValue("uid")
    print(ctx.FormValue("uid"))
	//7.主题的标题
	topic.Title = ctx.FormValue("title")
	//8.主题的主要内容(富文本)
	topic.Content = ctx.FormValue("content")
	println(ctx.FormValue("content"))
	if len(strings.TrimSpace(ctx.FormValue("uid")))==0||len(strings.TrimSpace(ctx.FormValue("content"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="参数错误"
		response.Data = response2.Data{}
		return response,nil
	}
	fmt.Println("调用了Diary方法")



	if len(strings.TrimSpace(ctx.FormValue("content")))!=0 {
		topic.TopicText = ctx.FormValue("content")
		image:=pb.TopicImage{}
		p,err:=goquery.NewDocumentFromReader(bytes.NewBufferString(ctx.FormValue("content")))
		//p.Find("video").Each(func(i int, s *goquery.Selection){
		//	v,t:=s.Attr("src")
		//	println(v,t)
		//})

		p.Find("img").Each(func(i int, s *goquery.Selection){
			v,t:=s.Attr("src")
			if t==true {
				image.ImageBase64 =v
				topic.Images = append(topic.Images,&image)
				println(v,t)
			}
		})
		println("1111",err)

	}

	//调用缓冲层的RPC，将内容传递下去
	err:= TopicSaveRPC(topic)
	if err !=nil{
		return nil,err
	}
	//producer:=topicWorkerPool.NewProducer(*topic)
	//producer.Output(*topic)
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}
//2.更新主题
func (t *TopicsService)UpdateTopic(ctx iris.Context)(interface{},error){
	topic:=new(model.Topic)
	//1.更新主题的时间
	topic.CreatedAt = time.Now().UnixNano()
	//3.主题的ID
	topic.TopicID = ctx.FormValue("tid")
	//4.发表主题者的ID
	topic.UserID = ctx.FormValue("uid")
	//7.日记的标题
	topic.Title = ctx.FormValue("title")
	//8.日记的主要内容(富文本)
	topic.Content = ctx.FormValue("content")
	fmt.Println("调用了Diary方法")
	if len(strings.TrimSpace(ctx.FormValue("uid")))==0||len(strings.TrimSpace(ctx.FormValue("content"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="参数错误"
		response.Data = response2.Data{}
		return response,nil
	}


	if len(strings.TrimSpace(ctx.FormValue("content")))!=0 {
		topic.TopicText = ctx.FormValue("content")
		image:= pb.TopicImage{}

		p,err:=goquery.NewDocumentFromReader(bytes.NewBufferString(ctx.FormValue("content")))
		//p.Find("video").Each(func(i int, s *goquery.Selection){
		//	v,t:=s.Attr("src")
		//	println(v,t)
		//})

		p.Find("img").Each(func(i int, s *goquery.Selection){
			v,t:=s.Attr("src")
			println(v,t)
			if t==true {
				image.ImageBase64 =v
				topic.Images = append(topic.Images,&image)
				println(v,t)
			}
		})
		println("1111",err)

	}

	//调用缓冲层的RPC，将内容传递下去
   err:=  TopicUpdateRPC(topic)
	if err!=nil {
		return nil,err
	}
	//producer:=topicWorkerPool.NewProducer(*topic)
	//producer.Output(*topic)
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}
//3.参与主题（或者叫回答主题）
func (t *TopicsService)CommentTopic(ctx iris.Context)(interface{},error){
	answer:=new(model.Answer)
	//1.参与的时间
	answer.CreateTime = time.Now().UnixNano()
	//2.参与的话题的ID
	answer.TopicID = ctx.FormValue("tid")
	//3.参与话题的人的UID
	answer.UserID = ctx.FormValue("uid")

	//4.回答的ID
	ansID,_ := uuid.NewV1()
	answer.AnswerID = ansID.String()
	//7.回答的内容富文本
	answer.Content = ctx.FormValue("content")
	if len(strings.TrimSpace(ctx.FormValue("uid")))==0||len(strings.TrimSpace(ctx.FormValue("content"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="参数错误"
		response.Data = response2.Data{}
		return response,nil
	}

	if len(strings.TrimSpace(ctx.FormValue("content")))!=0 {
		answer.AnswerText = ctx.FormValue("content")
		image:=pb.TopicImage{}
		video:=pb.TopicVideo{}
		p,err:=goquery.NewDocumentFromReader(bytes.NewBufferString(ctx.FormValue("content")))
		p.Find("video").Each(func(i int, s *goquery.Selection){
			v,t:=s.Attr("src")
			video.VideoBase64 = v
			answer.Videos = append(answer.Videos,&video)
			println(v,t)
		})

		p.Find("img").Each(func(i int, s *goquery.Selection){
			v,t:=s.Attr("src")
			println(v,t)
			if t==true {
				image.ImageBase64 =v
				answer.Videos = append(answer.Videos,&video)
			}
		})
		println("1111",err)

	}
	//8.回答的标题
	answer.AnswerTitle = ctx.FormValue("title")
	//回答主题RPC
	err:=TopicAnswerRPC(answer)
	if err!= nil {
		return nil,err
	}
	fmt.Println("调用了CommentTopic方法")
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}
//4.回复回答
func (t *TopicsService)ReplyComment(ctx iris.Context)(interface{},error){
	reply:=new(model.Reply)
	//1.创建回复的时间
	reply.CreatedAt = time.Now().UnixNano()
	//2.回复的ID
	RID,_ :=uuid.NewV1()
	reply.ReplyID = RID.String()
	//3.回复的内容
	reply.Content = ctx.FormValue("content")
	//4.回复的哪个answer
	reply.AnswerID = ctx.FormValue("answer_id")
	//5.回复者的UID
	reply.ReplierID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid")))==0||len(strings.TrimSpace(ctx.FormValue("content"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="参数错误"
		response.Data = response2.Data{}
		return response,nil
	}
	//6.回复的主题
	reply.TopicID = ctx.FormValue("tid")

	//回复回答的RPC
	err:= TopicReplyRPC(reply)
	if err != nil {
		return nil,err
	}
	fmt.Println("调用了ReplyComment方法")
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}
//5.点赞主题
func (t *TopicsService)LikeTopic(ctx iris.Context)(interface{},error){
	like:=new(model.Like)
	//1.点赞创建的时间
	like.CreatedAt = time.Now().UnixNano()
	//2.点赞的主题的ID
	like.TopicID = ctx.FormValue("tid")
	//3.点赞者的UID
	like.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
	err:=TopicLikeRPC(like)
	if err != nil {
		return nil,err
	}
	fmt.Println("调用了LikeTopic方法")
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}
//6.点赞（参与）评论
func (t *TopicsService)LikeComment(ctx iris.Context)(interface{},error){
	like:=new(model.Like)
	//1.点赞创建的时间
	like.CreatedAt = time.Now().UnixNano()
	//2.点赞的评论
	like.AnswerID = ctx.FormValue("answer_id")
	//3.点赞的人的UID
	like.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
	//4.属于哪个主题
	like.TopicID = ctx.FormValue("tid")
	  err:= TopicCommentLikeRPC(like)
	if err != nil {
		return nil,err
	}
	fmt.Println("调用了LikeComment方法")
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}
//7.点赞回复
func (t *TopicsService) LikeReply(ctx iris.Context)(interface{},error) {
	like:=new(model.Like)
	//1.点赞回复的时间
	like.CreatedAt = time.Now().UnixNano()
	//2.点赞的回复id
	like.ReplyID = ctx.FormValue("rid")
	//3.点赞者的uid
	like.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
	//4.属于哪个回答
	like.AnswerID = ctx.FormValue("answer_id")
	//5.属于哪个主题
	like.TopicID = ctx.FormValue("tid")
    err:= TopicReplyLikeRPC(like)
    return nil,err
	fmt.Println("调用了LikeReply方法")
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}
//8.取消回答点赞
func (t *TopicsService) DislikeComment(ctx iris.Context)(interface{},error){
	dislike:=new(model.Like)
   //1.取消点赞的时间
   dislike.CreatedAt = time.Now().UnixNano()
   //2.取消点赞的回答的ID
   dislike.AnswerID = ctx.FormValue("answer_id")
   //3.取消点赞的人
   dislike.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
   //回答属于哪个主题
   dislike.TopicID = ctx.FormValue("tid")

   err:= TopicCommentDislikeRPC(dislike)
	if err!=nil {
		return nil,err
	}
	fmt.Println("调用了DislikeComment方法")
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}

//9.取消主题点赞
func (t *TopicsService) DislikeTopic(ctx iris.Context)(interface{},error){
	dislike:=new(model.Like)
	//1.取消点赞的时间
	dislike.CreatedAt = time.Now().UnixNano()
	//2.取消点赞的主题的ID
	dislike.TopicID = ctx.FormValue("tid")
	//3.取消点赞的人
	dislike.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
	fmt.Println("调用了DislikeTopic方法")
	err:= TopicDislikeRPC(dislike)
	if err != nil {
		return nil,err
	}
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}
//10.取消回复点赞
func (t *TopicsService) DislikeReply(ctx iris.Context)(interface{},error){
	dislike:=new(model.Like)
	//1.取消点赞的时间
	dislike.CreatedAt = time.Now().UnixNano()
	//2.取消点赞的回复的ID
	dislike.ReplyID = ctx.FormValue("rid")
	//3.取消点赞的人
	dislike.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
	//4.属于哪个回答
	dislike.AnswerID = ctx.FormValue("answer_id")
	//5 属于哪个主题
	dislike.TopicID = ctx.FormValue("tid")
	fmt.Println("调用了DislikeReply方法")
	 err:=TopicReplyDislikeRPC(dislike)
	if err != nil {
		return nil,err
	}
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}


//1.发表主题RPC
func TopicSaveRPC(topic * model.Topic)(error){
	//动作是发表日记
	action :=&pb.TopicAction{ActionValue:pb.ActionTopicType_TOPIC_SAVE}
	//创建文章对象
	theTopic:=&pb.Topic{}
	//文章的富文本内容
	theTopic.Content = topic.Content
	////测试的图片
	//image:= &pb.TopicImage{}
	//image.ImageBase64 = topic.Images[0].ImageBase64
	//theTopic.Images = append(theTopic.Images,image)
	theTopics:=[]*pb.Topic{}
	theTopics = append(theTopics,theTopic)
	in := pb.DMLTopicRequest{Action:action,Topics:theTopics}
	r, err := rpc.ClientTopic.DMLTopic(rpc.ClientDiaryCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//2.更新主题RPC
func TopicUpdateRPC(topic *model.Topic)(error){
	//动作是发表日记
	action :=&pb.TopicAction{ActionValue:pb.ActionTopicType_TOPIC_MODIFY}
	//创建文章对象
	theTopic:=&pb.Topic{}
	//文章的富文本内容
	theTopic.Content = topic.Content
	//测试的图片
	//image:= &pb.TopicImage{}
	//image.ImageBase64 = topic.Images[0].ImageBase64
	//theTopic.Images = append(theTopic.Images,image)
	theTopics:=[]*pb.Topic{}
	theTopics = append(theTopics,theTopic)
	in := pb.DMLTopicRequest{Action:action,Topics:theTopics}
	r, err := rpc.ClientTopic.DMLTopic(rpc.ClientTopicCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//3.参与主题(回答主题 answer topic)
func TopicAnswerRPC(answer *model.Answer)(error){
	//动作是发表主题
	action :=&pb.TopicAction{ActionValue:pb.ActionTopicType_ANSWER_SAVE}

	//创建对象
	answerType:=&pb.Answer{}
    //1.回答创建的时间
	answerType.CreatedTime = time.Now().UnixNano()
    //2.回答ID
	answerType.AnswerID = answer.AnswerID
    //3.回答的主题的ID
	answerType.TopicID = answer.TopicID
	//4.回答的人的UID
	answerType.UserID =  answer.UserID

	//6.回答的标题
	answerType.AnswerTitle = answer.AnswerTitle
	//7.回答的内容
	answerType.Content = answer.Content
	//8.回答的图片
	//9.回答的视频
	//10.回答的文字

	answerTypes:=[]*pb.Answer{}
	answerTypes = append(answerTypes,answerType)

	in := pb.DMLTopicAnswerRequest{Action:action,Answers:answerTypes}
	r, err := rpc.ClientTopic.DMLAnswer(rpc.ClientTopicCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//4.回复评论（回复参与）
func TopicReplyRPC(reply *model.Reply)(error){
	//动作是回复参与的回答
	action :=&pb.TopicAction{ActionValue:pb.ActionTopicType_TOPIC_COMMENT_REPLY}
	//创建回复的对象
	replyType:=&pb.TopicComment{}
	//1.回复的时间
	replyType.CreateTime = reply.CreatedAt
	//2.回复的哪个问题
	replyType.AnswerID = reply.AnswerID
	//3.回复的ID
	replyType.CommentID = reply.ReplyID
	//4.回复的内容
	replyType.Content = reply.Content
	//5.回复者的ID
	replyType.UserID = reply.ReplierID
	replyTypes:=[]*pb.TopicComment{}
	replyTypes = append(replyTypes,replyType)

	in := pb.DMLTopicCommentRequest{Action:action,Comments:replyTypes}
	r, err := rpc.ClientTopic.DMLComment(rpc.ClientTopicCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//5.点赞主题
func TopicLikeRPC(like *model.Like)(error){
	//动作是点赞主题
	action :=&pb.TopicAction{ActionValue:pb.ActionTopicType_TOPIC_LIKE}
	//创建喜欢对象
	likeType:=&pb.TopicLike{}
	//1.喜欢的时间
	likeType.CreatedTime =  like.CreatedAt
	//2.喜欢的主题
	likeType.TopicId =like.TopicID
	//3.喜欢的人的UID
	likeType.UserId = like.UserID
	likeTypes:=[]*pb.TopicLike{}
	likeTypes = append(likeTypes,likeType)
	in := pb.DMLTopicLikeRequest{Action:action,Likes:likeTypes}
	r, err := rpc.ClientTopic.DMLLike(rpc.ClientTopicCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//6.取消点赞主题
func TopicDislikeRPC(dislike *model.Like)(error){
	//动作取消点赞主题
	action :=&pb.TopicAction{ActionValue:pb.ActionTopicType_TOPIC_DISLIKE}
	//创建喜欢对象
	dislikeType:=&pb.TopicLike{}
	//1.取消点赞创建的时间
	dislikeType.CreatedTime = dislike.CreatedAt
	//2.取消点赞的主题ID
	dislikeType.TopicId = dislike.TopicID
	//3.取消点赞者的UID
	dislikeType.UserId = dislike.UserID
	dislikeTypes:=[]*pb.TopicLike{}
	dislikeTypes = append(dislikeTypes,dislikeType)

	in := pb.DMLTopicLikeRequest{Action:action,Likes:dislikeTypes}
	r, err := rpc.ClientTopic.DMLLike(rpc.ClientTopicCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//7.点赞 answer（参与主题成为answer）
func TopicCommentLikeRPC(like *model.Like)(error){
	//动作是点赞comment
	action :=&pb.TopicAction{ActionValue:pb.ActionTopicType_TOPIC_COMMENT_LIKE}
	//创建文章对象
	likeType:=&pb.TopicLike{}
    //1.点赞的时间
    likeType.CreatedTime = like.CreatedAt
    //2.点赞的answerID
    likeType.AnswerId = like.AnswerID
    //3.点赞者的userID
    likeType.UserId = like.UserID

	likeTypes:=[]*pb.TopicLike{}
	likeTypes = append(likeTypes,likeType)
	in := pb.DMLTopicLikeRequest{Action:action,Likes:likeTypes}
	r, err := rpc.ClientTopic.DMLLike(rpc.ClientTopicCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//8.取消点赞评论
func TopicCommentDislikeRPC(dislike *model.Like)(error){
	//动作是取消回答点赞
	action :=&pb.TopicAction{ActionValue:pb.ActionTopicType_ANSWER_DISLIKE}
	//创建like对象
	dislikeType:=&pb.TopicLike{}
    //1.取消点赞的时间
	dislikeType.CreatedTime = dislike.CreatedAt
	//2.取消点赞的评论
	dislikeType.AnswerId = dislike.AnswerID
	//3.取消点赞的UID
	dislikeType.UserId = dislike.UserID
	dislikeTypes:=[]*pb.TopicLike{}
	dislikeTypes = append(dislikeTypes,dislikeType)
	in := pb.DMLTopicLikeRequest{Action:action,Likes:dislikeTypes}
	r, err := rpc.ClientTopic.DMLLike(rpc.ClientTopicCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//9.点赞回复
func TopicReplyLikeRPC(like *model.Like)(error){
	//动作是点赞回复
	action :=&pb.TopicAction{ActionValue:pb.ActionTopicType_TOPIC_REPLY_LIKE}
	//创建喜欢对象
	likeType:=&pb.TopicLike{}
	//1.喜欢创建的时间
	likeType.CreatedTime = like.CreatedAt
	//2.喜欢回复的ID
	likeType.CommentId = like.ReplyID
	//3.点赞者UID
	likeType.UserId = like.UserID
	likeTypes:=[]*pb.TopicLike{}
	likeTypes = append(likeTypes,likeType)
	in := pb.DMLTopicLikeRequest{Action:action,Likes:likeTypes}
	r, err := rpc.ClientTopic.DMLLike(rpc.ClientTopicCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//10.取消点赞回复
func TopicReplyDislikeRPC(like * model.Like)(error){
	//动作取消点赞回复
	action :=&pb.TopicAction{ActionValue:pb.ActionTopicType_TOPIC_REPLY_DISLIKE}
	//创建dislike对象
	dislikeType:=&pb.TopicLike{}
	//1.取消点赞的时间
	dislikeType.CreatedTime = like.CreatedAt
	//2.取消点赞的回复的ID
	dislikeType.CommentId = like.ReplyID
   //3.取消点赞者的UID
   dislikeType.UserId = like.UserID
	dislikeTypes:=[]*pb.TopicLike{}
	dislikeTypes = append(dislikeTypes,dislikeType)
	in := pb.DMLTopicLikeRequest{Action:action,Likes:dislikeTypes}
	r, err := rpc.ClientTopic.DMLLike(rpc.ClientTopicCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}