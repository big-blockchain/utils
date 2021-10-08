package baiduface

// 人脸检测
// http://ai.baidu.com/docs#/Face-Detect-V3/top
type FaceDetect struct {
	FaceNum  int              `json:"face_num"`  //检测到的图片中的人脸数量
	FaceList []FaceDetectList `json:"face_list"` //人脸信息列表，具体包含的参数参考下面的列表。
}

type FaceDetectList struct {
	FaceToken       string       `json:"face_token"`       //人脸图片的唯一标识
	Location        FaceLocation `json:"location"`         //人脸在图片中的位置
	FaceProbability float64      `json:"face_probability"` //人脸置信度，范围【0~1】，代表这是一张人脸的概率，0最小、1最大。
	Angel           struct {
		Yaw   float64 `json:"yaw"`   //三维旋转之左右旋转角[-90(左), 90(右)]
		Pitch float64 `json:"pitch"` //三维旋转之俯仰角度[-90(上), 90(下)]
		Roll  float64 `json:"roll"`  //平面内旋转角[-180(逆时针), 180(顺时针)
	} `json:"angel"` //人脸旋转角度参数
	Age        float64         `json:"age"`        // 年龄
	Beauty     float64         `json:"beauty"`     //美丑打分，范围0-100，越大表示越美。当face_fields包含beauty时返回
	Expression TypeProbability `json:"expression"` //表情，当 face_field包含expression时返回 type none:不笑；smile:微笑；laugh:大笑
	FaceShape  TypeProbability `json:"face_shape"` //脸型，当face_field包含face_shape时返回 type square: 正方形 triangle:三角形 oval: 椭圆 heart: 心形 round: 圆形
	Gender     TypeProbability `json:"gender"`     //性别，face_field包含gender时返回 type male:男性 female:女性
	Glasses    TypeProbability `json:"glasses"`    //是否带眼镜，face_field包含glasses时返回,type none:无眼镜，common:普通眼镜，sun:墨镜
	EyeStatus  struct {
		LeftEye  float64 `json:"left_eye"` //左眼状态 [0,1]取值，越接近0闭合的可能性越大
		RightEye float64 `json:"right_eye"`
	} `json:"eye_status"` //双眼状态（睁开/闭合） face_field包含eye_status时返回
	Emotion     TypeProbability `json:"emotion"`   //情绪 face_field包含emotion时返回 type angry:愤怒 disgust:厌恶 fear:恐惧 happy:高兴 sad:伤心 surprise:惊讶 neutral:无情绪
	Race        TypeProbability `json:"race"`      //人种 face_field包含race时返回 type ,yellow: 黄种人 white: 白种人 black:黑种人 arabs: 阿拉伯人
	FaceType    TypeProbability `json:"face_type"` //真实人脸/卡通人脸 face_field包含face_type时返回,type,human: 真实人脸 cartoon: 卡通人脸
	Landmark    []Landmark      `json:"landmark"`  //4个关键点位置，左眼中心、右眼中心、鼻尖、嘴中心。face_field包含landmark时返回
	Landmark72  []Landmark      `json:"landmark72"`
	Landmark150 []Landmark      `json:"landmark150"`
	Quality     struct {
		Occlusion struct {
			LeftEye    float64 `json:"left_eye"`    //左眼遮挡比例，[0-1] ，1表示完全遮挡
			RightEye   float64 `json:"right_eye"`   //右眼遮挡比例，[0-1] ， 1表示完全遮挡
			Nose       float64 `json:"nose"`        //鼻子遮挡比例，[0-1] ， 1表示完全遮挡
			Mouth      float64 `json:"mouth"`       //嘴巴遮挡比例，[0-1] ， 1表示完全遮挡
			LeftCheek  float64 `json:"left_cheek"`  //左脸颊遮挡比例，[0-1] ， 1表示完全遮挡
			RightCheek float64 `json:"right_cheek"` //右脸颊遮挡比例，[0-1] ， 1表示完全遮挡
			Chin       float64 `json:"chin"`        //下巴遮挡比例，，[0-1] ， 1表示完全遮挡
		} `json:"occlusion"` //人脸各部分遮挡的概率，范围[0~1]，0表示完整，1表示不完整
		Blur         float64 `json:"blur"`         //人脸模糊程度，范围[0~1]，0表示清晰，1表示模糊
		Illumination int64   `json:"illumination"` //取值范围在[0~255], 表示脸部区域的光照程度 越大表示光照越好
		Completeness int64   `json:"completeness"` //人脸完整度，0或1, 0为人脸溢出图像边界，1为人脸都在图像边界内
	} `json:"quality"` //人脸质量信息。face_field包含quality时返回
}

