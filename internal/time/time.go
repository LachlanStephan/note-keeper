package time

import (
	"strconv"
	"time"
)

type Times struct {
	Day           string
	Month         string
	Year          string
	FormattedDate string
}

func SetTimes() *Times {
	day := time.Now().Local().Weekday().String()
	month := time.Now().Local().Month().String()
	year := time.Now().Local().Year()
	yearString := strconv.Itoa(year)
	formattedDate := day + " " + month + " " + yearString

	return &Times{
		Day:           day,
		Month:         month,
		Year:          yearString,
		FormattedDate: formattedDate,
	}
}
