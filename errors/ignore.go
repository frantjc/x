package xerrors

import "errors"

func Ignore(err error, targets ...error) error {
	for _, target := range targets {
		if errors.Is(err, target) {
			return nil
		}
	}

	return err
}
