package time

import (
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

func TestFormat(t *testing.T) {
	testTime := time.Date(2009, time.November, 10, 15, 12, 10, int(time.Millisecond*time.Duration(300)), time.Local)

	assert.Equal(t, "2009-11-10 15:12:10", FormatDefault(testTime))
}
