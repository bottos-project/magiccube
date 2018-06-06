/*Copyright 2017~2022 The Bottos Authors
  This file is part of the Bottos Data Exchange Client
  Created by Developers Team of Bottos.

  This program is free software: you can distribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with Bottos. If not, see <http://www.gnu.org/licenses/>.
*/

package util

import (
	"time"
)

//DayRange struct
type DayRange struct {
	Begin time.Time
	End   time.Time
}

//dayDur is to dayDur

func dayDur(now time.Time) (time.Time, time.Time) {

	todayStr := now.Format("20060102")

	yesterdayDate := now.AddDate(0, 0, -1)
	yesterdayStr := yesterdayDate.Format("20060102")

	yesterdayTime, _ := time.Parse("20060102", yesterdayStr)
	todayTime, _ := time.Parse("20060102", todayStr)

	return yesterdayTime, todayTime
}

//YesterdayDur is to YesterdayDur
func YesterdayDur() (time.Time, time.Time) {
	now := time.Now()
	return dayDur(now)
}

//WeekDur is to WeekDur
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
