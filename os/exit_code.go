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

func ErrorExitCode(err error) int {
	if err == nil {
		return 0
	}

	e := &ExitCodeError{}

	// For some reason, sometimes this return true
	// but sets e = nil, so we have to double check.
	if errors.As(err, &e) && e != nil {
		return e.ExitCode
	}

	return 1
}

func ExitFromError(err error) {
	os.Exit(ErrorExitCode(err))
}
