package article
import (
	"github.com/kataras/iris"
	uuid2"merlot_write/sdk/src/uuid"
	"green-go-sdk/src/greensdksample"
	"fmt"
	"merlot_project/merlot_write/articleworkerpool"
	 "merlot_project/merlot_write/model"
	"github.com/satori/go.uuid"
	"time"
	"log"
	pb "merlot_project/merlot_write/service"
	"merlot_project/merlot_write/rpc"
	response2 "merlot_project/merlot_write/response"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"bytes"
)

//定义接口服务
type ArticleServiceInterface interface{
	//1.发送文章
	Article(ctx iris.Context)(interface{},error)
	////2.更新文章
	UpdateArticle(ctx iris.Context)(interface{},error)
	////3.评论文章
	CommentArtcle(ctx iris.Context)(interface{},error)
	////4.回复评论
	ReplyComment(ctx iris.Context)(interface{},error)
	////5.点赞文章
	LikeArticle(ctx iris.Context)(interface{},error)
	////6.点赞评论
	LikeComment(ctx iris.Context)(interface{},error)
	////7.点赞回复
	LikeReply(ctx iris.Context)(interface{},error)
	////8.取消评论点赞
	DislikeComment(ctx iris.Context)(interface{},error)
	////9.取消文章点赞
	DislikeArticle(ctx iris.Context)(interface{},error)
	////10.取消回复点赞
	DislikeReply(ctx iris.Context)(interface{},error)
	////11.文章初始化表单格式化
	ArticleInitFromForm(ctx iris.Context)(model.Article,error)
}
//外部调用的结构
type ArticlesService struct {
	artcle model.Article
}
const accessKeyId  = "LTAIvHoOwjUZuzf9"
const accessKeySecret  = "4EQGd1iUevUW48P4FUo1QBEWrmuzmK"
//创建了切片，并将结构体插入到切片当中
var structSlice []model.Article
//这个方法应该满足发送成功返回成功发表，发送失败的时候返回失败错误：这一种是error，一种是内容审核有问题
//METHOD：
//1.发表文章
func (a *ArticlesService) Article(ctx iris.Context) (interface{},error){
      article:=new(model.Article)
      //1.生成文章ID
      aid,_ := uuid.NewV1()
      article.ArticleID = aid.String()
      //2.1记录发表文章的时间戳
	  article.CreatedAt = time.Now().UnixNano()
	  ////2.2记录更新文章的时间（）
      article.UserID =ctx.FormValue("uid")
      println("uid:",ctx.FormValue("uid"))
      //4.获取文章title   ok
      article.Title = ctx.FormValue("title")
      //5.文章的审核状态   初始值设置为0

      //6.content富文本内容
      article.Content =ctx.FormValue("content")
      println("content",ctx.FormValue("content"))
     //goquery.NewDocumentFromReader()
	//document,err := goquery.NewDocumentFromReader(bytes.NewBufferString(ctx.FormValue("content")))

	//fmt.Println("0000",(document.),err)
	if len(strings.TrimSpace(ctx.FormValue("uid")))==0||len(strings.TrimSpace(ctx.FormValue("content"))) ==0||len(strings.TrimSpace(ctx.FormValue("top_category"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="参数错误"
		response.Data = response2.Data{}
		return response,nil
	}

	if len(strings.TrimSpace(ctx.FormValue("content")))!=0 {
		article.ArticleText = ctx.FormValue("content")

		image:= pb.ArticleImage{}

		video:=pb.ArticleVideo{}

		p,err:=goquery.NewDocumentFromReader(bytes.NewBufferString(ctx.FormValue("content")))
        p.Find("video").Each(func(i int, s *goquery.Selection){
			v,t:=s.Attr("src")
			if t==true {
			video.VideoBase64 = v
			article.Videos = append(article.Videos,&video)
			}
			println(v,t)
		})

		p.Find("img").Each(func(i int, s *goquery.Selection){
			v,t:=s.Attr("src")
			if t==true {
				image.ImageBase64 =v
				article.Images = append(article.Images,&image)
				println(v,t)
			}

		})
		println("1111",err)




        }

      //7.文章的一级分类
      article.ArticleType.ParentID = ctx.FormValue("parentID")
      //8.文章的二级分类
      article.ArticleType.ArticleTypeID = ctx.FormValue("ArticleTypeID")
      //9.文章的点赞，评论分享等数
      //article.CommentNum =0
      //article.LikeNum = 0
      //article.ShareNum =0

      //10.提取出来的文字详情内容
      //article.ArticleText = ctx.FormValue("article_text")
      //11.提取出来的图片（）//暂时不从富文本解析先传

      //12.提取出来的视频
      //******处理富文本，提取文本，视频，图片
     error:=ArticleSaveRpc(article)
	if error!=nil {
		response:=response2.Response{}
		response.Code = "301"
		return response,nil
	}
	 producer:=articleWorkerPool.NewProducer(*article)
	 producer.Output(*article)

	 response:=response2.Response{}
	 response.Code ="0"
	 response.Message = "success"
	 response.Data = response2.Data{}
	defer fmt.Println("文章=====",article.Images[0].ImageBase64)
     return response,nil
}
//2.更新文章
func (a *ArticlesService)UpdateArticle(ctx iris.Context)(interface{},error){
	article:=new(model.Article)

	article.ArticleID = ctx.FormValue("aid")
	println("aid:",ctx.FormValue("aid"))

	//2.2记录更新文章的时间（）  ok
	article.UpdatedAt = time.Now().UnixNano()
	//3.获取发表文章用户ID     ok
	article.UserID =ctx.FormValue("uid")
	println("uid",ctx.FormValue("uid"))
	//4.获取文章title   ok
	article.Title = ctx.FormValue("title")
	//5.文章的审核状态   初始值设置为0

	//6.content富文本内容
	article.Content =ctx.FormValue("content")
	println("content:",ctx.FormValue("content"))
	if len(strings.TrimSpace(ctx.FormValue("uid")))==0||len(strings.TrimSpace(ctx.FormValue("content"))) ==0||len(strings.TrimSpace(ctx.FormValue("aid")))==0||len(strings.TrimSpace(ctx.FormValue("top_category"))) ==0{
		response :=response2.Response{}
		response.Code ="202"
		response.Message="参数错误"
		response.Data = response2.Data{}
		return response,nil
	}

	if len(strings.TrimSpace(ctx.FormValue("content")))!=0 {
		article.ArticleText = ctx.FormValue("content")
		image:= pb.ArticleImage{}
        video:=pb.ArticleVideo{}



		p,err:=goquery.NewDocumentFromReader(bytes.NewBufferString(ctx.FormValue("content")))
		p.Find("video").Each(func(i int, s *goquery.Selection){
			v,t:=s.Attr("src")
			if t == true {
				video.VideoBase64 = v
				article.Videos = append(article.Videos,&video)
			}
			println(v,t)
		})

		p.Find("img").Each(func(i int, s *goquery.Selection){
			v,t:=s.Attr("src")
			if t==true {
				image.ImageBase64 =v
				article.Images = append(article.Images,&image)
				println(v,t)
			}

		})

		println("1111",err)
	}
	//7.文章的一级分类
	article.ArticleType.ParentID = ctx.FormValue("ParentID")
	//8.文章的二级分类
	article.ArticleType.ArticleTypeID = ctx.FormValue("ArticleTypeID")
	//9.文章的点赞，评论分享等数
	//article.CommentNum =0
	//article.LikeNum = 0
	//article.ShareNum =0
	//10.提取出来的文字详情内容
	//article.ArticleText = ctx.FormValue("article_text")
	//11.提取出来的图片（）


	//12.提取出来的视频
	//******处理富文本，提取文本，视频，图片
	error:= ArticleUpdateRpc(article)
	if error!=nil {
		response:=response2.Response{}
		response.Code = "301"
		return response,nil
	}
	producer:=articleWorkerPool.NewProducer(*article)
	producer.Output(*article)
	response:=response2.Response{}
	response.Code ="0"
	response.Message = "success"
	response.Data = response2.Data{}


	return response,nil
}
//3.发表评论
func (a *ArticlesService)CommentArtcle(ctx iris.Context)(interface{},error){
	comment:=new(model.Comment)
	//1.生成评论的UUID 唯一标示
	cid,_ := uuid.NewV1()
	comment.CommentID =cid.String()
   //2.1发表评论的时间
    comment.CreatedAt = time.Now().UnixNano()
    ////2.2更新的时间
    comment.UpdatedAt =time.Now().UnixNano()
    //3.评论者的ID
    comment.CommenterID = ctx.FormValue("uid")
    //4.评论的内容
    comment.Content = ctx.FormValue("content")
	if len(strings.TrimSpace(ctx.FormValue("uid")))==0||len(strings.TrimSpace(ctx.FormValue("content"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="参数错误"
		response.Data = response2.Data{}
		return response,nil
	}
    //5.
	fmt.Println("调用了CommentArtcle方法")
	comment.ArticleID = ctx.FormValue("aid")
	error:=  ArticleCommentRPC(comment)
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
//4.回复评论
func (a *ArticlesService)ReplyComment(ctx iris.Context)(interface{},error){
	reply:=new(model.Reply)
	fmt.Println("调用了ReplyComment方法")
	//1.生成回复的唯一标识
	 replyID,_ := uuid.NewV1()
	reply.ReplyID = replyID.String()
	//2.回复的内容
	reply.Content =ctx.FormValue("content")
	//3.回复的那条评论的ID
	reply.CommentID =ctx.FormValue("cid")
	//6.回复的时间
	reply.CreatedAt= time.Now().UnixNano()
	//7.回复人的UID
	reply.ReplierID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid")))==0||len(strings.TrimSpace(ctx.FormValue("content"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="参数错误"
		response.Data = response2.Data{}
		return response,nil
	}
	error:=  ArticleReplyRPC(reply)
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
//5.点赞文章
func (a *ArticlesService)LikeArticle(ctx iris.Context)(interface{},error){
	like:=new(model.Like)
	fmt.Println("调用了LikeArticle方法")
	//1.点赞的文章的ID
	like.ArticleID=ctx.FormValue("aid")
	//2.点赞文章的时间
	like.CreatedAt = time.Now().UnixNano()
	//3.点赞者的UID
	like.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
	error:= ArticleLikeRPC(like)
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
//6.取消文章赞
func (a *ArticlesService)DislikeArticle(ctx iris.Context)(interface{},error){
	like:=new(model.Like)
	fmt.Println("调用了DislikeArticle方法")
	//1.取消点赞的文章的ID
	like.ArticleID=ctx.FormValue("aid")
	//2.取消点赞文章的时间
	like.CreatedAt = time.Now().UnixNano()
	//3.取消点赞者的UID
	like.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}

	error:=  ArticleDislikeRPC(like)
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
//7.点赞评论
func (a *ArticlesService)LikeComment(ctx iris.Context)(interface{},error){
	like:=new(model.Like)
	fmt.Println("调用了LikeComment方法")
	//1.点赞的评论ID
	like.CommentID = ctx.FormValue("cid")
	//评论属于那篇文章
	like.ArticleID = ctx.FormValue("aid")
	//2.点赞评论的时间
	like.CreatedAt = time.Now().UnixNano()
	//3.点赞者的UID
	like.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}

	error:=  CommentLikeRPC(like)
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
func (a *ArticlesService)DislikeComment(ctx iris.Context)(interface{},error){
	like:=new(model.Like)
	fmt.Println("调用了DislikeComment方法")
	//1.取消点赞的评论ID
	like.CommentID = ctx.FormValue("cid")
	//2.取消点赞评论的时间
	like.CreatedAt = time.Now().UnixNano()
	//3.取消点赞者的UID
	like.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
	//4.评论隶属的文章
	like.ArticleID = ctx.FormValue("aid")
	 error:= CommentDislikeRPC(like)
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
//9.点赞回复
func (a *ArticlesService)LikeReply(ctx iris.Context)(interface{},error){
	like:=new(model.Like)
	fmt.Println("调用了LikeReply方法")

	//1.点赞的回复的ID
	like.ReplyID = ctx.FormValue("rid")
	//2.点赞回复的时间
	like.CreatedAt = time.Now().UnixNano()
	//3.点赞者的ID
	like.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
    //4.回复的那条评论
    like.CommentID = ctx.FormValue("cid")
    //隶属于那篇文章
    like.ArticleID = ctx.FormValue("aid")
	error:= ReplyLikeRPC(like)
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
func (a *ArticlesService)DislikeReply(ctx iris.Context)(interface{},error){
	like:=new(model.Like)
	fmt.Println("调用了DislikeReply方法")
	//1.取消点赞的回复的ID
	like.ReplyID = ctx.FormValue("rid")
	//2.取消点赞回复的时间
	like.CreatedAt = time.Now().UnixNano()
	//3.取消点赞者的ID
	like.UserID = ctx.FormValue("uid")
	if len(strings.TrimSpace(ctx.FormValue("uid"))) ==0 {
		response :=response2.Response{}
		response.Code ="202"
		response.Message="用户id丢失"
		return response,nil
	}
	//回复的哪条评论
	like.CommentID=ctx.FormValue("cid")
	//隶属于那篇文章
	like.ArticleID = ctx.FormValue("aid")
	error:= ReplyDislikeRPC(like)
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
//内容安全
func ContentSecurityWorker(worker int, msg<- chan model.Article, result chan<- string){
	for article:=range msg{
		println(article.Images[0].ImageUrl)
		//图片鉴定
		profile := greensdksample.Profile{AccessKeyId:accessKeyId, AccessKeySecret:accessKeySecret}
		path := "/green/image/scan";
		clientInfo := greensdksample.ClinetInfo{Ip:"127.0.0.1"}
		// 构造请求数据
		bizType := "Green"
		scenes := []string{"porn"}
		task := greensdksample.Task{DataId:uuid2.Rand().Hex(), Url:"http://t2.hddhhn.com/uploads/tu/201701/8485/rzoms1ljmcb.jpg"}
		tasks := []greensdksample.Task{task}
		bizData := greensdksample.BizData{ bizType, scenes, tasks}
		var client greensdksample.IAliYunClient = greensdksample.DefaultClient{Profile:profile}
		result<-client.GetResponse(path, clientInfo, bizData)
	}
}
//（1）发表文章Article的RPC函数
func ArticleSaveRpc(article *model.Article)(error) {
	// Set up a connection to the server.
	print("ArticleSaveRPC")
	action :=&pb.ArticleAction{ActionValue:pb.ActionArticleType_ARTICLE_SAVE}
	//创建文章对象
	theArticle:=&pb.Article{}

	//1.articleID
    theArticle.ArticleID = article.ArticleID
	//2.文章的富文本内容
	theArticle.Content = article.Content
	//3.
	theArticle.CreatedTime = article.CreatedAt

	theArticle.UserID = article.UserID
	//文章的标题
	theArticle.ArticleTitle = article.Title
	//文章的图片
	theArticle.Images= article.Images



	articleType := &pb.ArticleType{}
	articleType.ArticleTypeID ="01010101"
	articleType.ArticleTypeName=[]byte("眼睛")
	articleType.ArticleTypeOnlineStatus =1
	articleType.ParentID= "000000"
	articleType.CreateTime = article.CreatedAt

	articleTypes:=[]*pb.ArticleType{}
	articleTypes = append(articleTypes,articleType)

	theArticle.ArticleTypes = articleTypes

	theArticles:=[]*pb.Article{}
	theArticles = append(theArticles,theArticle)
	theArticles = append(theArticles,theArticle)

	in := pb.DMLArticleRequest{Action:action,Articles:theArticles}
	r, err := rpc.ClientArticle.DMLArticle(rpc.ClientArticleCtx,&in)

	if err != nil {
		log.Println("could not greet: %v", err)
		//articleBuildErrors <-err
		return err

	}
	log.Printf("client: %s", r.Message)
print("success")
	return nil

}



//（2）更新文章Article的RPC函数
func ArticleUpdateRpc(article *model.Article)(error){
	action :=&pb.ArticleAction{ActionValue:pb.ActionArticleType_ARTICLE_MODIFY}
	//创建文章对象
	theArticle:=&pb.Article{}
	//文章的富文本内容
	theArticle.Content = article.Content
	//
	// //确定图片个数
	// //创建图片每一个对应的赋值
	// imageNum := len(article.Images)
	//
	//for i:=0; i<imageNum;i++  {
	//  theArticle.Images[i].ImageBase64 = article.Images[i].ImageBase64
	//}

	//
	//image:= &pb.ArticleImage{}
	//
	//image.ImageBase64 = article.Images[0].ImageBase64
	//
	//theArticle.Images = append(theArticle.Images,image)

	articleType := &pb.ArticleType{}
	articleType.ArticleTypeID =article.ArticleID
	articleType.ArticleTypeName=[]byte("眼睛")
	articleType.ArticleTypeOnlineStatus =1
	articleType.ParentID= "123456"
	articleType.CreateTime = article.CreatedAt
	articleTypes:=[]*pb.ArticleType{}
	articleTypes = append(articleTypes,articleType)
	theArticle.ArticleTypes = articleTypes
	theArticles:=[]*pb.Article{}
	theArticles = append(theArticles,theArticle)
	   in := pb.DMLArticleRequest{Action:action,Articles:theArticles}
	r, err := rpc.ClientArticle.DMLArticle(rpc.ClientArticleCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		//<-articleUpdateMessages
		return err
	}
	log.Printf("client: %s", r.Message)
	//<-articleUpdateMessages
	return nil

}

//（3）给文章发表评论的RPC
 func ArticleCommentRPC(comment *model.Comment)(error){
 	//commentMessages<- *comment

	 //1.初始化动作
	 action :=&pb.ArticleAction{ActionValue:pb.ActionArticleType_ARTICLE_COMMENT}
     //2.初始化结构体对象
	 commentType :=&pb.ArticleComment{}
     //3.赋值结构体属性传
     // (1)CID
     commentType.CommentID = comment.CommenterID
	 //(2)创建时间
	 commentType.CreateTime = comment.CreatedAt
	 //（3）更新时间
	 commentType.UpDateTime = comment.UpdatedAt
	 //(4)发表评论的人的ID
	 commentType.UserID = comment.CommenterID
	 commentTypes:=[]*pb.ArticleComment{}
	 commentTypes = append(commentTypes,commentType)
	 in := pb.DMLArticleCommentRequest{Action:action,Comments:commentTypes}
	 r, err := rpc.ClientArticle.DMLComment(rpc.ClientArticleCtx,&in)
	 if err != nil {
		 log.Println("could not greet: %v", err)
		 return err

	 }
	 log.Printf("client: %s", r.Message)
return nil
 }

//(4)回复评论RPC
func ArticleReplyRPC(reply *model.Reply)(error){
	//replyMessages<- *reply
	action :=&pb.ArticleAction{ActionValue:pb.ActionArticleType_ARTICLE_REPLY}
    //评论与回复共用一个请求体
	replyType :=&pb.ArticleComment{}
	//当为回复时候，把评论当成文章
	replyType.ArticleID = reply.CommentID
	//当为回复时候，把回复当成评论
	replyType.CommentID = reply.ReplyID
	//回复创建的时间
	replyType.CreateTime = reply.CreatedAt
	//回复更新的时间
	replyType.UpDateTime = reply.UpdatedAt
	//回复的内容
	replyType.Content = reply.Content
	//回复者的用户ID
	replyType.UserID = reply.ReplierID

	replyTypes:=[]*pb.ArticleComment{}
	replyTypes = append(replyTypes,replyType)
	in := pb.DMLArticleCommentRequest{Action:action,Comments:replyTypes}
	r, err := rpc.ClientArticle.DMLComment(rpc.ClientArticleCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		//<-replyMessages
		return err
	}
	log.Printf("client: %s", r.Message)
	//<-replyMessages
	return nil

}

//(5)文章点赞RPC
func ArticleLikeRPC(like *model.Like)(error){
	action :=&pb.ArticleAction{ActionValue:pb.ActionArticleType_ARTICLE_LIKE}
	likeType := &pb.ArticleLike{}
	//喜欢的文章的ID
	likeType.ArticleId =like.ArticleID
	//点赞的用户的ID
	likeType.UserName = like.UserID
	likeTypes:=[]*pb.ArticleLike{}
	likeTypes = append(likeTypes,likeType)
	in := pb.DMLArticleLikeRequest{Action:action,Likes:likeTypes}
	r, err := rpc.ClientArticle.DMLLike(rpc.ClientArticleCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err

	}
	log.Printf("client: %s", r.Message)
	//<-likeArticleMessages
	return nil
}
//(6)点赞评论
func CommentLikeRPC(like *model.Like)(error){

	action :=&pb.ArticleAction{ActionValue:pb.ActionArticleType_ARTICLE_COMMENT_LIKE}
	likeType := &pb.ArticleLike{}
	//点赞的评论的ID
	likeType.CommentId = like.CommentID
	//点赞人的UID
	likeType.UserId = like.UserID

	//点赞的时间

	likeTypes:=[]*pb.ArticleLike{}
	likeTypes = append(likeTypes,likeType)
	in := pb.DMLArticleLikeRequest{Action:action,Likes:likeTypes}
	r, err := rpc.ClientArticle.DMLLike(rpc.ClientArticleCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		//<-likeCommentMessages
		return err
	}
	log.Printf("client: %s", r.Message)
	//<-likeCommentMessages
	return nil
}

//(7)点赞回复
func ReplyLikeRPC(like *model.Like)(error){
	action :=&pb.ArticleAction{ActionValue:pb.ActionArticleType_ARTICLE_REPLY_LIKE}
	likeType :=&pb.ArticleLike{}
    likeType.UserId = like.UserID
    likeType.CommentId = like.ReplyID
	likeTypes:=[]*pb.ArticleLike{}
	likeTypes = append(likeTypes,likeType)
	in := pb.DMLArticleLikeRequest{Action:action,Likes:likeTypes}
	r, err := rpc.ClientArticle.DMLLike(rpc.ClientArticleCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err

	}
	log.Printf("client: %s", r.Message)
	return nil
}
//(8)取消评论点赞
func CommentDislikeRPC(like *model.Like)(error){

	action :=&pb.ArticleAction{ActionValue:pb.ActionArticleType_ARTICLE_COMMENT_DISLIKE}
	dislikeType :=&pb.ArticleLike{}
	//取消点赞评论的ID
	dislikeType.CommentId = like.CommentID

	//取消点赞者的用户ID
	dislikeType.UserId = like.UserID

	dislikeTypes:=[]*pb.ArticleLike{}
	dislikeTypes = append(dislikeTypes,dislikeType)

	in := pb.DMLArticleLikeRequest{Action:action,Likes:dislikeTypes}
	r, err := rpc.ClientArticle.DMLLike(rpc.ClientArticleCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}

//(9)取消文章点赞RPC
func ArticleDislikeRPC(like *model.Like)(error) {
	action :=&pb.ArticleAction{ActionValue:pb.ActionArticleType_ARTICLE_DISLIKE}
	dislikeType := &pb.ArticleLike{}
	//取消点赞哪篇文章
	dislikeType.ArticleId = like.ArticleID
	//取消点赞者的用户ID
	dislikeType.UserId = like.UserID
	dislikeTypes:=[]*pb.ArticleLike{}
	dislikeTypes = append(dislikeTypes,dislikeType)
	in := pb.DMLArticleLikeRequest{Action:action,Likes:dislikeTypes}
	r, err := rpc.ClientArticle.DMLLike(rpc.ClientArticleCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}

//(10)取消回复点赞RPC
func ReplyDislikeRPC(like *model.Like)(error){
	action :=&pb.ArticleAction{ActionValue:pb.ActionArticleType_ARTICLE_REPLY_DISLIKE}
	dislikeType := &pb.ArticleLike{}
	//取消点赞哪个回复
	dislikeType.CommentId = like.CommentID
	//取消点赞者的用户ID
	dislikeType.UserId = like.UserID
	dislikeTypes:=[]*pb.ArticleLike{}
	dislikeTypes = append(dislikeTypes,dislikeType)
	in := pb.DMLArticleLikeRequest{Action:action,Likes:dislikeTypes}
	r, err := rpc.ClientArticle.DMLLike(rpc.ClientArticleCtx,&in)
	if err != nil {
		log.Println("could not greet: %v", err)
		return err
	}
	log.Printf("client: %s", r.Message)
	return nil
}