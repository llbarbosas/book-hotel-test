package main

import (
	"log"
	"strings"
	"time"

	"github.com/goodsign/monday"
)

const (
	defaultDateFormat = "2 de January de 2006"
)

func daysBetween(start, end time.Time) []time.Time {
	nDays := int(end.Sub(start).Hours()/24) + 1
	days := make([]time.Time, nDays)

	days[0] = start

	for i := range days[1:] {
		days[i+1] = days[i].Add(time.Hour * 24)
	}

	days[nDays-1] = end

	return days
}

func Date(day, month, year int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

func parseDate(dateStr string) time.Time {
	dateStr = strings.Trim(dateStr, "\n")

	date, err := monday.Parse(defaultDateFormat, dateStr, monday.LocalePtBR)

	if err != nil {
		log.Fatal(err)
	}

	return date
}
