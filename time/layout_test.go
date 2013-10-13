package time

import (
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

func TestChinaDateLayout(t *testing.T) {
	testTime := time.Date(2009, time.November, 10, 15, 12, 10, int(time.Millisecond*time.Duration(300)), time.Local)

	assert.Equal(t, "2009-11-10 15:12:10.3", testTime.Format(ChinaDefaultDateTimeFull))
	assert.Equal(t, "2009-11-10 15:12:10", testTime.Format(ChinaDefaultDateTime))
	assert.Equal(t, "2009-11-10", testTime.Format(ChinaDefaultDate))
	assert.Equal(t, "15:12:10", testTime.Format(ChinaDefaultTime))
}
