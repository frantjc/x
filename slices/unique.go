package xslices

import "slices"

// Unique creates a new array with only unique elements from the given array.
func Unique[T comparable](in []T) []T {
	i := -1

	return slices.DeleteFunc(in, func(a T) bool {
		i++
		return i != IndexOf(in, a)
	})
}
