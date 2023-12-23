package xslice_test

import (
	"testing"

	"github.com/frantjc/x/slice"
)

func TestCoalesceTrue(t *testing.T) {
	var (
		expected = "default"
		actual   = xslice.Coalesce("", "", "", expected, "")
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestCoalesceFalse(t *testing.T) {
	var (
		expected = ""
		actual   = xslice.Coalesce("", "", "")
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}
