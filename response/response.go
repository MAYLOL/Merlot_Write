package response
//针对内容请求类型的接口
type Response struct {
	Code string  `json:"code"`
	Message  string  `json:"message"`
	Data Data  `json:"data"`
}

type Data struct {

}