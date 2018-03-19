package util

import (
	"time"
)

type DayRange struct {
	Begin time.Time
	End   time.Time
}

func dayDur(now time.Time) (time.Time, time.Time) {

	today_str := now.Format("20060102")

	yesterday_Date := now.AddDate(0, 0, -1)
	yesterday_str := yesterday_Date.Format("20060102")

	yesterday_time, _ := time.Parse("20060102", yesterday_str)
	today_time, _ := time.Parse("20060102", today_str)

	return yesterday_time, today_time
}

func YesterdayDur() (time.Time, time.Time) {
	now := time.Now()
	return dayDur(now)
}
func WeekDur() []*DayRange {
	var week = []*DayRange{}
	now := time.Now()
	for i := 7; i > 0; i-- {
		begin, end := dayDur(now)
		day := &DayRange{begin, end}
		week = append(week, day)
		now = begin
	}
	return week
}
