package xtar

import (
	"archive/tar"
	"errors"
	"io"
	"strings"
)

var ErrEmptySubdir = errors.New("empty tarball subdirectory")

// Subdir reads the tarball from r and streams the files in
// the given subdirectory to the returned io.ReadCloser as a
// tarball with the subdirectory's path trimmed from each file's name.
//
// If the subdirectory is empty or non-existent, the returned io.ReadCloser
// is closed with ErrEmptySubdir.
func Subdir(r *tar.Reader, subdir string) io.ReadCloser {
	var (
		pr, pw              = io.Pipe()
		iw        io.Writer = pw
		found               = false
		lenSubdir           = len(subdir)
	)

	go func() {
		tw := tar.NewWriter(iw)
		err := func() error {
			for {
				f, err := r.Next()
				if errors.Is(err, io.EOF) {
					if !found {
						return ErrEmptySubdir
					}

					break
				} else if err != nil {
					return err
				}

				if !strings.HasPrefix(f.Name, subdir) {
					continue
				}

				found = true
				f.Name = f.Name[lenSubdir:]

				if f.Name == "" || f.Name == "/" {
					continue
				}

				if err := tw.WriteHeader(f); err != nil {
					return err
				}

				if _, err := io.Copy(tw, r); err != nil {
					return err
				}
			}

			return nil
		}()

		_ = pw.CloseWithError(errors.Join(err, tw.Close()))
	}()

	return pr
}
