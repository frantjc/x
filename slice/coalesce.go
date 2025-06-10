package xslice

func Coalesce[T comparable](ts ...T) T {
	empty := *new(T)
	for _, t := range ts {
		if t != empty {
			return t
		}
	}
	return empty
}
