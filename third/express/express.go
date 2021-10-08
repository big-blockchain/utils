/**
 * @Auth: Nuts
 * @Date: 2021/8/13 11:31 上午
 */
package express

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.onedao.com.cn/onedao/backend/backend/utils-go/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

type Express struct {
	Key      string
	Customer string
	Salt     string
}

func New(key, customer, salt string) *Express {
	return &Express{
		Key:      key,
		Customer: customer,
		Salt:     salt,
	}
}

type ListExpress struct {
	Message    string
	Result     bool
	ReturnCode string
	State      string
	Status     string
	Com        string //物流公司
	Nu         string //运单号
	Data       []ExpressData
}

type ExpressData struct {
	Context string
	Time    string
	Ftime   string
}

//查询物流信息
//com:查询的快递公司的编码
//num:快递单号
func (receiver *Express) QueryList(com, number string) (error, ListExpress) {
	url := "https://poll.kuaidi100.com/poll/query.do"
	param := utils.M{
		"com":      com,
		"num":      number,
		"phone":    "",
		"from":     "",
		"to":       "",
		"resultv2": 3,
		"show":     "0",
		"order":    "desc",
	}
	sign := strings.ToUpper(utils.Md5String(fmt.Sprintf("%s%s%s", utils.JsonEncode(param), receiver.Key, receiver.Customer)))
	bodyParamStr := fmt.Sprintf("customer=%s&param=%s&sign=%s", receiver.Customer, utils.JsonEncode(param), sign)

	list := ListExpress{}
	con := &http.Client{}
	body := ioutil.NopCloser(bytes.NewReader([]byte(bodyParamStr)))
	req1, _ := http.NewRequest(http.MethodPost, url, body)

	req1.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := con.Do(req1)

	if err != nil {
		return err, list
	}

	defer resp.Body.Close()

	bodys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, list
	}

	err = json.Unmarshal(bodys, &list)
	if err != nil {
		return err, list
	}

	return nil, list
}

type SubscribeResult struct {
	Result     bool   `json:"result"`
	ReturnCode string `json:"returnCode"`
	Message    string `json:"message"`
}

//订阅物流信息
func (receiver *Express) Subscribe(com, number, callback string) (error, SubscribeResult) {
	url := "https://poll.kuaidi100.com/poll"
	param := utils.M{
		"company": com,
		"number":  number,
		"from":    "",
		"to":      "",
		"key":     receiver.Key,
		"parameters": utils.M{
			"callbackurl":        callback,
			"salt":               receiver.Salt,
			"resultv2":           "1",
			"autoCom":            "0",
			"interCom":           "0",
			"departureCountry":   "",
			"departureCom":       "",
			"destinationCountry": "",
			"destinationCom":     "",
			"phone":              "",
		},
	}
	sign := strings.ToUpper(utils.Md5String(fmt.Sprintf("%s%s%s", utils.JsonEncode(param), receiver.Key, receiver.Customer)))
	bodyParamStr := fmt.Sprintf("customer=%s&param=%s&sign=%s", receiver.Customer, utils.JsonEncode(param), sign)

	res := SubscribeResult{}
	con := &http.Client{}
	body := ioutil.NopCloser(bytes.NewReader([]byte(bodyParamStr)))
	req1, _ := http.NewRequest(http.MethodPost, url, body)

	req1.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := con.Do(req1)

	if err != nil {
		return err, res
	}

	defer resp.Body.Close()

	bodys, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, res
	}

	err = json.Unmarshal(bodys, &res)
	if err != nil {
		return err, res
	}

	return nil, res
}

type SubscribeData struct {
	Sign  string    `json:"sign"`
	Param ParamData `json:"param"`
}
type ParamData struct {
	Status     string     `json:"status"`
	Billstatus string     `json:"billstatus"`
	Message    string     `json:"message"`
	AutoCheck  string     `json:"autoCheck"`
	ComOld     string     `json:"comOld"`
	ComNew     string     `json:"comNew"`
	LastResult LastResult `json:"lastLesult"`
}
type LastResult struct {
	Message   string           `json:"message"`
	State     string           `json:"state"`
	Status    string           `json:"status"`
	Condition string           `json:"condition"`
	Ischeck   string           `json:"ischeck"`
	Com       string           `json:"com"`
	Nu        string           `json:"nu"`
	Data      []LastResultData `json:"data"`
}
type LastResultData struct {
	Context  string `json:"context"`
	Time     string `json:"time"`
	Ftime    string `json:"ftime"`
	Status   string `json:"status"`
	AreaCode string `json:"areaCode"`
	AreaName string `json:"areaName"`
}

func (receiver Express) SubscribeCallback(ctx *gin.Context) (bool, LastResult) {
	lastResult := LastResult{}
	subscribeData := SubscribeData{}
	if err := ctx.ShouldBindJSON(&subscribeData); err != nil {
		return false, lastResult
	}

	sign := subscribeData.Sign
	paramStr := ctx.PostForm("param")
	signStr := utils.Md5String(paramStr + receiver.Salt)
	if sign == signStr {
		return false, lastResult
	}

	lastResult = subscribeData.Param.LastResult
	return true, lastResult
}
