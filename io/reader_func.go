package xio

type ReaderFunc func([]byte) (int, error)

func (r ReaderFunc) Read(b []byte) (int, error) {
	return r(b)
}
