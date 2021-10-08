package baiduface

type IdentifyParam struct {
	Image      string `json:"image"`
	GroupId    string `json:"group_id"`
	UserTopNum int64  `json:"user_top_num"`
}

type IdentifyResult struct {
	GroupId  string    `json:"group_id"`
	Uid      string     `json:"uid"`
	UserInfo string    `json:"user_info"`
	Score    []float64 `json:"score"`
}

func (client *Client) Identify(image, groupId string) (IdentifyResult, string, error) {
	param := IdentifyParam{
		Image:      image,
		GroupId:    groupId,
		UserTopNum: 1,
	}

	res := IdentifyResult{}
	str, err := client.reqPost("identify", param, &res)

	if err != nil {
		return res, str, err
	}
	return res, str, nil
}
