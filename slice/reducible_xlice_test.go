package xslice_test

import (
	"testing"

	xslice "github.com/frantjc/x/slice"
)

func TestMappableMap(t *testing.T) {
	var (
		a = xslice.ReducibleSlice[int, int]([]int{1, 2, 3, 4})
		f = func(a, i int) int {
			return a + 1
		}
		expected = xslice.AnySlice[int]([]int{2, 3, 4, 5})
		actual   = a.Map(f)
	)
	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}

func TestReducibleReduceRight(t *testing.T) {
	var (
		a = xslice.ReducibleSlice[int, int]([]int{1, 2, 3, 4})
		f = func(acc, a, i int) int {
			return acc + a
		}
		expected = 10
		actual   = a.ReduceRight(f, 0)
	)
	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestReducibleReduce(t *testing.T) {
	var (
		a = xslice.ReducibleSlice[int, int]([]int{1, 2, 3, 4})
		f = func(acc, a, i int) int {
			return acc + a
		}
		expected = 10
		actual   = a.Reduce(f, 0)
	)
	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}
