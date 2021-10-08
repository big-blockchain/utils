package utils

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

//CBC解密
func DecryptDES_CBC(src, key string) (string, error) {
	keyByte := []byte(key)
	data, err := base64.StdEncoding.DecodeString(src)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	iv := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(data))
	mode.CryptBlocks(plaintext, data)
	plaintext = PKCS5UnPadding(plaintext)
	return string(plaintext), nil
}

//明文减码算法
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	if length == 0 {
		return []byte{}
	}
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
