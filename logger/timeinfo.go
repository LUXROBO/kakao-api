package logger

import "time"

const (
	// DefaultLayout ...
	DefaultLayout = "2006-01-02T15:04:05Z"
	// DatetimeLayout datetime format
	DatetimeLayout = "2006-01-02 15:04:05"
	// DateLayout ...
	DateLayout = "2006-01-02"
	// PeriodTimeDLayout ...
	PeriodTimeDLayout = "2006.01.02"
	// DiffDateTimeLayput ...
	DiffDateTimeLayput = "2006-01-02 15:04"
)

// LocalTime ...
func LocalTime() time.Time {
	loc, _ := time.LoadLocation("Asia/Seoul")
	t := time.Now().In(loc)
	return t
}
