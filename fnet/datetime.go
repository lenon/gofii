package fnet

import (
	"fmt"
	"time"
)

const (
	DATE_FORMAT_MY     = "2"
	DATE_FORMAT_DMY    = "3"
	DATE_FORMAT_DMY_HM = "4"
)

var dateFormats = map[string]string{
	DATE_FORMAT_MY:     "01/2006",
	DATE_FORMAT_DMY:    "02/01/2006",
	DATE_FORMAT_DMY_HM: "02/01/2006 15:04",
}

func ParseDate(date string, format string) (time.Time, error) {
	layout, ok := dateFormats[format]
	if !ok {
		return time.Time{}, fmt.Errorf("unknown format: %s", format)
	}
	return time.ParseInLocation(layout, date, time.Local)
}