type TypeProbability struct {
	Type        string  `json:"type"`        // 类型
	Probability float64 `json:"probability"` //  可信度
}

type Landmark struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type FaceLocation struct {
	Left     float64 `json:"left"`     //人脸区域离左边界的距离
	Top      float64 `json:"top"`      //人脸区域离上边界的距离
	Width    float64 `json:"width"`    // 人脸区域的宽度
	Height   float64 `json:"height"`   // 人脸区域的高度
	Rotation int64   `json:"rotation"` //人脸框相对于竖直方向的顺时针旋转角，[-180,180]
}

//人脸检查
//
//image	 	string	图片信息(总数据大小应小于10M)，图片上传方式根据image_type来判断
//image_type	string	图片类型
//  BASE64:图片的base64值，base64编码后的图片数据，编码后的图片大小不超过2M；
//	URL:图片的 URL地址( 可能由于网络等原因导致下载图片时间过长)；
//	FACE_TOKEN: 人脸图片的唯一标识，调用人脸检测接口时，会为每个人脸图片赋予一个唯一的FACE_TOKEN，同一张图片多次检测得到的FACE_TOKEN是同一个。
//face_field		string	包括age,beauty,expression,face_shape,gender,glasses,landmark,landmark150,race,quality,eye_status,emotion,face_type信息
//	逗号分隔. 默认只返回face_token、人脸框、概率和旋转角度
//max_face_num		uint32	最多处理人脸的数目，默认值为1，仅检测图片中面积最大的那个人脸；最大值10，检测图片中面积最大的几张人脸。
//face_type		string	人脸的类型
//	LIVE表示生活照：通常为手机、相机拍摄的人像图片、或从网络获取的人像图片等
//	IDCARD表示身份证芯片照：二代身份证内置芯片中的人像照片
//	WATERMARK表示带水印证件照：一般为带水印的小图，如公安网小图
//	CERT表示证件照片：如拍摄的身份证、工卡、护照、学生证等证件图片
//		默认LIVE
//liveness_control	否	string	活体控制 检测结果中不符合要求的人脸会被过滤
//	NONE: 不进行控制
//	LOW:较低的活体要求(高通过率 低攻击拒绝率)
//	NORMAL: 一般的活体要求(平衡的攻击拒绝率, 通过率)
//	HIGH: 较高的活体要求(高攻击拒绝率 低通过率)
//		默认NONE
func (client *Client) Detect(image string, imageType ImageType, faceField string, faceType FaceType, livenessControl LivenessControl, maxFaceNum uint32) (FaceDetect, string, error) {

	body := make(map[string]interface{})
	body["image_type"] = imageType
	body["image"] = image
	if faceField != "" {
		body["face_field"] = faceField
	}
	if faceType != "" {
		body["face_type"] = faceType
	}

	if livenessControl != "" {
		body["liveness_control"] = livenessControl
	}

	if maxFaceNum > 0 {
		body["max_face_num"] = maxFaceNum
	}

	res := FaceDetect{}
	str, err := client.reqPost("detect", body, &res)

	if err != nil {
		return res, str, err
	}

	return res, str, nil
}
