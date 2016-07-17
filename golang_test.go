package gotang

import (
	"testing"
	"time"
	//"github.com/bmizerany/assert"
)

func TestTime(t *testing.T) {
	Time(func() {
		time.Sleep(time.Millisecond * 50)
	})
}
