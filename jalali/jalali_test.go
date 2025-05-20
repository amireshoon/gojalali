package jalali

import (
	"testing"
	"time"
)

func TestToJalali(t *testing.T) {
	tests := []struct {
		gregorianDate time.Time
		expectedJY    int
		expectedJM    int
		expectedJD    int
	}{
		{time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC), 1402, 1, 1},
		{time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC), 1403, 1, 1}, // Nowruz 1403
		{time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC), 1401, 7, 18},
		{time.Date(2020, 2, 29, 0, 0, 0, 0, time.UTC), 1398, 12, 10}, // Gregorian leap year
	}

	for _, test := range tests {
		jy, jm, jd := ToJalali(test.gregorianDate)
		if jy != test.expectedJY || jm != test.expectedJM || jd != test.expectedJD {
			t.Errorf("ToJalali(%v) = %d/%02d/%02d; want %d/%02d/%02d",
				test.gregorianDate, jy, jm, jd, test.expectedJY, test.expectedJM, test.expectedJD)
		}
	}
}

func TestToGregorian(t *testing.T) {
	tests := []struct {
		jy           int
		jm           int
		jd           int
		expectedDate time.Time
	}{
		{1402, 1, 1, time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC)},
		{1402, 12, 30, time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC)}, // Jalali leap
		{1401, 7, 18, time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC)},
		{1398, 12, 10, time.Date(2020, 2, 29, 0, 0, 0, 0, time.UTC)},
	}

	for _, test := range tests {
		g := ToGregorian(test.jy, test.jm, test.jd)
		if g.Year() != test.expectedDate.Year() || g.Month() != test.expectedDate.Month() || g.Day() != test.expectedDate.Day() {
			t.Errorf("ToGregorian(%d, %d, %d) = %v; want %v",
				test.jy, test.jm, test.jd, g.Format("2006-01-02"), test.expectedDate.Format("2006-01-02"))
		}
	}
}

func TestToJalaliFormat(t *testing.T) {
	tests := []struct {
		input    time.Time
		format   string
		expected string
	}{
		{time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC), "yyyy/mm/dd", "1403/01/01"},
		{time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC), "yy-m-d", "03-1-1"},
		{time.Date(2023, 3, 21, 0, 0, 0, 0, time.UTC), "yyyy.mm.dd", "1402.01.01"},
		{time.Date(2022, 10, 10, 0, 0, 0, 0, time.UTC), "d/m/yyyy", "18/7/1401"},
	}

	for _, test := range tests {
		result := ToJalaliFormat(test.input, test.format)
		if result != test.expected {
			t.Errorf("ToJalaliFormat(%v, %q) = %q; want %q",
				test.input, test.format, result, test.expected)
		}
	}
}

func TestNewJalaliTime(t *testing.T) {
	jt := NewJalaliTime(1402, 12, 29)
	want := "2024-03-19"
	got := jt.ToTime().Format("2006-01-02")
	if got != want {
		t.Errorf("NewJalaliTime(1402,12,30) = %v; want %v", got, want)
	}
}

func TestNow(t *testing.T) {
	jt := Now()
	if jt.Year < 1400 || jt.Year > 1600 {
		t.Errorf("Now() returned Jalali year out of range: %v", jt.Year)
	}
}

func TestJalaliTimeFormat(t *testing.T) {
	tm := time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC)
	jy, jm, jd := ToJalali(tm)
	jt := JalaliTime{Year: jy, Month: jm, Day: jd, time: tm}

	got := jt.Format("yyyy/mm/dd")
	want := "1403/01/01"
	if got != want {
		t.Errorf("Format = %v; want %v", got, want)
	}
}

func TestJalaliMonthName(t *testing.T) {
	jt := NewJalaliTime(1402, 12, 29) // 2024-03-19
	// which is the last day of Esfand
	want := "اسفند"
	got := jt.MonthName()
	if got != want {
		t.Errorf("MonthName = %v; want %v", got, want)
	}
}

func TestJalaliWeekdayName(t *testing.T) {
	jt := NewJalaliTime(1402, 12, 29) // 2024-03-19
	// which is a Tuesday
	want := "سه‌شنبه"
	got := jt.WeekdayName()
	if got != want {
		t.Errorf("WeekdayName = %v; want %v", got, want)
	}
}

func TestAddDays(t *testing.T) {
	jt := NewJalaliTime(1402, 12, 29)
	next := jt.AddDays(1)
	if next.Year != 1403 || next.Month != 1 || next.Day != 1 {
		t.Errorf("AddDays(1) failed: got %d/%02d/%02d", next.Year, next.Month, next.Day)
	}
}

func TestRemoveDay(t *testing.T) {
	jt := NewJalaliTime(1402, 12, 29)
	prev := jt.RemoveDays(1)
	if prev.Year != 1402 || prev.Month != 12 || prev.Day != 28 {
		t.Errorf("RemoveDays(1) failed: got %d/%02d/%02d", prev.Year, prev.Month, prev.Day)
	}
}

func TestAddWeeks(t *testing.T) {
	jt := NewJalaliTime(1402, 12, 29)
	next := jt.AddWeeks(1)
	if next.Year != 1403 || next.Month != 1 || next.Day != 7 {
		t.Errorf("AddWeeks(1) failed: got %d/%02d/%02d", next.Year, next.Month, next.Day)
	}
}

func TestRemoveWeeks(t *testing.T) {
	jt := NewJalaliTime(1402, 12, 29)
	prev := jt.RemoveWeeks(1)
	if prev.Year != 1402 || prev.Month != 12 || prev.Day != 22 {
		t.Errorf("RemoveWeeks(1) failed: got %d/%02d/%02d", prev.Year, prev.Month, prev.Day)
	}
}

func TestAddMonths(t *testing.T) {
	jt := NewJalaliTime(1402, 12, 29)
	next := jt.AddMonths(1)
	if next.Year != 1403 || next.Month != 1 || next.Day != 29 {
		t.Errorf("AddMonths(1) failed: got %d/%02d/%02d", next.Year, next.Month, next.Day)
	}
}

func TestRemoveMonths(t *testing.T) {
	jt := NewJalaliTime(1402, 12, 29)
	prev := jt.RemoveMonths(1)
	if prev.Year != 1402 || prev.Month != 11 || prev.Day != 29 {
		t.Errorf("RemoveMonths(1) failed: got %d/%02d/%02d", prev.Year, prev.Month, prev.Day)
	}
}

func TestAddYears(t *testing.T) {
	jt := NewJalaliTime(1402, 12, 29)
	next := jt.AddYears(1)
	if next.Year != 1403 || next.Month != 12 || next.Day != 29 {
		t.Errorf("AddYears(1) failed: got %d/%02d/%02d", next.Year, next.Month, next.Day)
	}
}

func TestRemoveYears(t *testing.T) {
	jt := NewJalaliTime(1402, 12, 29)
	prev := jt.RemoveYears(1)
	if prev.Year != 1401 || prev.Month != 12 || prev.Day != 29 {
		t.Errorf("RemoveYears(1) failed: got %d/%02d/%02d", prev.Year, prev.Month, prev.Day)
	}
}
