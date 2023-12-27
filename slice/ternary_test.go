package xslice_test

import (
	"testing"

	xslice "github.com/frantjc/x/slice"
)

func TestTernaryTrue(t *testing.T) {
	var (
		expected = 1
		actual   = xslice.Ternary(true, 1, 0)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestTernaryFalse(t *testing.T) {
	var (
		expected = 0
		actual   = xslice.Ternary(false, 1, 0)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}
