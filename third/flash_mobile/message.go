/**
 * @Auth: Nuts
 * @Date: 2021/5/24 11:30 上午
 */
package flash_mobile

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

//发送短信
func SendSmsMessage(account, password, urlStr, msg, mobile string) (string, error) {
	params := make(map[string]interface{})
	//请登录zz.253.com获取API账号、密码以及短信发送的URL
	params["account"] = account   //创蓝API账号
	params["password"] = password //创蓝API密码
	params["phone"] = mobile      //手机号码
	//设置您要发送的内容：其中“【】”中括号为运营商签名符号，多签名内容前置添加提交
	params["msg"] = url.QueryEscape(msg)
	params["report"] = "true"
	bytesData, err := json.Marshal(params)

	if err != nil {
		return "", err
	}

	reader := bytes.NewReader(bytesData)
	request, err := http.NewRequest("POST", urlStr, reader)

	if err != nil {
		return "", err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	bodys, err := ioutil.ReadAll(resp.Body)
	respBody := string(bodys)
	var mapResult struct {
		Code     string `json:"code"`
		ErrorMsg string `json:"errorMsg"`
	}
	if err2 := json.Unmarshal([]byte(respBody), &mapResult); err2 != nil {
		return "", err2
	}
	if mapResult.Code == "0" {
		return mapResult.Code, nil
	} else {
		return mapResult.Code, errors.New(mapResult.ErrorMsg)
	}
}
