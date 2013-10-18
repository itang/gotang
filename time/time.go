package time

import (
	"time"
)

func FormatDefault(t time.Time) string {
	return time.Time.Format(t, ChinaDefaultDateTime)
}
