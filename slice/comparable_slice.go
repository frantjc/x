package xslice

// ComparableSlice is an array of comparable and has methods that
// mimic a subset of the JavaScript array functions that are
// unique to comparables.
type ComparableSlice[T comparable] AnySlice[T]

// Includes determines whether an array includes a certain value
// among its entries, returning true or false as appropriate.
func (a ComparableSlice[T]) Includes(b T) bool {
	return Includes(a, b)
}

// IndexOf returns the first index at which a given element
// can be found in the array, or -1 if it is not present.
func (a ComparableSlice[T]) IndexOf(b T) int {
	return IndexOf(a, b)
}

// LastIndexOf returns the last index at which a given element
// can be found in the array, or -1 if it is not present.
// The array is searched backwards, starting at fromIndex.
func (a ComparableSlice[T]) LastIndexOf(b T, from int) int {
	return LastIndexOf(a, b, from)
}

// Unique creates a new array with only unique elements from the given array.
func (a ComparableSlice[T]) Unique() ComparableSlice[T] {
	return Unique(a)
}

// Filter creates a new array with all elements that pass
// the test implemented by the provided function.
func (a ComparableSlice[T]) Filter(f func(T, int) bool) ComparableSlice[T] {
	return ComparableSlice[T](Filter(a, f))
}

// Reverse creates a new array by reversing the given array.
// The first array element becomes the last, and the last array element becomes the first.
func (a ComparableSlice[T]) Reverse() ComparableSlice[T] {
	return ComparableSlice[T](Reverse(a))
}

// Slice returns a shallow copy of a portion
// of an array into a new array object selected
// from start to end (end not included) where
// start and end represent the index of items in that array.
// The original array will not be modified.
//
// If start<0, it is treated as distance from the end of the array.
// If end<=0, it is treated as distance from the end of the array.
func (a ComparableSlice[T]) Slice(start, end int) ComparableSlice[T] {
	return ComparableSlice[T](Slice(a, start, end))
}
