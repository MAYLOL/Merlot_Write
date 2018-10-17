package model
type Video struct {
	//视频名
	Name  string `json:"name"`
	//视频的ID
	VideoID string `json:"video_id"`
	//视频是谁发的
	UserID  string    `json:"user_id"`
	//视频审核状态
	State   int64 `json:"state"`
	//视频属于哪个文章
	ArticleID string `json:"article_id"`
	//视频的宽高比
	Ratio float64 `json:"ratio"`
	//视频URL
	VideoUrl string `json:"video_url"`
    //图片base64
    VideoBase64 string `json:"video_base_64"`
    //视频类型
    VideoType  string `json:"video_type"`
    //视频预览图
    VideoImage string `json:"video_image"`
    //视频不过审核的原因
    Cause string `json:"cause"`
}

const (
	//VideoVerifying正在审核
	VideoVerifying = 1
	//VideoVerifySuccess审核成功
	VideoVerifySuccess = 2
	//VideoVerifyFail 审核失败
	VideoVerifyFali = 3
)
