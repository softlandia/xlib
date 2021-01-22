package xlib

import (
	"fmt"
	"math"
	"time"
)

// NowB - returns the beginning of today
func NowB() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

// WeekdayR - Сокращённое название дня недели на русском языке
func WeekdayR(t time.Time) string {
	days := []string{
		"Вс",
		"Пн",
		"Вт",
		"Ср",
		"Чт",
		"Пт",
		"Сб",
	}
	wd := t.Weekday()
	return days[wd]
}

// FirstDay - первый день текущего года
func FirstDay() time.Time {
	currentYear, _, _ := time.Now().Date()
	firstDate, _ := time.Parse("2006-01-02", fmt.Sprintf("%d-01-01", currentYear))
	return firstDate
}

// LastDay - последний день текущего года
func LastDay() time.Time {
	currentYear, _, _ := time.Now().Date()
	lastDate, _ := time.Parse("2006-01-02", fmt.Sprintf("%d-12-31", currentYear))
	return lastDate
}

// Index - номер дня t в текущем году, первый день года имеет индекс 0
func Index(t time.Time) int {
	return int(math.Round(t.Sub(FirstDay()).Hours() / 24))
}

// Day - получить день в текущем году по его номеру, первый день года имеет индекс 0
func Day(i int) time.Time {
	if (i < 0) || (i >= DayCount()) {
		return time.Time{}
	}
	return FirstDay().AddDate(0, 0, i)
}

// DayCount - количество дней в текущем году
func DayCount() int {
	return Index(LastDay()) + 1
}

// Yesterday - ровно на начало вчера
func Yesterday() time.Time {
	now, _ := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	return now.AddDate(0, 0, -1) // целевым является вчерашний день
}
