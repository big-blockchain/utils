/**
 * @Auth: Nuts
 * @Date: 2021/6/21 2:56 下午
即构直播
权限相关
*/

package zego

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type ZegoClient struct {
	SecretId   int
	SecretKey  string
	SecretSign string
}

func NewClient(secretId int, secretKey, secretSign string) *ZegoClient {
	return &ZegoClient{
		SecretId:   secretId,
		SecretKey:  secretKey,
		SecretSign: secretSign,
	}
}

/**
获取服务器 access_token，用于调用服务端 API 时进行鉴权。
*/
func (cli ZegoClient) GetAccessToken() (accessToken string, err error) {
	nonce := fmt.Sprintf("%d", rand.Int63())
	expired := time.Now().Unix() + 3600

	m := md5.New()
	m.Write([]byte(fmt.Sprintf("%d%s%s%d", cli.SecretId, strings.ToLower(cli.SecretKey), nonce, expired)))
	hash := fmt.Sprintf("%x", m.Sum(nil))

	var tokenInfo struct {
		Ver     int    `json:"ver"`
		Hash    string `json:"hash"`
		Nonce   string `json:"nonce"`
		Expired int    `json:"expired"`
	}
	tokenInfo.Hash = hash
	tokenInfo.Expired = int(expired)
	tokenInfo.Ver = 1
	tokenInfo.Nonce = nonce
	tokenInfoStr, _ := json.Marshal(tokenInfo)

	var reqStruct struct {
		SecretId int    `json:"secret_id"`
		Token    string `json:"token"`
	}
	reqStruct.SecretId = cli.SecretId
	reqStruct.Token = base64.StdEncoding.EncodeToString(tokenInfoStr)
	reqBody, _ := json.Marshal(reqStruct)

	req, err := http.NewRequest(http.MethodPost, "https://roomkit-api.zego.im/auth/get_access_token", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var ret struct {
		Ret struct {
			Code    int    `json:"code"`
			Msg     string `json:"msg"`
			Version string `json:"version"`
		} `json:"ret"`
		Data struct {
			AccessToken string `json:"access_token"`
			ExpiresIn   int    `json:"expires_in"`
		}
	}

	err = json.Unmarshal(respBody, &ret)
	if err != nil {
		return
	}
	if ret.Ret.Code != 0 {
		return "", errors.New(ret.Ret.Msg)
	}

	return ret.Data.AccessToken, nil
}

/**
获取 sdk_token，用于客户端登录鉴权。
*/
func (cli ZegoClient) GetSDKToken(deviceId string, platform int) (sdkToken string, err error) {
	// 过期时间
	expired := time.Now().Unix() + 3600

	// 计算hash
	m := md5.New()
	m.Write([]byte(fmt.Sprintf("%s%s%d%d%d", cli.SecretSign[:32], deviceId, 3, 1, expired)))
	hash := fmt.Sprintf("%x", m.Sum(nil))

	var reqStruct struct {
		CommonData struct {
			Platform int `json:"platform"`
		}
		SecretId  int    `json:"secret_id"`
		Sign      string `json:"sign"`
		DeviceId  string `json:"device_id"`
		Timestamp int64  `json:"timestamp"`
	}
	reqStruct.CommonData.Platform = platform
	reqStruct.SecretId = cli.SecretId
	reqStruct.Sign = hash
	reqStruct.DeviceId = deviceId
	reqStruct.Timestamp = expired
	reqBody, _ := json.Marshal(reqStruct)

	// 发送post请求
	req, err := http.NewRequest(http.MethodPost, "https://roomkit-api.zego.im/auth/get_sdk_token", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var ret struct {
		Ret struct {
			Code    int64    `json:"code"`
			Msg     string `json:"msg"`
			Version string `json:"version"`
		} `json:"ret"`
		Data struct {
			SDKToken string `json:"sdk_token"`
		}
	}

	err = json.Unmarshal(respBody, &ret)
	if err != nil {
		return
	}
	if ret.Ret.Code != 0 {
		return "", errors.New(ret.Ret.Msg)
	}

	return ret.Data.SDKToken, nil
}

func requestPost(url string, param interface{}, response interface{}) error {
	reqBody, _ := json.Marshal(param)
	// 发送post请求
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var respBody []byte
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var ret struct {
		Ret struct {
			Code    int    `json:"code"`
			Msg     string `json:"msg"`
			Version string `json:"version"`
		} `json:"ret"`
		Data interface{}
	}

	err = json.Unmarshal(respBody, &ret)
	if err != nil {
		return err
	}
	if ret.Ret.Code != 0 {
		return errors.New(ret.Ret.Msg)
	}

	jsonBody, err := json.Marshal(&ret.Data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonBody, &response)
	if err != nil {
		return err
	}

	return nil
}