package adaptor

import (
	"github.com/kataras/iris"
	"merlot_project/merlot_write/controller"
	response2 "merlot_project/merlot_write/response"
	"os"
	"fmt"
	"io/ioutil"


)

//func empty(){
//	gopherjspkg.FS
//	dom.AnimationEvent{}
//	dom.ImageData{}
//
//}
//注册并运行所有接口
func Run(app* iris.Application){
	//文章相关接口
   RunArticles(app)
   //日记相关接口
   RunDiaries(app)
   //主题相关接口
   RunTopics(app)
	 //ImageRouter
	 // ImageRouter(app)
}
//一、发布文章
func RunArticles(app *iris.Application){
	//1.发文章
	SendArticle(app)
	//2.更新文章
	UpDateArticle(app)
	//3.发表评论
	CommentArticle(app)
	//4.回复文章评论
	ReplyArticle(app)
	//5.点赞文章
	LikeArticle(app)
	//6.取消赞文章
	DisLikeArticle(app)
	//7.点赞文章评论
	LikeArticleComment(app)
	//8.取消评论点赞
	DisLikeArticleComment(app)
	//9.回复点赞
	LikeArticleReply(app)
	//10.取消回复点赞
	DisLikeArticleReply(app)
}




//二、发布日记
func RunDiaries(app *iris.Application){
	//1.发日记
	SendDiary(app)
	//2.更新日记
	UpdateDiary(app)
	//3.发表评论
	CommentDiary(app)
	//4.回复日记评论
	ReplyDiary(app)
	//5.点赞日记
	LikeDiary(app)
	//6.取消点赞日记
	DislikeDiary(app)
	//7.点赞日记评论
	LikeDiaryComment(app)
	//8.取消日记评论点赞
	DisLikeDiaryComment(app)
	//9.点赞日记回复
	LikeDiaryReply(app)
	//10.取消日记回复的赞
	DisLikeDiaryReply(app)
}

//运行话题所有接口
func RunTopics(app *iris.Application){
//1.发表话题
	SendTopic(app)
//2.更新话题
	UpdateTopic(app)
//3.评论话题
	CommentTopic(app)
//4.回复话题评论
    ReplyTopic(app)
//5.点赞话题
    LikeTopic(app)
//6.取消话题点赞
    DisLikeTopic(app)
//7.点赞话题评论
    LikeTopicComment(app)
//8.取消话题评论点赞
    DisLikeTopicComment(app)
//9.点赞话题的回复
    LikeTopicReply(app)
//10.取消回复的点赞
    DislikeTopicReply(app)
}

//三、发表话题

