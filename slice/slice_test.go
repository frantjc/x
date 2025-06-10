package xslice_test

import (
	"testing"

	xslice "github.com/frantjc/x/slice"
)

func TestEveryTrue(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4}
		f = func(a, _ int) bool {
			return a > 0
		}
		expected = true
		actual   = xslice.Every(a, f)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestEveryFalse(t *testing.T) {
	var (
		a = []int{0, 1, 2, 3}
		f = func(a, _ int) bool {
			return a > 0
		}
		expected = false
		actual   = xslice.Every(a, f)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestFilter(t *testing.T) {
	var (
		a = []int{0, 1, 2, 3}
		f = func(a, _ int) bool {
			return a > 0
		}
		expected = []int{1, 2, 3}
		actual   = xslice.Filter(a, f)
	)

	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}

func TestFindIndex(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4}
		f = func(a, _ int) bool {
			return a == 2
		}
		expected = 1
		actual   = xslice.FindIndex(a, f)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestFind(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4}
		f = func(a, _ int) bool {
			return a > 0
		}
		expected = 1
		actual   = xslice.Find(a, f)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestIncludesTrue(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4}
		expected = true
		actual   = xslice.Includes(a, 1)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestIncludesFalse(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4}
		expected = false
		actual   = xslice.Includes(a, 0)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestIndexOf(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4}
		expected = 0
		actual   = xslice.IndexOf(a, 1)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestLastIndexOf(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4, 3, 2}
		expected = 4
		actual   = xslice.LastIndexOf(a, 3, 0)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestLastIndexOfFrom(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4, 3, 2}
		expected = 2
		actual   = xslice.LastIndexOf(a, 3, 3)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestMap(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4, 3, 2}
		f = func(a, _ int) int {
			return a + 1
		}
		expected = []int{2, 3, 4, 5, 4, 3}
		actual   = xslice.Map(a, f)
	)

	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}

func TestReduceRight(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4, 3, 2}
		f = func(acc, a, _ int) int {
			return acc + a
		}
		expected = 15
		actual   = xslice.ReduceRight(a, f, 0)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestReduce(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4, 3, 2}
		f = func(acc, a, _ int) int {
			return acc + a
		}
		expected = 15
		actual   = xslice.Reduce(a, f, 0)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestReverse(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4, 3, 2}
		expected = []int{2, 3, 4, 3, 2, 1}
		actual   = xslice.Reverse(a)
	)

	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}

func TestSliceFromStart(t *testing.T) {
	var (
		a        = []int{0, 1, 2, 3}
		expected = []int{0, 1}
		actual   = xslice.Slice(a, 0, 2)
	)

	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}

func TestSliceFromEnd(t *testing.T) {
	var (
		a        = []int{0, 1, 2, 3}
		expected = []int{2}
		actual   = xslice.Slice(a, -2, -1)
	)

	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}

func TestSomeTrue(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4}
		f = func(a, _ int) bool {
			return a == 4
		}
		expected = true
		actual   = xslice.Some(a, f)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestSomeFalse(t *testing.T) {
	var (
		a = []int{1, 2, 3, 4}
		f = func(a, _ int) bool {
			return a == 5
		}
		expected = false
		actual   = xslice.Some(a, f)
	)

	if expected != actual {
		t.Error("actual", actual, "does not equal expected", expected)
		t.FailNow()
	}
}

func TestUnique(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4, 3, 2, 1}
		expected = []int{1, 2, 3, 4}
		actual   = xslice.Unique(a)
	)

	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}

func TestSortWorst(t *testing.T) {
	var (
		a        = []int{4, 3, 2, 1}
		expected = []int{1, 2, 3, 4}
		actual   = xslice.Sort(a, func(a, b int) int {
			return a - b
		})
	)

	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}

func TestSortBest(t *testing.T) {
	var (
		a        = []int{1, 2, 3, 4}
		expected = []int{1, 2, 3, 4}
		actual   = xslice.Sort(a, func(a, b int) int {
			return a - b
		})
	)

	for i := range expected {
		if expected[i] != actual[i] {
			t.Error("actual", actual, "does not equal expected", expected)
			t.FailNow()
		}
	}
}
