//2018 August
//     (`・ω・´)
//Developed by WangYao - BlockChain & Go language backstage & iOS Developer
//The right of final interpretation is owned by the @MAYLOL

/*
				   _ooOoo_
				  o8888888o
				  88" . "88
				  (| -_- |)
				  O\  =  /O
			   ____/`---'\____
			 .'  \\|     |//  `.
			/  \\|||  :  |||//  \
		   /  _||||| -:- |||||-  \
		   |   | \\\  -  /// |   |
		   | \_|  ''\---/''  |   |
		   \  .-\__  `-`  ___/-. /
		 ___`. .'  /--.--\  `. . __
	  ."" '<  `.___\_<|>_/___.'  >'"".
	 | | :  `- \`.;`\ _ /`;.`/ - ` : | |
	 \  \ `-.   \_ __\ /__ _/   .-` /  /
======`-.____`-.___\_____/___.-`____.-'======
				   `=---='
^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
		 佛祖保佑       永无BUG
*/
package main

import (
	"github.com/kataras/iris"
    "merlot_project/merlot_write/articleworkerpool"
    "merlot_project/merlot_write/diaryworkerpool"
    "merlot_project/merlot_write/topicworkerpool"

	"merlot_project/merlot_write/adaptor"
)

const (
	address2     = "172.16.3.168:50051"
)
func main() {
	//iris初始化
	app:=iris.Default()
	adaptor.Run(app)
	//创建发送文章相关的工作池
	articleDispatcher :=articleWorkerPool.NewDispatcher(10)
	articleDispatcher.Run()

   //创建发表主题相关的任务工作吃
	topicDispatcher :=topicWorkerPool.NewDispatcher(10)
	topicDispatcher.Run()

  //创建发表日记相关的工作池
	diaryDispatcher :=diaryWorkerPool.NewDispatcher(10)
	diaryDispatcher.Run()


	app.Run(iris.Addr(":8081"))
	//调用缓冲RPC服务
	//Rpc()
}






//敏感词过滤
func WordFilter(){
	//加载读取敏感词
     updateSensitiveWord()
     //通过循环匹配是否有敏感词汇
	 //发现有敏感词汇以后，将返回
}
    //敏感词库是否需要更新
func  updateSensitiveWord(){
	//查询敏感词库的版本，当发现词库需要更新，则下载词库
}


//鉴图（是否健康）阿里云API
func ImageAppraisal(){
//查看阿里巴巴API文档
}

//视频鉴定
func VideoAppraisal(){
//人工和其他解决方案
}









//将已经审核通过的文章打标签放在合格文章的切片中传递给数据缓冲层
func GRpc(){

}




