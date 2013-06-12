package test

import (
	"github.com/bmizerany/assert"
	"testing"
)

func AssertTrue(t *testing.T, got interface{}, args ...interface{}) {
	assert.Equal(t, true, got, args...)
}

func AssertFalse(t *testing.T, got interface{}, args ...interface{}) {
	assert.Equal(t, false, got, args...)
}
