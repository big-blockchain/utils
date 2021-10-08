package utils

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
)

func Sha1(str string) string {
	h := sha1.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Sha256(str string) string {
	h := sha256.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
