package date

import "time"

func NowToString() string {
	return time.Now().Format("2006-01-02T15:04:05Z07:00")
}
