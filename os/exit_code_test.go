package xos_test

import (
	"fmt"
	"testing"

	xos "github.com/frantjc/x/os"
)

func TestExitCodeUnwrap(t *testing.T) {
	var (
		expected = 11
		err      = xos.NewExitCodeError(fmt.Errorf("test"), expected)
		actual   = xos.ErrorExitCode(err)
	)

	if actual != expected {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestExitCodeNotUnwrap(t *testing.T) {
	var (
		expected = 1
		err      = fmt.Errorf("test")
		actual   = xos.ErrorExitCode(err)
	)

	if actual != expected {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestExitCodeNil(t *testing.T) {
	var (
		expected = 0
		err      error
		actual   = xos.ErrorExitCode(err)
	)

	if actual != expected {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}
