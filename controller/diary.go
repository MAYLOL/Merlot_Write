package article

import (
	"github.com/kataras/iris"
	"merlot_project/merlot_write/model"
	"fmt"
	"merlot_project/merlot_write/diaryworkerpool"
	"time"
	"github.com/satori/go.uuid"
	"log"
	pb "merlot_project/merlot_write/service"
	"merlot_project/merlot_write/rpc"
	response2 "merlot_project/merlot_write/response"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)


type DiaryServiceInterface interface {
	//1.发送日记
	Diary(ctx iris.Context) (interface{},error)
	//2.更新日记
	UpdateDiary(ctx iris.Context)(interface{},error)
	//3.评论日记
	CommentDiary(ctx iris.Context)(interface{},error)
	//4.回复评论
	ReplyComment(ctx iris.Context)(interface{},error)
	//5.点赞日记
	LikeDiary(ctx iris.Context)(interface{},error)
	//6.点赞评论
	LikeComment(ctx iris.Context)(interface{},error)
	//7.点赞回复
	LikeReply(ctx iris.Context)(interface{},error)
	//8.取消评论点赞
	DislikeComment(ctx iris.Context)(interface{},error)
	//9.取消日记点赞
	DislikeDiary(ctx iris.Context)(interface{},error)
	//10.取消回复点赞
	DislikeReply(ctx iris.Context)(interface{},error)
}

