package xio

type CloserFunc func() error

func (c CloserFunc) Close() error {
	return c()
}

type ReaderFunc func([]byte) (int, error)

func (r ReaderFunc) Read(b []byte) (int, error) {
	return r(b)
}

type WriterFunc func([]byte) (int, error)

func (w WriterFunc) Write(b []byte) (int, error) {
	return w(b)
}
