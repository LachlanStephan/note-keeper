package time

import (
	"strconv"
	"time"
)

type Times struct {
	Day           string
	Date          string
	Month         string
	Year          string
	FormattedDate string
}

func SetTimes() *Times {
	day := time.Now().Local().Weekday().String()
	date := strconv.Itoa(time.Now().Local().Day())
	month := time.Now().Local().Month().String()
	year := time.Now().Local().Year()
	yearString := strconv.Itoa(year)
	formattedDate := date + " " + month + " " + yearString + " (" + day + ")"

	return &Times{
		Day:           day,
		Date:          date,
		Month:         month,
		Year:          yearString,
		FormattedDate: formattedDate,
	}
}