type DiariesServices struct {
	diary model.Diary
}
//1.发表日记（只有文字图片，没有视频）
func(d *DiariesServices)Diary(ctx iris.Context) (interface{},error){
	diary:=new(model.Diary)
	//1.发表日记的时间
	diary.CreatedAt = time.Now().UnixNano()
	////2.日记更新的时间
	//3.日记的ID
	diaryID,_ := uuid.NewV1()
	diary.DiaryID = diaryID.String()
	//4.发表日记者的ID
	diary.UserID = ctx.FormValue("uid")
	//7.日记的标题
	diary.Title = ctx.FormValue("title")
	//8.日记的主要内容(富文本)
	diary.Content = ctx.FormValue("content")
	if len(strings.TrimSpace(ctx.FormValue("uid")))==0||len(strings.TrimSpace(ctx.FormValue("content"))) ==0||len(strings.TrimSpace(ctx.FormValue("secondary_category"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="参数错误"
		response.Data = response2.Data{}
		return response,nil
	}
	if len(strings.TrimSpace(ctx.FormValue("content")))!=0 {
		diary.DiaryText = ctx.FormValue("content")
		image:= pb.DiaryImage{}
        //暂时不支持视频
		//video:=pb.DiaryVideo{}
		p,err:=goquery.NewDocumentFromReader(bytes.NewBufferString(ctx.FormValue("content")))
		//p.Find("video").Each(func(i int, s *goquery.Selection){
		//	v,t:=s.Attr("src")
		//	println(v,t)
		//})

		p.Find("img").Each(func(i int, s *goquery.Selection){
			v,t:=s.Attr("src")
			if t==true {
				image.ImageBase64 =v
				diary.Images = append(diary.Images,&image)
				println(v,t)
			}

		})

		println("1111",err)
	}

	//10.日记的二级分类
	diary.DiaryType.DiaryTypeID = ctx.FormValue("DiaryTypeID")
	fmt.Println("调用了Diary方法")
	//调用缓冲层的RPC，将内容传递下去
	error:=DiarySaveRPC(diary)
	if error!=nil {
		response:=response2.Response{}
		response.Code = "301"
		return response,nil
	}
	producer:=diaryWorkerPool.NewProducer(*diary)
	producer.Output(*diary)
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}
//2.更新日记
func(d *DiariesServices)UpdateDiary(ctx iris.Context)(interface{},error){

	diary:=new(model.Diary)
	////1.发表日记的时间
	//diary.CreatedAt = time.Now().UnixNano()
	//2.日记更新的时间
	diary.UpdatedAt = time.Now().UnixNano()
	//3.日记的ID
	diaryID,_ := uuid.NewV1()
	diary.DiaryID = diaryID.String()
	//4.发表日记者的ID
	diary.UserID = ctx.FormValue("uid")

	//7.日记的标题
	diary.Title = ctx.FormValue("title")
	//8.日记的主要内容(富文本)
	diary.Content = ctx.FormValue("content")
	if len(strings.TrimSpace(ctx.FormValue("uid")))==0||len(strings.TrimSpace(ctx.FormValue("content"))) ==0||len(strings.TrimSpace(ctx.FormValue("did"))) ==0||len(strings.TrimSpace(ctx.FormValue("secondary_category"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="参数错误"
		response.Data = response2.Data{}
		return response,nil
	}
	if len(strings.TrimSpace(ctx.FormValue("content")))!=0 {
		diary.DiaryText = ctx.FormValue("content")
		image:= pb.DiaryImage{}

		p,err:=goquery.NewDocumentFromReader(bytes.NewBufferString(ctx.FormValue("content")))
		//p.Find("video").Each(func(i int, s *goquery.Selection){
		//	v,t:=s.Attr("src")
		//	println(v,t)
		//})

		p.Find("img").Each(func(i int, s *goquery.Selection){
			v,t:=s.Attr("src")
			if t==true {
				image.ImageBase64 =v
				diary.Images = append(diary.Images,&image)
				println(v,t)
			}
		})
		println("1111",err)

	}

	fmt.Println("调用了Diary方法")
    //10.日记的二级分类
    diary.DiaryType.DiaryTypeID =ctx.FormValue("DiaryTypeID")
	//调用缓冲层的RPC，将内容传递下去
	error:= DiaryUpdateRPC(diary)
	if error!=nil {
		response:=response2.Response{}
		response.Code = "301"
		return response,nil
	}
	producer:=diaryWorkerPool.NewProducer(*diary)
	producer.Output(*diary)
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}

//3.评论日记
func(d *DiariesServices)CommentDiary(ctx iris.Context)(interface{},error){
	comment:=new(model.Comment)

	//1.发布评论的时间
	comment.CreatedAt = time.Now().UnixNano()
	//2.评论的日记的ID
	comment.DiaryID = ctx.FormValue("did")
	//3.评论的ID
	CID,_:=uuid.NewV1()
	comment.DiaryID = CID.String()
	//4.评论者的UID
	comment.CommenterID =ctx.FormValue("uid")

    //5 评论的内容
    comment.Content = ctx.FormValue("content")
	if len(strings.TrimSpace(ctx.FormValue("uid")))==0||len(strings.TrimSpace(ctx.FormValue("content"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="参数错误"
		response.Data = response2.Data{}
		return response,nil
	}
	//RPC
	error:= DiaryCommentRPC(comment)
	if error!=nil {
		response:=response2.Response{}
		response.Code = "301"
		return response,nil
	}
	fmt.Println("调用了CommentDiary方法")
	//producer:=diaryWorkerPool.NewProducer(*diary)
	//producer.Output(*diary)
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}

//4.回复日记
func(d *DiariesServices)ReplyComment(ctx iris.Context)(interface{},error){
	reply:= new(model.Reply)
    //1.回复的时间
    reply.CreatedAt = time.Now().UnixNano()
    //2.回复的内容
    reply.Content = ctx.FormValue("content")

    //3.回复的ID
    RID,_ := uuid.NewV1()
    reply.ReplyID = RID.String()

    //4.回复的那个评论的ID
    reply.CommentID = ctx.FormValue("cid")
    //5.回复者的UID
    reply.ReplierID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid")))==0||len(strings.TrimSpace(ctx.FormValue("content"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="参数错误"
		response.Data = response2.Data{}
		return response,nil
	}
    //6日记的id
    reply.DiaryID = ctx.FormValue("did")
    //
  //回复日记RPC
    error:= DiaryReplyRPC(reply)
	if error!=nil {
		response:=response2.Response{}
		response.Code = "301"
		return response,nil
	}
	fmt.Println("调用了ReplyComment方法")
	//producer:=diaryWorkerPool.NewProducer(*diary)
	//producer.Output(*diary)
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}

//5.点赞日记
func(d *DiariesServices)LikeDiary(ctx iris.Context)(interface{},error){

	like :=new(model.Like)

	//1.点赞的时间
	like.CreatedAt = time.Now().UnixNano()
	//2.点赞的日记
	like.DiaryID = ctx.FormValue("did")
	//3.点赞用户的UID
	like.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
	//点赞日记RPC
	error:= DiaryLikeRPC(like)
	if error!=nil {
		response:=response2.Response{}
		response.Code = "301"
		return response,nil
	}
	fmt.Println("调用了LikeDiary方法")
	//producer:=diaryWorkerPool.NewProducer(*diary)
	//producer.Output(*diary)
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}

//6.点赞评论
func(d *DiariesServices)LikeComment(ctx iris.Context)(interface{},error){
	like:=new(model.Like)
	//1.点赞评论创建的时间
	like.CreatedAt = time.Now().UnixNano()
	//2.点赞的评论的ID
	like.CommentID = ctx.FormValue("cid")
	//3.点赞这的ID
	like.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
	//4.日记的ID
	like.DiaryID = ctx.FormValue("did")

	fmt.Println("调用了LikeComment方法")
	//producer:=diaryWorkerPool.NewProducer(*diary)
	//producer.Output(*diary)

	//点赞评论RPC

	 error:= DiaryCommentLikeRPC(like)
	if error!=nil {
		response:=response2.Response{}
		response.Code = "301"
		return response,nil
	}

	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}
//7.点赞回复
func(d *DiariesServices)LikeReply(ctx iris.Context)(interface{},error){
	like:=new(model.Like)
	//1.点赞创建的时间
	like.CreatedAt = time.Now().UnixNano()
	//2.点赞的回复的ID
	like.ReplyID = ctx.FormValue("rid")
	//3.点赞者的ID
	like.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
	//4.回复属于哪个评论
	like.CommentID = ctx.FormValue("cid")
	//5.属于哪个日记
	like.DiaryID  = ctx.FormValue("did")
	fmt.Println("调用了LikeReply方法")
	//producer:=diaryWorkerPool.NewProducer(*diary)
	//producer.Output(*diary)
	//点赞回复RPC
	error:=  DiaryReplyLikeRPC(like)
	if error!=nil {
		response:=response2.Response{}
		response.Code = "301"
		return response,nil
	}
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}
//8.取消评论点赞
func (d *DiariesServices)DislikeComment(ctx iris.Context)(interface{},error){
	dislike:=new(model.Like)
	//1.取消点赞的时间
	dislike.CreatedAt = time.Now().UnixNano()
	//2.取消点赞的评论的ID
	dislike.CommentID = ctx.FormValue("cid")
	//3.取消点赞的人的ID
	dislike.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
	//4.属于哪个日记
	dislike.DiaryID = ctx.FormValue("did")

	fmt.Println("调用了DislikeComment方法")
	//producer:=diaryWorkerPool.NewProducer(*diary)
	//producer.Output(*diary)
	//取消评论点赞RPC
	error:= DiaryCommentDislikeRPC(dislike)
	if error!=nil {
		response:=response2.Response{}
		response.Code = "301"
		return response,nil
	}
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}
//9.取消日记点赞
func (d *DiariesServices) DislikeDiary(ctx iris.Context)(interface{},error) {
	dislike:=new(model.Like)
	//1.取消点赞的时间
	dislike.CreatedAt = time.Now().Unix()
	//2.取消点赞的日记
	dislike.DiaryID = ctx.FormValue("did")
	//3.取消点赞者的UID
	dislike.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
	fmt.Println("调用了DislikeDiary方法")
	//producer:=diaryWorkerPool.NewProducer(*diary)
	//producer.Output(*diary)
	//取消日记点赞RPC
	error:=  DiaryDislikeRPC(dislike)
	if error!=nil {
		response:=response2.Response{}
		response.Code = "301"
		return response,nil
	}
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}
//10.取消回复点赞
func (d *DiariesServices)DislikeReply(ctx iris.Context)(interface{},error){
	dislike:=new(model.Like)
	fmt.Println("调用了DislikeReply方法")
	//1.取消点赞的时间
	dislike.CreatedAt = time.Now().UnixNano()
	//2.取消点赞的回复的ID
	dislike.ReplyID = ctx.FormValue("rid")
	//3.取消回复者的ID
	dislike.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		response.Data=response2.Data{}
		return response,nil
	}
   //4回复属于哪个评论
   dislike.CommentID = ctx.FormValue("cid")
   //5.回复属于哪个日记
   dislike.DiaryID = ctx.FormValue("did")
	//producer:=diaryWorkerPool.NewProducer(*diary)
	//producer.Output(*diary)
	//取消回复点赞RPC
	error:= DiaryReplyDislikeRPC(dislike)
	if error!=nil {
		response:=response2.Response{}
		response.Code = "301"
		return response,nil
	}
	//
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	return response,nil
}

//1.发表日记RPC
func DiarySaveRPC(diary *model.Diary)(error){
    //动作是发表日记
	action :=&pb.DiaryAction{ActionValue:pb.ActionDiaryType_DIARY_SAVE}
	//创建文章对象
	theDiary:=&pb.Diary{}
	//文章的富文本内容
	theDiary.Content = diary.Content
	diaryType := &pb.DiaryType{}
	diaryType.DiaryTypeID ="123456"
	diaryType.DiaryTypeName="眼睛"
	diaryType.DiaryTypeOnlineStatus =1
	diaryType.ParentID= "123456"
	diaryType.CreateTime = diary.CreatedAt

	diaryTypes:=[]*pb.DiaryType{}
	diaryTypes = append(diaryTypes,diaryType)

	theDiary.DiaryTypes = diaryTypes

	theDiaries:=[]*pb.Diary{}
	theDiaries = append(theDiaries,theDiary)

	in := pb.DMLDiaryRequest{Action:action,Diarys:theDiaries}
	r, err := rpc.ClientDiary.DMLDiary(rpc.ClientDiaryCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//2.更新日记RPC
func DiaryUpdateRPC(diary *model.Diary)(error){
	//动作是发表日记
	action :=&pb.DiaryAction{ActionValue:pb.ActionDiaryType_DIARY_MODIFY}

	//创建文章对象
	theDiary:=&pb.Diary{}
	//文章的富文本内容
	theDiary.Content = diary.Content

	theDiary.DiaryTitle = diary.Title



	diaryType := &pb.DiaryType{}
	diaryType.DiaryTypeID ="123456"
	diaryType.DiaryTypeName="眼睛"
	diaryType.DiaryTypeOnlineStatus =1
	diaryType.ParentID= "123456"
	diaryType.CreateTime = diary.CreatedAt



	diaryTypes:=[]*pb.DiaryType{}
	diaryTypes = append(diaryTypes,diaryType)

	theDiary.DiaryTypes = diaryTypes

	theDiaries:=[]*pb.Diary{}
	theDiaries = append(theDiaries,theDiary)

	in := pb.DMLDiaryRequest{Action:action,Diarys:theDiaries}
	r, err := rpc.ClientDiary.DMLDiary(rpc.ClientDiaryCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//3.发表评论RPC
func DiaryCommentRPC(comment *model.Comment)(error){
	//动作是为日记发表评论
	action :=&pb.DiaryAction{ActionValue:pb.ActionDiaryType_DIARY_COMMENT}

	//1.创建评论对象
	commentType:=&pb.DiaryComment{}
	//2.评论的内容
	commentType.Content = comment.Content
	//3.评论日记的时间
	commentType.CreateTime = time.Now().UnixNano()
	//4.评论的日记的ID
	commentType.CommentID = comment.DiaryID
	//5.评论者的UID
	commentType.UserID = comment.CommenterID
	commentTypes:=[]*pb.DiaryComment{}
	commentTypes = append(commentTypes,commentType)


	in := pb.DMLDiaryCommentRequest{Action:action,Comments:commentTypes}
	r, err := rpc.ClientDiary.DMLComment(rpc.ClientDiaryCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//4.发表回复RPC
func DiaryReplyRPC(reply *model.Reply)(error){
	//动作是为日记的评论发表回复
	action :=&pb.DiaryAction{ActionValue:pb.ActionDiaryType_DIARY_REPLY}

	//1.创建回复对象
	replyType:=&pb.DiaryComment{}
	//2.回复的内容
	replyType.Content = reply.Content
	//3.回复日记的时间
	replyType.CreateTime = time.Now().UnixNano()
	//4.回复的日记的ID  记住替换为DiaryID
	replyType.DiaryID = reply.CommentID
	//5.回复的ID
	replyType.CommentID = reply.ReplyID
	//6.回复者的name

	//7.回复者的ID
	replyType.UserID = reply.ReplierID


	replyTypes:=[]*pb.DiaryComment{}
	replyTypes = append(replyTypes,replyType)


	in := pb.DMLDiaryCommentRequest{Action:action,Comments:replyTypes}
	r, err := rpc.ClientDiary.DMLComment(rpc.ClientDiaryCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil

}
//5.点赞日记RPC
func DiaryLikeRPC(like * model.Like)(error){
	//动作是点赞日记
	action :=&pb.DiaryAction{ActionValue:pb.ActionDiaryType_DIARY_LIKE}

	//创建日记点赞对象
	 likeType:=&pb.DiaryLike{}
	 //1.点赞的日记的ID
     likeType.DiaryId = like.DiaryID
     //2.点赞日记的人的UID
     likeType.UserId = like.UserID
	 likeTypes:=[]*pb.DiaryLike{}
	 likeTypes = append(likeTypes,likeType)

	in := pb.DMLDiaryLikeRequest{Action:action,Likes:likeTypes}
	r, err := rpc.ClientDiary.DMLLike(rpc.ClientDiaryCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//6.取消点赞日记RPC
func DiaryDislikeRPC(dislike *model.Like)(error){
	//动作是取消点赞日记
	action :=&pb.DiaryAction{ActionValue:pb.ActionDiaryType_DIARY_DISLIKE}
	//创建日记点赞对象
	dislikeType:=&pb.DiaryLike{}
	//1.点赞的日记的ID
	dislikeType.DiaryId = dislike.DiaryID
	//2.点赞日记的人的UID
	dislikeType.UserId = dislike.UserID

	dislikeTypes:=[]*pb.DiaryLike{}
	dislikeTypes = append(dislikeTypes,dislikeType)

	in := pb.DMLDiaryLikeRequest{Action:action,Likes:dislikeTypes}
	r, err := rpc.ClientDiary.DMLLike(rpc.ClientDiaryCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//7.点赞评论RPC
func DiaryCommentLikeRPC(like *model.Like)(error){
	//动作是取消点赞日记
	action :=&pb.DiaryAction{ActionValue:pb.ActionDiaryType_DIARY_COMMENT_LIKE}

	//创建日记点赞对象
	likeType:=&pb.DiaryLike{}
	//1.点赞的评论的
	likeType.CommentId = like.CommentID
	//2.点赞评论的人的UID
	likeType.UserId = like.UserID

    likeTypes:=[]*pb.DiaryLike{}
	likeTypes = append(likeTypes,likeType)

	in := pb.DMLDiaryLikeRequest{Action:action,Likes:likeTypes}
	r, err := rpc.ClientDiary.DMLLike(rpc.ClientDiaryCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//8.取消点赞评论RPC
func DiaryCommentDislikeRPC(dislike *model.Like)(error){

	//动作是取消点赞评论
	action :=&pb.DiaryAction{ActionValue:pb.ActionDiaryType_DIARY_COMMENT_DISLIKE}

	//创建日记点赞对象
	dislikeType:=&pb.DiaryLike{}
	//1.取消点赞的评论的ID
	dislikeType.CommentId = dislike.CommentID
	//2.取消点赞评论的人的UID
	dislikeType.UserId = dislike.UserID
	//3.取消点赞日记的人的昵称
	dislikeType.UserName = dislike.UserName
	//4.取消点赞者的头像
	//
	dislikeTypes:=[]*pb.DiaryLike{}
	dislikeTypes = append(dislikeTypes,dislikeType)
	in := pb.DMLDiaryLikeRequest{Action:action,Likes:dislikeTypes}
	r, err := rpc.ClientDiary.DMLLike(rpc.ClientDiaryCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//9.点赞回复RPC
func DiaryReplyLikeRPC(like *model.Like)(error){
	//动作是取消点赞评论
	action :=&pb.DiaryAction{ActionValue:pb.ActionDiaryType_DIARY_REPLY_LIKE}

	//创建日记点赞对象
	likeType:=&pb.DiaryLike{}
	//1.取消点赞的回复的ID
	likeType.CommentId = like.CommentID
	//2.取消点赞回复的人的UID
	likeType.UserId = like.UserID
	likeTypes:=[]*pb.DiaryLike{}
	likeTypes = append(likeTypes,likeType)

	in := pb.DMLDiaryLikeRequest{Action:action,Likes:likeTypes}
	r, err := rpc.ClientDiary.DMLLike(rpc.ClientDiaryCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}
//10.取消点赞回复RPC
func DiaryReplyDislikeRPC(dislike *model.Like)(error){

	//动作是取消点赞评论
	action :=&pb.DiaryAction{ActionValue:pb.ActionDiaryType_DIARY_REPLY_DISLIKE}
	//创建日记点赞对象
	dislikeType:=&pb.DiaryLike{}
	//1.取消点赞的回复的ID
	dislikeType.CommentId = dislike.CommentID
	//2.取消点赞回复的人的UID
	dislikeType.UserId = dislike.UserID
	dislikeTypes:=[]*pb.DiaryLike{}
	dislikeTypes = append(dislikeTypes,dislikeType)

	in := pb.DMLDiaryLikeRequest{Action:action,Likes:dislikeTypes}
	r, err := rpc.ClientDiary.DMLLike(rpc.ClientDiaryCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}