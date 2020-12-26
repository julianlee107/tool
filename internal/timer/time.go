package timer

import "time"

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Milisecond                = 1000 * Microsecond
	Second                    = 1000 * Milisecond
	Minute                    = 60 * Second
	Hour                      = 60 * Minute
	Day                       = 24 * Hour
)

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTime.Add(duration), nil
}

