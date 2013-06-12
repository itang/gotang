package io

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

func TestExists(t *testing.T) {
	AssertTrue(t, Exists("file.go"))
	AssertTrue(t, Exists("file_test.go"))
	AssertTrue(t, Exists("../io"))

	AssertFalse(t, Exists("some.go"))
}

func TestIsDir(t *testing.T) {
	AssertTrue(t, IsDir("../io"))

	AssertFalse(t, IsDir("file.go"))
	AssertFalse(t, IsDir("file_test.go"))
	AssertFalse(t, IsDir("some.go"))
}

func TestIsFile(t *testing.T) {
	AssertTrue(t, IsFile("file.go"))
	AssertTrue(t, IsFile("file_test.go"))

	AssertFalse(t, IsFile("../io"))
	AssertFalse(t, IsFile("some.go"))
}
