package xos

import (
	"errors"
	"os"
)

type ExitCodeError struct {
	Err      error
	ExitCode int
}

func (e *ExitCodeError) Error() string {
	if e != nil && e.Err != nil {
		return e.Err.Error()
	}

	return ""
}

func (e *ExitCodeError) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Err
}

func NewExitCodeError(err error, exitCode int) *ExitCodeError {
	e := &ExitCodeError{}

	switch {
	case err == nil:
		return nil
	case errors.As(err, &e):
		e.ExitCode = exitCode
	default:
		e = &ExitCodeError{err, exitCode}
	}

	return e
}

func ExitFromError(err error) {
	if err == nil {
		os.Exit(0)
	}

	e := &ExitCodeError{}

	if errors.As(err, &e) {
		os.Exit(e.ExitCode)
	}

	os.Exit(1)
}
