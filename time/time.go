package time

import (
	"time"
)

func FormatDefault(t time.Time) string {
	return time.Time.Format(t, ChinaDefaultDateTime)
}

type RichTime time.Time

func (self RichTime) Yesterday() time.Time {
	var t = time.Time(self)
	return t.AddDate(0, 0, -1)
}
