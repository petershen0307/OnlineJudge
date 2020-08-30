package q1185

import "time"

func dayOfTheWeek(day int, month int, year int) string {
	dateTime := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	return dateTime.Weekday().String()
}
