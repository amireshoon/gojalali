package main

import (
	"fmt"
	"time"

	"github.com/amireshoon/gojalali/jalali"
)

func main() {
	now := time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC)
	year, month, day := jalali.ToJalali(now)
	fmt.Printf("Gregorian: %s\n", now.Format("2006-01-02"))
	fmt.Printf("Jalali: %d/%02d/%02d\n", year, month, day)

	// Convert back to Gregorian
	gregorian := jalali.ToGregorian(year, month, day)
	fmt.Printf("Back to Gregorian: %s\n", gregorian.Format("2006-01-02"))
}
