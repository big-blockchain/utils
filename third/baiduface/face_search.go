package baiduface

import "time"

// 人脸搜索
// 文档地址：http://ai.baidu.com/docs#/Face-Search-V3/top
type FaceSearch struct {
	FaceToken string         `json:"face_token"` //人脸标志
	UserList  []FaceUserList `json:"user_list"`  //匹配的用户信息列表
}

type FaceUserList struct {
	GroupId  string  `json:"group_id"`  //用户所属的group_id
	UserId   string  `json:"user_id"`   //用户的user_id
	UserInfo string  `json:"user_info"` //注册用户时携带的user_info
	Score    float64 `json:"score"`     //用户的匹配得分，推荐阈值80分
}

// 图片搜索
//
// image	是	string	图片信息(总数据大小应小于10M)，图片上传方式根据image_type来判断
// image_type	是	string	图片类型
//  BASE64:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M；
//	URL:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)；
//	FACE_TOKEN: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
//group_id_list	是	string	从指定的group中进行查找 用逗号分隔，上限10个
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
//user_id	否	string	当需要对特定用户进行比对时，指定user_id进行比对。即人脸认证功能。
//max_user_num	否	unit32	查找后返回的用户数量。返回相似度最高的几个用户，默认为1，最多返回50个。
func (client *Client) Search(image string, imageType ImageType, groupIdList string, qualityControl QualityControl, livenessControl LivenessControl, userId string, maxUserNum uint32) (FaceSearch, string, error) {

	body := make(map[string]interface{})
	body["image_type"] = imageType
	body["image"] = image
	if groupIdList != "" {
		body["group_id_list"] = groupIdList
	}
	if qualityControl != "" {
		body["quality_control"] = qualityControl
	}
	if userId != "" {
		body["user_id"] = userId
	}

	if livenessControl != "" {
		body["liveness_control"] = livenessControl
	}

	if maxUserNum > 0 {
		body["max_user_num"] = maxUserNum
	}

	res := FaceSearch{}
	str, err := client.reqPost("search", body, &res)

	if err != nil {
		if err.Error() == "Open api qps request limit reached" { //超出QPS限制时，休眠0.5秒再查
			time.Sleep(300 * time.Millisecond)
			return client.Search(image, imageType, groupIdList, qualityControl, livenessControl, userId, maxUserNum)
		}
		return res, str, err
	}

	return res, str, nil
}
