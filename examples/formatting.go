package main

import (
	"fmt"
	"time"

	"github.com/amireshoon/gojalali/jalali"
)

func main() {
	t := time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC)

	formatted1 := jalali.ToJalaliFormat(t, "yyyy/mm/dd")
	formatted2 := jalali.ToJalaliFormat(t, "d/m/yyyy")
	formatted3 := jalali.ToJalaliFormat(t, "yyyy-mm-dd")

	fmt.Println("Formatted Jalali Dates:")
	fmt.Println("Format yyyy/mm/dd:", formatted1)
	fmt.Println("Format d/m/yyyy:", formatted2)
	fmt.Println("Format yyyy-mm-dd:", formatted3)
}
