package main

import (
	"fmt"
	"time"

	"github.com/amireshoon/gojalali/jalali"
)

func main() {
	t := time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC)
	jt := jalali.ToJalaliTime(t)

	weekdayName := jt.WeekdayName()
	fmt.Printf("Gregorian date: %s\n", t.Format("2006-01-02 Monday"))
	fmt.Printf("Jalali date: %d/%02d/%02d %s\n", jt.Year, jt.Month, jt.Day, weekdayName)
}
