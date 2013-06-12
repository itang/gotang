package net

import (
	. "github.com/itang/gotang/test"
	"strings"
	"testing"
)

func TestLookupWlanIP4addr(t *testing.T) {
	ip, err := LookupWlanIP4addr()
	// FIX buggly?
	if err == nil {
		AssertTrue(t, strings.HasPrefix(ip, "192"))
	}
}
