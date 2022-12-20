package dateutils

import "time"

const (
	DATE_FORMT = "2006-01-02T15:04:05Z"
)

func TimeNow() string {
	currentDateTime := time.Now()

	return currentDateTime.Format(DATE_FORMT)
}
