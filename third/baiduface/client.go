package baiduface

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	BaseUrl   = "https://aip.baidubce.com"
	tokenUrl  = BaseUrl + "/oauth/2.0/token"
	faceUrl   = BaseUrl + "/rest/2.0/face/v3"
	filterUrl = BaseUrl + "/rest/2.0/solution/v1"
)

//图片类型
type ImageType string

const (
	ImageTypeUrl       ImageType = "URL"
	ImageTypeBase64    ImageType = "BASE64"
	ImageTypeFaceToken ImageType = "FACE_TOKEN"
)

//人脸的类型
type FaceType string

const (
	FaceTypeLIVE      FaceType = "LIVE"      //表示生活照：通常为手机、相机拍摄的人像图片、或从网络获取的人像图片等
	FaceTypeIDCARD    FaceType = "IDCARD"    //表示身份证芯片照：二代身份证内置芯片中的人像照片，
	FaceTypeWATERMARK FaceType = "WATERMARK" //表示带水印证件照：一般为带水印的小图，如公安网小图
	FaceTypeCERT      FaceType = "CERT"      //表示证件照片：如拍摄的身份证、工卡、护照、学生证等证件图片
)

//图片质量控制
type QualityControl string

const (
	QualityControlNONE   QualityControl = "NONE"   //不进行控制
	QualityControlLOW    QualityControl = "LOW"    //较低的质量要求
	QualityControlNORMAL QualityControl = "NORMAL" //一般的质量要求
	QualityControlHIGH   QualityControl = "HIGH"   //较高的质量要求
)

//活体检测控制
type LivenessControl string

const (
	LivenessControlNONE   QualityControl = "NONE"   //不进行控制
	LivenessControlLOW    QualityControl = "LOW"    //较低的质量要求
	LivenessControlNORMAL QualityControl = "NORMAL" //一般的质量要求
	LivenessControlHIGH   QualityControl = "HIGH"   //较高的质量要求
)

type Error struct {
	ErrorCode int    `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
}

type Result struct {
	*Error
	Result interface{} `json:"result"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	*Error
}

type Client struct {
	apiKey      string
	secretKey   string
	AccessToken string
	ExpiresTime time.Time
}

func NewClient(apiKey, secretKey string) *Client {
	return &Client{
		apiKey:    apiKey,
		secretKey: secretKey,
	}
}

// 获取AccessToken
func (client *Client) getAccessToken() string {

	if client.AccessToken != "" {
		if time.Now().Before(client.ExpiresTime) {
			return client.AccessToken
		}
	}

	con := &http.Client{}

	req1, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s?grant_type=client_credentials&client_id=%s&client_secret=%s", tokenUrl, client.apiKey, client.secretKey), nil)

	req1.Header.Set("Content-Type", "application/json")
	resp, err := con.Do(req1)

	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	bodys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	res := AccessToken{}
	err = json.Unmarshal(bodys, &res)
	if err != nil {

	}

	client.AccessToken = res.AccessToken
	client.ExpiresTime = time.Now().Add(time.Second * time.Duration(res.ExpiresIn))

	return client.AccessToken
}

// url 请求地址
// param 请求参数
// response 响应的结构体
// 返回参数 字符串为当前
func (client *Client) reqPost(url string, param interface{}, response interface{}) (string, error) {

	con := &http.Client{}

	byts, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	body := ioutil.NopCloser(bytes.NewReader(byts))
	urlStr := fmt.Sprintf("%s/%s?access_token=%s", faceUrl, url, client.getAccessToken())
	req1, _ := http.NewRequest(http.MethodPost, urlStr, body)

	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("Content-Length", fmt.Sprintf("%d", int64(len(byts))))
	resp, err := con.Do(req1)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	res := Result{}

	err = json.Unmarshal(bodys, &res)

	if err != nil {
		return string(bodys), err
	}

	if res.Error != nil {
		if res.ErrorCode > 0 {
			return string(bodys), errors.New(res.ErrorMsg)
		}
	}

	jsonBody, err := json.Marshal(&res.Result)
	if err != nil {
		return string(bodys), err
	}

	err = json.Unmarshal(jsonBody, &response)

	return string(bodys), err
}

// url 请求地址
// param 请求参数
// response 响应的结构体
// 返回参数 字符串为当前
func (client *Client) reqFilterPost(urlStr string, params map[string]interface{}, response interface{}) (string, error) {
	cli := &http.Client{}

	//post要提交的数据
	DataUrlVal := url.Values{}
	for key, val := range params {
		DataUrlVal.Add(key, fmt.Sprintf("%v", val))
	}
	urlStr = fmt.Sprintf("%s/%s?access_token=%s", filterUrl, urlStr, client.getAccessToken())
	req, err := http.NewRequest("POST", urlStr, strings.NewReader(DataUrlVal.Encode()))
	if err != nil {
		return "", err
	}
	//伪装头部
	byts, err := json.Marshal(params)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", fmt.Sprintf("%d", int64(len(byts))))
	//提交请求
	resp, err := cli.Do(req)
	defer resp.Body.Close()


	if err != nil {
		return "", err
	}

	bodys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}


	err = json.Unmarshal(bodys, &response)

	if err != nil {
		return string(bodys), err
	}

	return string(bodys), err
}

// url 请求地址
// param 请求参数
// response 响应的结构体
// 返回参数 字符串为当前
func (client *Client) reqSetFacePost(url string, param interface{}, response interface{}) (string, error) {

	con := &http.Client{}

	byts, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	body := ioutil.NopCloser(bytes.NewReader(byts))

	req1, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/%s?access_token=%s", faceUrl, url, client.getAccessToken()), body)

	req1.Header.Set("Content-Type", "application/json")
	req1.Header.Set("Content-Length", fmt.Sprintf("%d", int64(len(byts))))
	resp, err := con.Do(req1)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(bodys, &response)

	if err != nil {
		return string(bodys), err
	}

	return string(bodys), err
}
