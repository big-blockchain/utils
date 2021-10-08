package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5String(str string) string {
	h := md5.New()
	h.Write(StringToBytes(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5(str []byte) []byte {
	h := md5.New()
	h.Write(str)
	return h.Sum(nil)
}

func Md5Byte(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}
