/**
 * @Auth: Nuts
 * @Date: 2021/3/10 3:57 下午
 */
package aliyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

type EmailClient struct {
	AccessKeyId     string
	AccessKeySecret string
	AccountName     string
}

func NewEmailClient(accessKeyId, accessKeySecret, accountName string) *EmailClient {
	return &EmailClient{
		AccessKeyId:     accessKeyId,
		AccessKeySecret: accessKeySecret,
		AccountName:     accountName,
	}
}

type Result struct {
	RequestId string
	EnvId     string
	Message   string
}

func (client EmailClient) SingleSendMail(email, subject, body string) Result {
	rand.Seed(time.Now().UnixNano())

	data := map[string]string{
		"SignatureMethod":  "HMAC-SHA1",
		"SignatureNonce":   fmt.Sprintf("%s", rand.Int63()),
		"AccessKeyId":      client.AccessKeyId,
		"SignatureVersion": "1.0",
		"Timestamp":        time.Now().UTC().Format("2006-01-02T15:04:05Z"),
		"Format":           "JSON",

		"Action":         "SingleSendMail",
		"Version":        "2015-11-23",
		"AccountName":    client.AccountName,
		"ReplyToAddress": "true",
		"AddressType":    "1",
		"ToAddress":      email,
		"Subject":        subject,
		"TextBody":       body,
		"RegionId":       "cn-hangzhou",
	}

	sortQueryString := SortedString(data)
	stringToSign := fmt.Sprintf("GET&%s&%s", UrlEncode("/"), UrlEncode(sortQueryString[1:]))
	url := fmt.Sprintf("http://dm.aliyuncs.com/?Signature=%s%s", client.Sign(stringToSign), sortQueryString)

	res := Result{}
	r, err := http.Get(url)
	if err != nil {
		res.Message = err.Error()
		return res
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res.Message = err.Error()
		return res
	}

	err = json.Unmarshal(b, &res)
	if err != nil {
		res.Message = err.Error()
		return res
	}

	return res
}

func UrlEncode(in string) string {
	r := strings.NewReplacer("+", "%20", "*", "%2A", "%7E", "~")
	return r.Replace(url.QueryEscape(in))
}

func (client EmailClient) Sign(stringToSign string) string {
	h := hmac.New(sha1.New, []byte(fmt.Sprintf("%s&", client.AccessKeySecret)))
	h.Write([]byte(stringToSign))
	return UrlEncode(base64.StdEncoding.EncodeToString(h.Sum(nil)))
}

func Keys(data map[string]string) []string {
	var keys []string
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func SortedString(data map[string]string) string {
	var sortQueryString string
	for _, v := range Keys(data) {
		sortQueryString = fmt.Sprintf("%s&%s=%s", sortQueryString, v, UrlEncode(data[v]))
	}
	return sortQueryString
}
