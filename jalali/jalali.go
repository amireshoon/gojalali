package jalali

import (
	"fmt"
	"strings"
	"time"
)

type JalaliTime struct {
	Year  int
	Month int
	Day   int
	time  time.Time
}

func NewJalaliTime(year, month, day int) JalaliTime {
	t := ToGregorian(year, month, day)
	return JalaliTime{
		Year:  year,
		Month: month,
		Day:   day,
		time:  t,
	}
}

func Now() JalaliTime {
	t := time.Now()
	jy, jm, jd := ToJalali(t)
	return JalaliTime{
		Year:  jy,
		Month: jm,
		Day:   jd,
		time:  t,
	}
}

func (jt JalaliTime) ToTime() time.Time {
	return jt.time
}

func (jt JalaliTime) Format(layout string) string {
	return ToJalaliFormat(jt.time, layout)
}

func isLeapJalali(jy int) bool {
	leapYears := []int{1, 5, 9, 13, 17, 22, 26, 30}
	mod := jy % 33
	for _, y := range leapYears {
		if mod == y {
			return true
		}
	}
	return false
}

var gDaysInMonth = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var jDaysInMonth = [12]int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}
var persianMonthNames = [...]string{
	"", "فروردین", "اردیبهشت", "خرداد", "تیر", "مرداد", "شهریور",
	"مهر", "آبان", "آذر", "دی", "بهمن", "اسفند",
}
var persianWeekdayNames = [...]string{
	"یک‌شنبه", "دوشنبه", "سه‌شنبه", "چهارشنبه", "پنج‌شنبه", "جمعه", "شنبه",
}

// Convert Gregorian to Jalali
func ToJalali(t time.Time) (jy, jm, jd int) {
	gy := t.Year()
	gm := int(t.Month())
	gd := t.Day()

	gy2 := gy - 1600
	gm2 := gm - 1
	gd2 := gd - 1

	gDayNo := 365*gy2 + (gy2+3)/4 - (gy2+99)/100 + (gy2+399)/400
	for i := 0; i < gm2; i++ {
		gDayNo += gDaysInMonth[i]
	}
	if gm2 > 1 && isLeapGregorian(gy) {
		gDayNo++
	}
	gDayNo += gd2

	jDayNo := gDayNo - 79

	jNp := jDayNo / 12053
	jDayNo %= 12053

	jy = 979 + 33*jNp + 4*(jDayNo/1461)
	jDayNo %= 1461

	if jDayNo >= 366 {
		jy += (jDayNo - 1) / 365
		jDayNo = (jDayNo - 1) % 365
	}

	var i int
	for i = 0; i < 11 && jDayNo >= jDaysInMonth[i]; i++ {
		jDayNo -= jDaysInMonth[i]
	}
	jm = i + 1
	jd = jDayNo + 1

	return
}

func ToJalaliTime(t time.Time) JalaliTime {
	y, m, d := ToJalali(t)
	return JalaliTime{
		Year:  y,
		Month: m,
		Day:   d,
		time:  t,
	}
}

// ToJalaliFormat formats a time.Time to a Jalali date string using a custom format.
func ToJalaliFormat(t time.Time, format string) string {
	jy, jm, jd := ToJalali(t)

	replacer := map[string]string{
		"yyyy": fmt.Sprintf("%04d", jy),
		"yy":   fmt.Sprintf("%02d", jy%100),
		"mm":   fmt.Sprintf("%02d", jm),
		"m":    fmt.Sprintf("%d", jm),
		"dd":   fmt.Sprintf("%02d", jd),
		"d":    fmt.Sprintf("%d", jd),
	}

	// Replace longer tokens first
	tokens := []string{"yyyy", "yy", "mm", "m", "dd", "d"}
	result := format
	for _, token := range tokens {
		result = strings.ReplaceAll(result, token, replacer[token])
	}

	return result
}

// Convert Jalali to Gregorian
func ToGregorian(jy, jm, jd int) time.Time {
	jy -= 979
	jm--
	jd--

	jDayNo := 365*jy + jy/33*8 + (jy%33+3)/4
	for i := 0; i < jm; i++ {
		jDayNo += jDaysInMonth[i]
	}
	jDayNo += jd

	gDayNo := jDayNo + 79

	gy := 1600 + 400*(gDayNo/146097)
	gDayNo %= 146097

	leap := true
	if gDayNo >= 36525 {
		gDayNo--
		gy += 100 * (gDayNo / 36524)
		gDayNo %= 36524

		if gDayNo >= 365 {
			gDayNo++
		} else {
			leap = false
		}
	}

	gy += 4 * (gDayNo / 1461)
	gDayNo %= 1461

	if gDayNo >= 366 {
		leap = false
		gDayNo--
		gy += gDayNo / 365
		gDayNo %= 365
	}

	var i int
	for i = 0; gDayNo >= gDaysInMonth[i]+boolToInt(i == 1 && leap); i++ {
		gDayNo -= gDaysInMonth[i]
		if i == 1 && leap {
			gDayNo--
		}
	}
	gm := i + 1
	gd := int(gDayNo) + 1

	return time.Date(gy, time.Month(gm), gd, 0, 0, 0, 0, time.UTC)
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func isLeapGregorian(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func daysInJalaliMonth(year, month int) int {
	if month <= 6 {
		return 31
	} else if month <= 11 {
		return 30
	}
	if isLeapJalali(year) {
		return 30
	}
	return 29
}

func (jt JalaliTime) MonthName() string {
	if jt.Month >= 1 && jt.Month <= 12 {
		return persianMonthNames[jt.Month]
	}
	return ""
}

func (jt JalaliTime) WeekdayName() string {
	return persianWeekdayNames[jt.time.Weekday()]
}

func (jt JalaliTime) AddDays(n int) JalaliTime {
	return ToJalaliTime(jt.time.AddDate(0, 0, n))
}

func (jt JalaliTime) AddWeeks(n int) JalaliTime {
	return ToJalaliTime(jt.time.AddDate(0, 0, n*7))
}

func (jt JalaliTime) AddMonths(n int) JalaliTime {
	y := jt.Year
	m := jt.Month
	d := jt.Day

	// Shift months
	m += n
	for m > 12 {
		m -= 12
		y++
	}
	for m < 1 {
		m += 12
		y--
	}

	// Clamp day if needed
	maxDay := daysInJalaliMonth(y, m)
	if d > maxDay {
		d = maxDay
	}

	return NewJalaliTime(y, m, d)
}

func (jt JalaliTime) AddYears(n int) JalaliTime {
	y := jt.Year + n
	m := jt.Month
	d := jt.Day

	maxDay := daysInJalaliMonth(y, m)
	if d > maxDay {
		d = maxDay
	}

	return NewJalaliTime(y, m, d)
}

func (jt JalaliTime) RemoveDays(n int) JalaliTime {
	return jt.AddDays(-n)
}

func (jt JalaliTime) RemoveWeeks(n int) JalaliTime {
	return jt.AddWeeks(-n)
}

func (jt JalaliTime) RemoveMonths(n int) JalaliTime {
	return jt.AddMonths(-n)
}

func (jt JalaliTime) RemoveYears(n int) JalaliTime {
	y := jt.Year - n
	m := jt.Month
	d := jt.Day

	maxDay := daysInJalaliMonth(y, m)
	if d > maxDay {
		d = maxDay
	}

	return NewJalaliTime(y, m, d)
}