//1.发表话题
func SendTopic(app *iris.Application){
	app.Handle("POST","/topics", func(ctx iris.Context) {
		top :=article.TopicsService{}
		msg,err:= top.Topic(ctx)
	    data:=new([]interface{})
		responseFunc(ctx,msg,data,err)
	})
}
//2.更新话题
func UpdateTopic(app *iris.Application){
	app.Handle("PUT","/topics", func(ctx iris.Context) {
		top :=article.TopicsService{}
		msg,err:= top.UpdateTopic(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//3.评论话题
func CommentTopic(app *iris.Application){
	app.Handle("POST","/topics/answer", func(ctx iris.Context) {
		top :=article.TopicsService{}
		msg,err:= top.CommentTopic(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})

}
//4.回复话题评论
func ReplyTopic(app *iris.Application){
	app.Handle("POST","/topics/reply", func(ctx iris.Context) {
		top :=article.TopicsService{}
		msg,err:= top.ReplyComment(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}

//5.点赞话题
func LikeTopic(app *iris.Application){
	app.Handle("PUT","/topics/like", func(ctx iris.Context) {
		top :=article.TopicsService{}
		msg,err:= top.LikeTopic(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}

//6.取消话题点赞
func DisLikeTopic(app *iris.Application){
	app.Handle("PUT","/topics/dislike", func(ctx iris.Context) {
		top :=article.TopicsService{}

		msg,err:= top.DislikeTopic(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//7.点赞话题回答
func LikeTopicComment(app *iris.Application){
	app.Handle("PUT","/topics/answer/like", func(ctx iris.Context) {
		top :=article.TopicsService{}
		msg,err:= top.LikeComment(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}

//8.取消话题回答点赞
func DisLikeTopicComment(app *iris.Application){
	app.Handle("PUT","/topics/answer/dislike", func(ctx iris.Context) {
		top :=article.TopicsService{}
		msg,err:= top.DislikeComment(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//9.点赞话题的回复
func LikeTopicReply(app *iris.Application){
	app.Handle("PUT","/topics/reply/like", func(ctx iris.Context) {
		top :=article.TopicsService{}
		msg,err:= top.LikeReply(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}

//10.取消回复的点赞
func DislikeTopicReply(app *iris.Application){
	app.Handle("PUT","/topics/reply/dislike", func(ctx iris.Context) {
		top :=article.TopicsService{}
		msg,err:= top.DislikeReply(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//1.发日记
func SendDiary(app *iris.Application){

	app.Handle("POST","/diaries", func(ctx iris.Context) {
		di :=article.DiariesServices{}
		msg,err:= di.Diary(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)

	})

}

//2.更新日记
func UpdateDiary(app *iris.Application){
	app.Handle("PUT","/diaries", func(ctx iris.Context) {
		di :=article.DiariesServices{}
		msg,err:= di.UpdateDiary(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//3.发表评论
func CommentDiary(app *iris.Application){
	app.Handle("PUT","/diaries/comment", func(ctx iris.Context) {
		di :=article.DiariesServices{}
		msg,err:= di.CommentDiary(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//4.回复日记评论
func ReplyDiary(app *iris.Application){
	app.Handle("PUT","/diaries/reply", func(ctx iris.Context) {
		di :=article.DiariesServices{}
		msg,err:= di.ReplyComment(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}

//5.点赞日记
func LikeDiary(app *iris.Application){
	app.Handle("PUT","/diaries/like", func(ctx iris.Context) {
		di :=article.DiariesServices{}
		msg,err:= di.LikeDiary(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//6.取消点赞日记
func DislikeDiary(app *iris.Application){
	app.Handle("PUT","/diaries/dislike", func(ctx iris.Context) {
		di :=article.DiariesServices{}
		msg,err:= di.DislikeDiary(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//7.点赞日记评论
func LikeDiaryComment(app *iris.Application){
	app.Handle("PUT","/diaries/comment/like", func(ctx iris.Context) {
		di :=article.DiariesServices{}
		msg,err:= di.LikeComment(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//8.取消日记评论点赞
func DisLikeDiaryComment(app *iris.Application){
	app.Handle("PUT","/diaries/comment/dislike", func(ctx iris.Context) {
		di :=article.DiariesServices{}
		msg,err:= di.DislikeComment(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//9.日记的回复点赞
func LikeDiaryReply(app *iris.Application){
	app.Handle("PUT","/diaries/reply/like", func(ctx iris.Context) {
		di :=article.DiariesServices{}
		msg,err:= di.LikeReply(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}

//10.取消日记的回复点赞
func DisLikeDiaryReply(app *iris.Application)  {
	app.Handle("PUT","/diaries/reply/dislike", func(ctx iris.Context) {
		di :=article.DiariesServices{}
		msg,err:= di.DislikeReply(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
func responseFunc(ctx iris.Context,msg interface{},data interface{},err error){
	theResponse:=response2.Response{}
	if err !=nil {
		theResponse.Code="1"
		theResponse.Message =err.Error()
		theResponse.Data = response2.Data{}
		ctx.JSON(err)

	}else{

		ctx.JSON(msg)
	}


}
//一、发布文章
//1.发布文章（支持图片、视频、12）
func SendArticle(app *iris.Application){
	app.Handle("POST","/articles", func(ctx iris.Context) {
		ar := article.ArticlesService{}

		msg,err:=  ar.Article(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//2.更新文章
func UpDateArticle(app *iris.Application){
	app.Handle("PUT","/articles", func(ctx iris.Context) {
		ar := article.ArticlesService{}
		msg,err:=  ar.UpdateArticle(ctx)
		data:="{}"
		responseFunc(ctx,msg,data,err)
	})
}
//3.评论文章
func CommentArticle(app *iris.Application){
	app.Handle("POST","/articles/comment", func(ctx iris.Context) {
		ar := article.ArticlesService{}
		msg,err:=  ar.CommentArtcle(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}

//4.回复文章
func ReplyArticle(app *iris.Application){
	app.Handle("POST","/articles/reply", func(ctx iris.Context) {
		ar := article.ArticlesService{}
		msg,err:=  ar.ReplyComment(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//5.点赞文章
func LikeArticle(app *iris.Application){
	app.Handle("PUT","/articles/like", func(ctx iris.Context) {

		ar := article.ArticlesService{}
		msg,err:=  ar.LikeArticle(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//6.取消文章点赞
func DisLikeArticle(app *iris.Application){
	app.Handle("PUT","/articles/dislike", func(ctx iris.Context) {

		ar := article.ArticlesService{}
		msg,err:=  ar.DislikeArticle(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)

	})
}
//7.评论点赞
func LikeArticleComment(app  *iris.Application){
	app.Handle("PUT","/articles/comment/like", func(ctx iris.Context) {
		ar := article.ArticlesService{}
		msg,err:=  ar.LikeComment(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}

//8.取消评论点赞
func DisLikeArticleComment(app *iris.Application){
	app.Handle("PUT","/articles/comment/dislike", func(ctx iris.Context) {
		ar := article.ArticlesService{}
		msg,err:=  ar.DislikeComment(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}

//9.回复的点赞
func LikeArticleReply(app *iris.Application){
	app.Handle("PUT","/articles/reply/like", func(ctx iris.Context) {
		ar := article.ArticlesService{}
		msg,err:=  ar.LikeReply(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}
//10.取消恢复的点赞
func DisLikeArticleReply(app *iris.Application){
	app.Handle("PUT","/articles/reply/dislike", func(ctx iris.Context) {
		ar := article.ArticlesService{}
		msg,err:=  ar.DislikeReply(ctx)
		data:=new([]interface{})
		responseFunc(ctx,msg,*data,err)
	})
}






//访问服务器图片接口(拼接图片ID可获得图片)

func ImageRouter(app *iris.Application){

	app.Handle("GET","/Resources/Images", func(ctx iris.Context) {
		var ID =  ctx.FormValue("ID")

		file,err:=os.Open(ID)
		if err!=nil {
			fmt.Println(err)
		}
		defer file.Close()
		buff,_:=ioutil.ReadAll(file)
		ctx.Write(buff)
	})
}