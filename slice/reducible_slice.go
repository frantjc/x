package xslice

// ReducibleSlice is an array of any and has methods that
// mimic the JavaScript array reduce functions.
//
// T1 represents the type of the array, while T2 represents the
// type of the array intended to be mapped to.
type ReducibleSlice[T1, T2 any] AnySlice[T1]

// Map creates a new array populated with the results of
// calling a provided function on every element in the calling array.
func (a ReducibleSlice[T1, T2]) Map(f func(T1, int) T2) AnySlice[T2] {
	return Map(a, f)
}

// ReduceRight applies a function against an accumulator and each value of the array
// (from right-to-left) to reduce it to a single value.
func (a ReducibleSlice[T1, T2]) ReduceRight(f func(T2, T1, int) T2, initial T2) T2 {
	return ReduceRight(a, f, initial)
}

// Reduce executes a user-supplied "reducer" callback function
// on each element of the array, in order, passing in the return
// value from the calculation on the preceding element.
// The final result of running the reducer across all
// elements of the array is a single value.
//
// The first time that the callback is run there is no
// "return value of the previous calculation". If supplied,
// an initial value may be used in its place.
// Otherwise the array element at index 0 is used
// as the initial value and iteration starts from
// the next element (index 1 instead of index 0).
func (a ReducibleSlice[T1, T2]) Reduce(f func(T2, T1, int) T2, initial T2) T2 {
	return Reduce(a, f, initial)
}

// Filter creates a new array with all elements that pass
// the test implemented by the provided function.
func (a ReducibleSlice[T1, T2]) Filter(f func(T1, int) bool) ReducibleSlice[T1, T2] {
	return ReducibleSlice[T1, T2](Filter(a, f))
}

// Reverse creates a new array by reversing the given array.
// The first array element becomes the last, and the last array element becomes the first.
func (a ReducibleSlice[T1, T2]) Reverse() ReducibleSlice[T1, T2] {
	return ReducibleSlice[T1, T2](Reverse(a))
}

// Slice returns a shallow copy of a portion
// of an array into a new array object selected
// from start to end (end not included) where
// start and end represent the index of items in that array.
// The original array will not be modified.
//
// If start<0, it is treated as distance from the end of the array.
// If end<=0, it is treated as distance from the end of the array.
func (a ReducibleSlice[T1, T2]) Slice(start, end int) ReducibleSlice[T1, T2] {
	return ReducibleSlice[T1, T2](Slice(a, start, end))
}
