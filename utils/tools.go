package utils

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
	"math"
	r "math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unsafe"
)

//根据slice结构体获取某个字段的的int64 slice
func GetSliceColsInt64(rows interface{}, col ...string) []int64 {
	sliceValue := reflect.Indirect(reflect.ValueOf(rows))
	var cols []int64
	if len(col) == 0 {
		return cols
	}
	if sliceValue.Kind() != reflect.Slice && sliceValue.Kind() != reflect.Map {
		return cols
	}
	for i := 0; i < sliceValue.Len(); i++ {
		for _, n := range col {
			cols = append(cols, sliceValue.Index(i).FieldByName(n).Int())
		}
	}
	return cols
}

//根据slice结构体获取某个字段的的int64 slice
func GetSliceColsString(rows interface{}, col ...string) []string {
	sliceValue := reflect.Indirect(reflect.ValueOf(rows))
	var cols []string
	if len(col) == 0 {
		return cols
	}
	if sliceValue.Kind() != reflect.Slice && sliceValue.Kind() != reflect.Map {
		return cols
	}
	for i := 0; i < sliceValue.Len(); i++ {
		for _, n := range col {
			cols = append(cols, sliceValue.Index(i).FieldByName(n).String())
		}
	}
	return cols
}

//查询两个数组中交集的个数
func GetIntersectionArray(items1, items2 []string) ([]string, int) {
	items := make([]string, 0)
	for _, i := range items1 {
		if InArray(items2, i) {
			items = append(items, i)
		}
	}

	return items, len(items)
}

// 判断slice是否包含某个item
func InArray(items interface{}, item interface{}) bool {
	value := reflect.ValueOf(items)
	size := value.Len()
	for i := 0; i < size; i++ {
		if value.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func ParamsAppend(Org map[string]interface{}, params interface{}) map[string]interface{} {
	paramStr := JsonUtils{}.JsonEncode(params)
	paramMap := M{}
	JsonUtils{}.JsonDecode(paramStr, &paramMap)

	for key, val := range paramMap {
		Org[key] = val
	}

	return Org
}

// 驼峰命名转小写，中间的大写转_
func CamelCaseLower(name string) string {
	reg := regexp.MustCompile(`([A-Z])`)
	str := strings.Trim(reg.ReplaceAllStringFunc(name, func(s string) string {
		return "_" + strings.ToLower(s)
	}), "_")
	return strings.Trim(str, "_")
}

// 驼峰转大写
func CamelCaseUpper(name string) string {
	split := strings.Split(name, "_")
	var strs []string
	for _, n := range split {
		strArry := []rune(n)
		if len(strArry) == 0 {
			continue
		}
		if strArry[0] >= 97 && strArry[0] <= 122 {
			strArry[0] -= 32
		}
		strs = append(strs, string(strArry))
	}

	return strings.Join(strs, "")
}

// 隐藏字符中的部分，str 原始字符串，rep替换为，start 从哪个开始，l 长度
func HideString(str string, start, l int, rep rune) string {
	var i, j int
	s := strings.Map(func(r rune) rune {
		i++
		if i > start && j < l {
			j++
			return rep
		}
		return r
	}, str)

	return s
}

// RandString 生成随机字符串
func RandString(len int64) string {
	bytes := make([]byte, len)
	var i int64
	for i = 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

// 根据经纬度计算距离
func GetDistance(lat1, lng1, lat2, lng2 float64) float64 {

	if (lat1 == 0 && lng1 == 0) || (lat2 == 0 && lng2 == 0) {
		return 0
	}

	radius := 6378.138 // 6378137
	rad := math.Pi / 180.0

	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	dist := math.Asin(math.Sqrt(math.Pow(math.Sin((lat1-lat2)/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin((lng1-lng2)/2), 2)))

	return dist * radius * 2
}

// 判断是否为直辖市，且返回直辖市
func Municipality(city string) (string, bool) {
	cities := []string{"北京", "上海", "天津", "重庆"}

	for _, n := range cities {
		ok := strings.Contains(city, n)
		if ok {
			return n, ok
		}
	}

	return "", false
}

//根据省市返回具体城市
func GetAbodeCity(abode string) string {
	municipality, ok := Municipality(abode)

	if ok {
		return municipality
	} else {
		abode = strings.Replace(abode, "~", "-", 1)
		abodeS := strings.Split(abode, "-")

		if len(abodeS) > 0 && abodeS[len(abodeS)-1] != "" {
			if strings.Contains(abode, "全省") {
				return abodeS[0]
			}
			if strings.Contains(abode, "不限") {
				return abodeS[0]
			}
			return abodeS[len(abodeS)-1]
		}
	}
	return abode
}

func GetAbodeProvince(abode string) string {
	municipality, ok := Municipality(abode)

	if ok {
		return municipality
	} else {
		abode = strings.Replace(abode, "~", "-", 1)
		abodeS := strings.Split(abode, "-")

		if len(abodeS) > 0 {
			return abodeS[0]
		}
	}
	return abode
}

func ReverseArray(arr *[]interface{}) {
	var temp interface{}
	length := len(*arr)
	for i := 0; i < length/2; i++ {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[length-1-i]
		(*arr)[length-1-i] = temp
	}
}


// 将字符串转，字符串为数字拼接
func StrToSliceInt64(str string, sep string) []int64 {
	var rId []int64
	for _, n := range strings.Split(str, sep) {
		id, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			continue
		}
		rId = append(rId, id)
	}
	return rId
}

func StringToBytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func JsonEncode(m interface{}) string {
	res, _ := json.Marshal(m)
	return BytesToString(res)
}

type M map[string]interface{}

func (m M) Json() []byte {
	if m == nil {
		return []byte{}
	}
	res, _ := json.Marshal(m)
	return res
}

func (m M) JsonString() string {
	return BytesToString(m.Json())
}

//校验身份证有效性
func CheckIDCard(idCard string) bool {
	valid := validation.Validation{}
	reg := regexp.MustCompile(`^[1-9]\d{5}(18|19|20)\d{2}((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
	return valid.Match(idCard, reg, "identification_number").Ok
}

func GetPageLimit(limit, page int64) (int, int) {
	return int(limit), int((page - 1) * limit)
}

//截取字符串，兼容中文
func SubString(str string, begin, length int) string {
	rs := []rune(str)
	lth := len(rs)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length

	if end > lth {
		end = lth
	}
	return string(rs[begin:end])
}
