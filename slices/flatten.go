package xslices

func Flatten[T any](in... []T) []T {
	out := []T{}
	for _, i := range in {
		out = append(out, i...)
	}
	return out
}
