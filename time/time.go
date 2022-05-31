// Copyright 29-May-2017 ºDeme
// GNU General Public License - V3 <http://www.gnu.org/licenses/>

// Functions to manage time-dates.
//
// Time is stored as an int64 corresponding to the Unix time (milliseconds
// since January 1, 1970).
//
// Therefore it is possible to use integer functions like '+', '-', '>', '==',
// etc.
package time

import (
	"strconv"
	"strings"
	gtime "time"
)

func allDigits(s string) bool {
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

func toInt(s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}

// Add 'days' to 'tm'. 'days' can be negative.
func AddDays(tm int64, days int) int64 {
	return gtime.UnixMilli(tm).AddDate(0, 0, days).UnixMilli()
}

// Returns the day of 'tm', between 1 and 31, both inclusive.
func Day(tm int64) int {
	return gtime.UnixMilli(tm).Day()
}

// Returns t1 - t2 in days,
func DfDays(t1, t2 int64) int {
	t1m := t1 / int64(86400000)
	if t1 < int64(0) {
		t1m--
	}
	t2m := t2 / int64(86400000)
	if t2 < int64(0) {
		t2m--
	}
	return int(t1m - t2m)
}

// Returns 'true' if 'time.dfDays(tm1, tm2) == 0'.
func EqDay(t1, t2 int64) bool {
	return DfDays(t1, t2) == 0
}

// Returns a string representation of 'tm'.
//    'template' uses the folowing variables to do substitution:
//      %d  Day in number 06 -> 6.
//      %D  Day with tow digits 06 -> 06.
//      %m  Month in number 03 -> 3.
//      %M  Month with two digits 03 -> 03.
//      %y  Year with two digits 2010 -> 10.
//      %Y  Year with four digits 2010 -> 2010.
//      %t  Time without milliseconds -> 15:03:55
//      %T  Time with milliseconds -> 15:03:55.345
//      %%  The sign '%'.
func Fmt(template string, tm int64) string {
	d := gtime.UnixMilli(tm)
	tmp := strings.ReplaceAll(template, "%d", gtime.Time(d).Format("2"))
	tmp = strings.ReplaceAll(tmp, "%D", gtime.Time(d).Format("02"))
	tmp = strings.ReplaceAll(tmp, "%m", gtime.Time(d).Format("1"))
	tmp = strings.ReplaceAll(tmp, "%M", gtime.Time(d).Format("01"))
	tmp = strings.ReplaceAll(tmp, "%y", gtime.Time(d).Format("06"))
	tmp = strings.ReplaceAll(tmp, "%Y", gtime.Time(d).Format("2006"))
	tmp = strings.ReplaceAll(tmp, "%t", gtime.Time(d).Format("15:04:05"))
	tmp = strings.ReplaceAll(tmp, "%T", gtime.Time(d).Format("15:04:05.000"))
	tmp = strings.ReplaceAll(tmp, "%%", "%")
	return tmp
}

// Returns a new time equals to 'tm', but setting hour, minute and second
// matching a string type: "HH:MM:SS", and milliseconds to 0.
func FromClock(tm int64, s string) int64 {
	ps := strings.Split(s, ":")
	if len(ps) != 3 || !allDigits(ps[0]) ||
		!allDigits(ps[1]) || !allDigits(ps[2]) ||
		len(ps[0]) != 2 || len(ps[1]) != 2 || len(ps[2]) != 2 ||
		toInt(ps[0]) < 0 || toInt(ps[0]) > 23 ||
		toInt(ps[1]) < 0 || toInt(ps[1]) > 59 ||
		toInt(ps[2]) < 0 || toInt(ps[1]) > 59 {
		panic("'" + s + "' bad clock.")
	} else {
		t := gtime.UnixMilli(tm)
		lc, err := gtime.LoadLocation("Local")
		if err != nil {
			panic(err)
		}
		return (gtime.Date(t.Year(), t.Month(), t.Day(),
			toInt(ps[0]), toInt(ps[1]), toInt(ps[2]),
			0, lc).UnixMilli())
	}
}

// Returns a time from a string type: MM*DD*YYYY, where '*' is the separator 'sep'.
func FromEn(s, sep string) int64 {
	ps := strings.Split(s, sep)
	if len(ps) != 3 || !allDigits(ps[0]) ||
		!allDigits(ps[1]) || !allDigits(ps[2]) {
		panic("'" + s + "' bad english date.")
	} else {
		lc, err := gtime.LoadLocation("Local")
		if err != nil {
			panic(err)
		}
		return gtime.Date(toInt(ps[2]), gtime.Month(toInt(ps[0])),
			toInt(ps[1]), 12, 0, 0, 0, lc).UnixMilli()
	}
}

// Returns a time from a string type: DD*MM*YYYY, where '*' is the separator 'sep'.
func FromIso(s, sep string) int64 {
	ps := strings.Split(s, sep)
	if len(ps) != 3 || !allDigits(ps[0]) ||
		!allDigits(ps[1]) || !allDigits(ps[2]) {
		panic("'" + s + "' bad ISO date.")
	} else {
		lc, err := gtime.LoadLocation("Local")
		if err != nil {
			panic(err)
		}
		return gtime.Date(toInt(ps[2]), gtime.Month(toInt(ps[1])),
			toInt(ps[0]), 12, 0, 0, 0, lc).UnixMilli()
	}
}

// Returns a time from a string type: YYYYMMDD.
func FromStr(s string) int64 {
	if len(s) != 8 || !allDigits(s) {
		panic("'" + s + "' bad date.")
	} else {
		lc, err := gtime.LoadLocation("Local")
		if err != nil {
			panic(err)
		}
		return gtime.Date(toInt(s[:4]), gtime.Month(toInt(s[4:6])),
			toInt(s[6:]), 12, 0, 0, 0, lc).UnixMilli()
	}
}

// Returns the hour of 'tm', between 0 and 23, both inclusive.
func Hour(tm int64) int {
	return gtime.UnixMilli(tm).Hour()
}

// Returns the minute of 'tm', between 0 and 59, both inclusive.
func Minute(tm int64) int {
	return gtime.UnixMilli(tm).Minute()
}

// Returns the month of 'tm', between 1 and 12, both inclusive.
func Month(tm int64) int {
	return int(gtime.UnixMilli(tm).Month())
}

// Returns a new time
//    day: Between 1 and 31, both inclusive.
//    month: Between 1 and 12, both inclusive.
//    year: With 4 digits.
//    hour: Between 0 and 23, both inclusive.
//    minute: Between 0 and 59, both inclusive.
//    second: Between 0 and 59, both inclusive.
func New(day, month, year, hour, minute, second int) int64 {
	lc, err := gtime.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	return (gtime.Date(year, gtime.Month(int(month)), day,
		hour, minute, second, 0, lc).UnixMilli())
}

// Returns a new time with hour wt to 12 and  minute, seconds and millisecond to 0.
//    day: Between 1 and 31, both inclusive.
//    month: Between 1 and 12, both inclusive.
//    year: With 4 digits.
func NewDate(day, month, year int) int64 {
	lc, err := gtime.LoadLocation("Local")
	if err != nil {
		panic(err)
	}
	return (gtime.Date(year, gtime.Month(int(month)), day, 12, 0, 0, 0, lc).UnixMilli())
}

// Returns the current time.
func Now() int64 {
	return gtime.Now().UnixMilli()
}

// Returns the second of 'tm', between 0 and 59, both inclusive.
func Second(tm int64) int {
	return gtime.UnixMilli(tm).Second()
}

// Equals to 'time.Fmt("%M-%D-%Y", tm)'. (e.g.; "02-30-2022")
func ToEn(tm int64) string {
	return Fmt("%M-%D-%Y", tm)
}

// Equals to 'time.Fmt("%D/%M/%Y", tm)'. (e.g.; "30/01/2022")
func ToIso(tm int64) string {
	return Fmt("%D/%M/%Y", tm)
}

// Equals to 'time.Fmt("%Y%M%D", tm)'. (e.g.; "20220130")
func ToStr(tm int64) string {
	return Fmt("%Y%M%D", tm)
}

// Returns the weekday of 'tm', between 0 (Sunday) and 6 (Saturday).
func Weekday(tm int64) int {
	return int(gtime.UnixMilli(tm).Weekday())
}

// Returns the year of 'tm'.
func Year(tm int64) int {
	return gtime.UnixMilli(tm).Year()
}

// Returns the year day of 'tm'.
func YearDay(tm int64) int {
	return gtime.UnixMilli(tm).YearDay()
}
