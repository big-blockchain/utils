/**
 * @Auth: Nuts
 * @Date: 2021/3/3 6:48 下午
 */
package baiduface

import "time"

type FaceUserParam struct {
	UserId    string    `json:"user_id"`
	GroupId   string    `json:"group_id"`
	Image     string    `json:"image"`
	ImageType ImageType `json:"image_type"`
}

type FaceUserResult struct {
	FaceToken string      `json:"face_token"` //人脸标志
	Location  interface{} `json:"location"`   //匹配的用户信息列表
}

type Location struct {
	Left     int64 `json:"left"`
	Top      int64 `json:"top"`
	Width    int64 `json:"width"`
	Height   int64 `json:"height"`
	Rotation int64 `json:""`
}

func (client *Client) FaceSet(faceUid, image, groupId string, imageType ImageType) (FaceUserResult, string, error) {
	param := FaceUserParam{
		Image:     image,
		ImageType: imageType,
		GroupId:   groupId,
		UserId:    faceUid,
	}

	body := make(map[string]interface{})
	body["image_type"] = imageType
	body["image"] = image
	if groupId != "" {
		body["group_id"] = groupId
	}
	if faceUid != "" {
		body["user_id"] = faceUid
	}

	res := FaceUserResult{}
	str, err := client.reqPost("faceset/user/add", param, &res)

	if err != nil {
		if err.Error() == "Open api qps request limit reached" { //超出QPS限制时，休眠0.5秒再查
			time.Sleep(500 * time.Millisecond)
			return client.FaceSet(faceUid, image, groupId, imageType)
		}
		return res, str, err
	}
	return res, str, nil
}
