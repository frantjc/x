package xio

import "io"

type ReadCloser struct {
	io.Reader
	io.Closer
}

type WriterCloser struct {
	io.Writer
	io.Closer
}

type ReadWriter struct {
	io.Reader
	io.Writer
}
