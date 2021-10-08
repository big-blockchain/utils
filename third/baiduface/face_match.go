package baiduface

// 人脸对比
// http://ai.baidu.com/docs#/Face-Match-V3/top
type Match struct {
	Score    float64 `json:"score"`
	FaceList []struct {
		FaceToken string `json:"face_token"`
	} `json:"face_list"`
}

type MatchParam struct {
	ImageType       ImageType       `json:"image_type"`
	Image           string          `json:"image"`
	FaceType        FaceType        `json:"face_type"`
	QualityControl  QualityControl  `json:"quality_control"`
	LivenessControl LivenessControl `json:"liveness_control"`
}

// 图片搜索
//
// image	是	string	图片信息(总数据大小应小于10M)，图片上传方式根据image_type来判断
// image_type	是	string	图片类型
//  BASE64:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M；
//	URL:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)；
//	FACE_TOKEN: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
//quality_control	否	string	图片质量控制
//	NONE: 不进行控制
//	LOW:较低的质量要求
//	NORMAL: 一般的质量要求
//	HIGH: 较高的质量要求
//	默认 NONE
//	若图片质量不满足要求，则返回结果中会提示质量检测失败
//liveness_control	否	string	活体检测控制
//	NONE: 不进行控制
//	LOW:较低的活体要求(高通过率 低攻击拒绝率)
//	NORMAL: 一般的活体要求(平衡的攻击拒绝率, 通过率)
//	HIGH: 较高的活体要求(高攻击拒绝率 低通过率)
//	默认NONE
//	若活体检测结果不满足要求，则返回结果中会提示活体检测失败
func (client *Client) Match(param []MatchParam) (Match, string, error) {

	res := Match{}
	str, err := client.reqPost("match", param, &res)

	if err != nil {
		return res, str, err
	}

	return res, str, nil
}
