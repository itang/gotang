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

func TestMustNoError(t *testing.T) {
	MustNoError(nil)

	e := errors.New("err")
	assert.Panic(t, e, func() {
		MustNoError(e)
	})
}
