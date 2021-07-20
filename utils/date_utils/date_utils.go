package dateutils

import "time"

const (
	apiDateLayout = "2006-01-30T15:04:05Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
