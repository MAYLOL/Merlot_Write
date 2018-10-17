package model
//调度池

type DispatchPool struct {
	//开多少个协程goroutine去处理
	WorkerNum  int64
	//消息缓冲管道的缓冲数量
	MsgNum  int64
	//结果返回的缓冲数量
	ResultNum int64


}
