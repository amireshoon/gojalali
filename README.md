# gojalali

A rich Go library for converting and interacting with Jalali (Persian) calendar dates.  
This package provides easy-to-use functions to convert between Gregorian and Jalali calendars, format Jalali dates, and perform date arithmetic in the Jalali calendar system.

---

## Features

- Convert Gregorian `time.Time` to Jalali year/month/day and vice versa  
- Create Jalali dates and convert them to `time.Time`  
- Format Jalali dates using customizable formats  
- Date arithmetic functions (add/remove days, weeks, months, years) in Jalali calendar  
- Retrieve Jalali weekday names and month names  
- Simple and idiomatic Go API  

---

## Installation

Use Go Modules to install the package:

```bash
go get github.com/yourusername/gojalali/jalali
```

## Quick Start
Basic Conversion (Gregorian to Jalali)
```go
t := time.Date(2025, 5, 20, 0, 0, 0, 0, time.UTC)
y, m, d := jalali.ToJalali(t)
fmt.Printf("Jalali date: %d/%02d/%02d\n", y, m, d)
```
Or you can create date from Jalali calendar
```go
jt := jalali.NewJalaliTime(1404, 2, 30)
fmt.Printf("Jalali Date: %d/%02d/%02d\n", jt.Year, jt.Month, jt.Day)
```
Then you can even convert it to Gregorian calendar
```go
fmt.Println("Gregorian date:", jt.ToGregorian())
```
Also you can add `days`,`weeks`,`months` and `years` to your JalaliTime
```go
jt.AddDays(1)
jt.AddWeeks(1)
jt.AddMonths(1)
jt.AddYears(1)
```
And you can format it however you want it
```go
jalali.ToJalaliFormat(time.Now(), "yyyy/mm/dd")
```

## Examples
See [Examples](examples/) directory for runnable example programs demonstrating various use cases.

## Documentaion
Full API documentation is available at [pkg.go.dev](https://pkg.go.dev/github.com/amireshoon/gojalali)

## License
This project is licensed under the MIT License
