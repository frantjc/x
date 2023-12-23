package xslice_test

import (
	"testing"

	"github.com/frantjc/x/slice"
)

func TestComparableIncludesTrue(t *testing.T) {
	var (
		a        = xslice.ComparableSlice[int]([]int{1, 2, 3, 4})
		expected = true
		actual   = a.Includes(1)
	)
	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestComparableIncludesFalse(t *testing.T) {
	var (
		a        = xslice.ComparableSlice[int]([]int{1, 2, 3, 4})
		expected = false
		actual   = a.Includes(0)
	)
	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestComparableIndexOf(t *testing.T) {
	var (
		a        = xslice.ComparableSlice[int]([]int{1, 2, 3, 4})
		expected = 0
		actual   = a.IndexOf(1)
	)
	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestComparableLastIndexOf(t *testing.T) {
	var (
		a        = xslice.ComparableSlice[int]([]int{1, 2, 3, 4, 3, 2})
		expected = 4
		actual   = a.LastIndexOf(3, 0)
	)
	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestComparableLastIndexOfFrom(t *testing.T) {
	var (
		a        = xslice.ComparableSlice[int]([]int{1, 2, 3, 4, 3, 2})
		expected = 2
		actual   = a.LastIndexOf(3, 3)
	)
	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestComparableUnique(t *testing.T) {
	var (
		a        = xslice.ComparableSlice[int]([]int{1, 2, 3, 4, 3, 2, 1})
		expected = xslice.ComparableSlice[int]([]int{1, 2, 3, 4})
		actual   = a.Unique()
	)
	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}
