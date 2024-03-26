package time

import (
	"strconv"
	"time"
)

type Times struct {
	day           string
	month         string
	year          string
	formattedDate string
}

func SetTimes() *Times {
	day := time.Now().Local().Weekday().String()
	month := time.Now().Local().Month().String()
	year := time.Now().Local().Year()
	yearString := strconv.Itoa(year)
	formattedDate := day + " " + month + " " + yearString

	return &Times{
		day:           day,
		month:         month,
		year:          yearString,
		formattedDate: formattedDate,
	}
}
