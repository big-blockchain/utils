package utils

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
)

func HmacSha1(dataStr string, keyStr string) []byte {
	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(dataStr))
	return mac.Sum(nil)
}

func HmacSha1Hex(dataStr string, keyStr string) string {
	row := HmacSha1(dataStr, keyStr)
	return hex.EncodeToString(row)
}
func HmacSha1Base64(dataStr string, keyStr string) string {
	row := HmacSha1(dataStr, keyStr)
	return base64.StdEncoding.EncodeToString(row)
}
