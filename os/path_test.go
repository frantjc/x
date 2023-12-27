package xos_test

import (
	"testing"

	xos "github.com/frantjc/x/os"
)

func TestJoinPath(t *testing.T) {
	var (
		expected = "/bin:/usr/bin"
		actual   = xos.JoinPath("/bin", "/usr/bin", "", "/usr/bin")
	)

	if actual != expected {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}
