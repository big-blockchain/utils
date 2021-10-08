package utils

import (
	"testing"
	"time"
)

func TestConstellation(t *testing.T) {

	birthday := "2018-03-21"
	con, _ := Constellation(birthday)
	if con != "白羊座" {
		t.Error()
	}

	birthday = "2018-03-21 00:00:00"
	con, _ = Constellation(birthday)
	if con != "" {
		t.Error()
	}

	birthday = "2018-04-21"
	con, _ = Constellation(birthday)
	if con != "金牛座" {
		t.Error("金牛座")
	}

	birthday = "2018-02-20"
	con, _ = Constellation(birthday)
	if con != "双鱼座" {
		t.Error("双鱼座")
	}
	birthday = "2018-10-23"
	con, _ = Constellation(birthday)
	if con != "天秤座" {
		t.Error("天秤座")
	}
	birthday = "2000-01-01"

	for i := 0; i < 600; i++ {
		bir, _ := time.Parse("2006-01-02", birthday)
		bir = bir.AddDate(0, 0, i)
		con, _ = Constellation(bir.Format("2006-01-02"))
	}
}

func TestNowUnix(t *testing.T) {
	if NowUnix() != time.Now().Unix() {
		t.Error()
	}
}

func TestTodayUnix(t *testing.T) {

	i := TodayUnix()
	if i == 0 {
		t.Error()
	}

	today, err := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), time.Local)
	if err != nil {
		t.Error()
	}

	if i != today.Unix() {
		t.Error()
	}
}

func TestBirthdayToAge(t *testing.T) {
	birthday := "2000-01-01 11:00:00"
	age, err := BirthdayToAge(birthday)
	if age != 0 {
		t.Error()
	}

	if err == nil {
		t.Error()
	}

	birth := time.Now().AddDate(-20, 0, 0)
	birthday = birth.Format("2006-01-02")

	age, err = BirthdayToAge(birthday)

	if err != nil {
		t.Error()
	}

	if age != 21 {
		t.Error("age error", age)
	}

}

func BenchmarkConstellation(b *testing.B) {
	birthday := "2000-01-01"
	for i := 0; i < b.N; i++ {
		bir, _ := time.Parse("2006-01-02", birthday)
		bir = bir.AddDate(0, 0, i)

		Constellation(bir.Format("2006-01-02"))
	}
}
