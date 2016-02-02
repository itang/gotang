package net

import (
	. "github.com/itang/gotang/test"
	"strings"
	"testing"
)

func TestLookupWlanIP4addr(t *testing.T) {
	ip, _ := LookupWlanIP4addr()

	if ip != "" {
		AssertTrue(t, !strings.HasPrefix(ip, "127"))
	}
}
