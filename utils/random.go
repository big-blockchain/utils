package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//生成随机数
func Random(lens int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	max := "1" + strings.Repeat("0", lens)
	b, _ := strconv.Atoi(max)
	num := r.Intn(b)
	var s string
	s = strconv.Itoa(num)
	z := lens - len(strconv.Itoa(num))
	if z > 0 {
		for i := 0; i < z; i++ {
			s += strconv.Itoa(rand.Intn(9))
		}
	}
	return s
}

// 随机数
func RandomNum(min, max int) int64 {

	if min > max {
		m := min
		min = max
		max = m
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	num := r.Intn(max)

	if num < min {
		return RandomNum(min, max)
	}

	return int64(num)
}

func RandomStr(lens int) string {
	str := "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm1234567890"
	strLen := len(str)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	result := ""
	for i := 0; i < lens; i++ {
		i := r.Int()

		result = result + string(str[i%strLen])

	}
	return result
}
