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

func AssertNull(t *testing.T, got interface{}, args ...interface{}) {
	assert.Equal(t, nil, got, args...)
}

func AssertNotNull(t *testing.T, got interface{}, args ...interface{}) {
	assert.NotEqual(t, nil, got, args...)
}
