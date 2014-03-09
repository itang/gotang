package gotang

import (
	"errors"
	"github.com/bmizerany/assert"
	"testing"
)

func TestAssertNoError(t *testing.T) {
	AssertNoError(nil, "")

	assert.Panic(t, "assertion failed: should no error (err)", func() {
		e := errors.New("err")
		AssertNoError(e, "should no error")
	})
}

func TestAssert(t *testing.T) {
	Assert(true, "")

	e := "assertion failed: 1 + 1 should equals 2"
	assert.Panic(t, e, func() {
		Assert(1+1 != 2, "1 + 1 should equals 2")
	})
}
