package gotang

import (
	"errors"
	"github.com/bmizerany/assert"
	. "github.com/itang/gotang/test"
	"testing"
)

func TestCheckError(t *testing.T) {
	AssertNull(t, CheckError(nil))
	e := errors.New("err")
	assert.Panic(t, e, func() {
		CheckError(e)
	})
}

func TestAssertNoError(t *testing.T) {
	AssertNoError(nil)

	e := errors.New("err")
	assert.Panic(t, e, func() {
		AssertNoError(e)
	})
}

func TestAssert(t *testing.T) {
	Assert(true, "")

	e := "assertion failed: 1 + 1 should equals 2"
	assert.Panic(t, e, func() {
		Assert(1+1 != 2, "1 + 1 should equals 2")
	})
}
