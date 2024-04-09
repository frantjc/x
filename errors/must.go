package xerrors

func Must[T1 any](f func() (T1, error)) T1 {
	v, err := f()
	if err != nil {
		panic(err)
	}

	return v
}
