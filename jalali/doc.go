/*
Package jalali provides utilities to work with the Jalali (Persian) calendar in Go.

It wraps Go's standard time.Time with Jalali calendar conversion, formatting, and
date arithmetic support including adding/removing days, weeks, months, and years.

Features:
- Convert Gregorian time.Time to Jalali date (year, month, day).
- Convert Jalali date back to Gregorian time.Time.
- Format Jalali dates with custom patterns (yyyy, mm, dd).
- Create JalaliTime structs that wrap time.Time with Jalali fields.
- Jalali date arithmetic (AddDays, AddMonths, RemoveYears, etc).
- Retrieve Jalali weekday names.

Example usage:

	jt := jalali.NewJalaliTime(1402, 12, 29)
	fmt.Println(jt.String()) // 1402/12/29

	now := time.Now()
	y, m, d := jalali.ToJalali(now)
	fmt.Printf("Today Jalali: %d/%02d/%02d\n", y, m, d)

	jt2 := jt.AddMonths(1)
	fmt.Println(jt2) // 1403/01/29
*/
package jalali
