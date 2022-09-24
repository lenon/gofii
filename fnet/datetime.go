package fnet

import (
	"fmt"
	"time"
)

const (
	DateFormatMY    = "2"
	DateFormatDMY   = "3"
	DateFormatDMYHM = "4"
)

var dateFormats = map[string]string{
	DateFormatMY:    "01/2006",
	DateFormatDMY:   "02/01/2006",
	DateFormatDMYHM: "02/01/2006 15:04",
}

func ParseDate(date string, format string) (time.Time, error) {
	layout, ok := dateFormats[format]
	if !ok {
		return time.Time{}, fmt.Errorf("unknown format: %s", format)
	}
	return time.ParseInLocation(layout, date, time.Local)
}
