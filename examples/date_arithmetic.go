package main

import (
	"fmt"

	"github.com/amireshoon/gojalali/jalali"
)

func main() {
	jt := jalali.NewJalaliTime(1402, 12, 29)

	fmt.Printf("Original: %d/%02d/%02d\n", jt.Year, jt.Month, jt.Day)

	// Add days
	addDays := jt.AddDays(5)
	fmt.Printf("Add 5 days: %d/%02d/%02d\n", addDays.Year, addDays.Month, addDays.Day)

	// Remove days
	removeDays := jt.RemoveDays(10)
	fmt.Printf("Remove 10 days: %d/%02d/%02d\n", removeDays.Year, removeDays.Month, removeDays.Day)

	// Add months
	addMonths := jt.AddMonths(1)
	fmt.Printf("Add 1 month: %d/%02d/%02d\n", addMonths.Year, addMonths.Month, addMonths.Day)

	// Remove months
	removeMonths := jt.RemoveMonths(2)
	fmt.Printf("Remove 2 months: %d/%02d/%02d\n", removeMonths.Year, removeMonths.Month, removeMonths.Day)

	// Add years
	addYears := jt.AddYears(1)
	fmt.Printf("Add 1 year: %d/%02d/%02d\n", addYears.Year, addYears.Month, addYears.Day)

	// Remove years
	removeYears := jt.RemoveYears(1)
	fmt.Printf("Remove 1 year: %d/%02d/%02d\n", removeYears.Year, removeYears.Month, removeYears.Day)
}
