package model
type Image struct {
	//图片名
	Name  string `json:"name"`
	//图片是谁发的
    UserID   string  `json:"user_id"`
	//是图片审核状态
	Status   int64 `json:"status"`
	//图片属于文章的ID
	ArticleID int64 `json:"article_id"`
	//图片的宽高比
	Ratio float64 `json:"ratio"`
	//图片编码格式（图片类型）
	ImageType string `json:"image_type"`
	//图片URL
	ImageUrl string `json:"image_url"`
	//图片base64
	ImageBase64 string `json:"image_base_64"`
	//图片不过审核的原因
	Cause  string  `json:"cause"`
}
const (
	//ImageVerifying正在审核
	ImageVerifying = 1
	//ImageVerifySuccess审核成功
	ImageVerifySuccess = 2
	//ImageVerifyFail审核失败
	ImageVerifyFali = 3
)
