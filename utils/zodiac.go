package utils

import "strconv"

//根据出生年份获取生肖
func GetZodiac(birthday string) (zodiac string) {
	var year int64
	if len(birthday) <= 4 {
		year, _ = strconv.ParseInt(birthday, 10, 64)
	} else {
		yearStr := SubString(birthday, 0, 4)
		year, _ = strconv.ParseInt(yearStr, 10, 64)
	}

	if year <= 0 {
		zodiac = ""
	}
	var start int64
	start = 1901
	x := (start - year) % 12
	if x == 1 || x == -11 {
		zodiac = "鼠"
	}
	if x == 0 {
		zodiac = "牛"
	}
	if x == 11 || x == -1 {
		zodiac = "虎"
	}
	if x == 10 || x == -2 {
		zodiac = "兔"
	}
	if x == 9 || x == -3 {
		zodiac = "龙"
	}
	if x == 8 || x == -4 {
		zodiac = "蛇"
	}
	if x == 7 || x == -5 {
		zodiac = "马"
	}
	if x == 6 || x == -6 {
		zodiac = "羊"
	}
	if x == 5 || x == -7 {
		zodiac = "猴"
	}
	if x == 4 || x == -8 {
		zodiac = "鸡"
	}
	if x == 3 || x == -9 {
		zodiac = "狗"
	}
	if x == 2 || x == -10 {
		zodiac = "猪"
	}
	return
}
