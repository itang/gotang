package time

import (
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

func TestChinaDateLayout(t *testing.T) {
	testTime := time.Date(2009, time.November, 10, 15, 12, 10, int(time.Millisecond*time.Duration(300)), time.Local)

	var tests = []struct {
		time     time.Time
		layout   string
		expected string
	}{
		{testTime, ChinaDefaultDateTimeFull, "2009-11-10 15:12:10.3"},
		{testTime, ChinaDefaultDateTime, "2009-11-10 15:12:10"},
		{testTime, ChinaDefaultDate, "2009-11-10"},
		{testTime, ChinaDefaultTime, "15:12:10"},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, test.time.Format(test.layout))
	}
}
