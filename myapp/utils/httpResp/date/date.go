package date

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

func GetDate() string {
	date := time.Now().UTC()          // get current date and time
	return date.Format(apiDateLayout) // convert to string
}
