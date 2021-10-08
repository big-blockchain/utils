/**
 * @Auth: Nuts
 * @Date: 2021/5/24 10:54 上午
 */
package flash_mobile

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//闪验V2版本
func GetMobileNo(appId, appKey, token string) (string, error) {
	u := "http://api.253.com/open/flashsdk/mobile-query"
	formData := url.Values{}
	formData.Set("appId", appId)
	formData.Set("token", token)
	formData.Set("sign", signParam(formData, appKey))
	resp, err := http.PostForm(u, formData)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	respBody := string(body)
	var mapResult map[string]interface{}
	if err2 := json.Unmarshal([]byte(respBody), &mapResult); err2 != nil {
		return "", err2
	}
	fmt.Println("mapResult", mapResult)
	if mapResult["code"] == "200000" {
		data := (mapResult["data"]).(map[string]interface{})
		mobileName := data["mobileName"].(string)
		return decryptPhone(mobileName, appKey), nil
	} else {
		return "", errors.New("置换手机号失败")
	}
}

func PKCS5Unpadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func signParam(data url.Values, key string) string {
	message := "appId" + data.Get("appId") + "token" + data.Get("token")
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(message))
	signature := hex.EncodeToString(mac.Sum(nil))
	return string(signature)
}

func decryptPhone(data string, key string) string {
	hash := md5.Sum([]byte(key))
	hashString := hex.EncodeToString(hash[:])
	block, _ := aes.NewCipher([]byte(hashString[:16]))
	ecb := cipher.NewCBCDecrypter(block, []byte(hashString[16:]))
	source, _ := hex.DecodeString(data)
	decrypted := make([]byte, len(source))
	ecb.CryptBlocks(decrypted, source)
	return string(PKCS5Unpadding(decrypted))
}
