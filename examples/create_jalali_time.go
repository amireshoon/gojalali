package main

import (
	"fmt"

	"github.com/amireshoon/gojalali/jalali"
)

func main() {
	jt := jalali.NewJalaliTime(1402, 12, 29)
	fmt.Printf("Jalali Date: %d/%02d/%02d\n", jt.Year, jt.Month, jt.Day)
}
