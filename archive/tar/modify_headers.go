package xtar

import (
	"archive/tar"
	"errors"
	"io"
)

func ModifyHeaders(r *tar.Reader, fns ...func(*tar.Header)) io.ReadCloser {
	pr, pw := io.Pipe()

	go func() {
		tw := tar.NewWriter(pw)
		defer pw.Close()
		defer tw.Close()

		if err := func() error {
			for {
				hdr, err := r.Next()
				if errors.Is(err, io.EOF) {
					return nil
				} else if err != nil {
					return err
				}

				for _, fn := range fns {
					fn(hdr)
				}

				if err = tw.WriteHeader(hdr); err != nil {
					return err
				}

				if _, err = io.Copy(tw, r); err != nil {
					return err
				}
			}
		}(); err != nil {
			_ = pw.CloseWithError(err)
		}
	}()

	return pr
}
