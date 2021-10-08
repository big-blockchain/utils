package baiduface

type VerifyParam struct {
	ImageType ImageType `json:"image_type"`
	Image     string    `json:"image"`
	FaceField string    `json:"face_field"`
}

type VerifyResult struct {
	FaceLiveness float64 `json:"face_liveness"`
}

func (client *Client) Verify(param []VerifyParam) (VerifyResult, string, error) {

	res := VerifyResult{}
	str, err := client.reqPost("faceverify", param, &res)

	if err != nil {
		return res, str, err
	}
	return res, str, nil
}
