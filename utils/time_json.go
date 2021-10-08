/**
 * @Auth: Nuts
 * @Date: 2021/3/31 10:25 上午
 */
package utils

import (
	"time"
)

//格式化时间
type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(j).Format("2006-01-02 15:04:05") + `"`), nil
}

func (t *JsonTime) UnmarshalJSON(data []byte) (err error) {
	// Fractional seconds are handled implicitly by Parse.
	ts, err := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*t = JsonTime(ts)
	return
}

func (t JsonTime) Parse(datetime string) (JsonTime, error) {
	ts, err := time.ParseInLocation("2006-01-02 15:04:05", string(datetime), time.Local)
	return JsonTime(ts), err
}

func (t JsonTime) Time() time.Time {
	return time.Time(t)
}

func (t JsonTime) Now() JsonTime {
	return JsonTime(time.Now())
}

type JsonDate time.Time

func (d JsonDate) MarshalJSON() ([]byte, error) {
	return []byte(`"` + time.Time(d).Format("2006-01-02") + `"`), nil
}

func (d *JsonDate) UnmarshalJSON(data []byte) (err error) {
	// Fractional seconds are handled implicitly by Parse.
	ts, err := time.ParseInLocation(`"2006-01-02"`, string(data), time.Local)
	*d = JsonDate(ts)
	return
}

func (d JsonDate) Parse(date string) (JsonDate, error) {
	ts, err := time.ParseInLocation("2006-01-02", string(date), time.Local)
	return JsonDate(ts), err
}
