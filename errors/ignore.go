package xerrors

import "errors"

func Ignore(err, target error) error {
	if errors.Is(err, target) {
		return nil
	}

	return err
}
