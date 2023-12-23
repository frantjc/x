package xtar

import (
	"archive/tar"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var (
	// ModTime is used as the ModTime of each
	// tar.Header so as to make the resulting
	// tarball have the same digest.
	ModTime = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC).
		Add(func() time.Duration {
			sourceDateEpoch := os.Getenv("SOURCE_DATE_EPOCH")
			if sourceDateEpoch == "" {
				return 0
			}

			if seconds, err := strconv.Atoi(sourceDateEpoch); err == nil {
				return time.Second * time.Duration(seconds)
			}

			return 0
		}())
)

func Compress(dir string) io.ReadCloser {
	var (
		pr, pw = io.Pipe()
		tw     = tar.NewWriter(pw)
	)

	go func() {
		err := filepath.WalkDir(dir, func(path string, di fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			fi, err := di.Info()
			if err != nil {
				return err
			}

			hdr, err := tar.FileInfoHeader(fi, fi.Name())
			if err != nil {
				return err
			}
			hdr.ModTime = ModTime

			rel, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}
			hdr.Name = rel

			if err := tw.WriteHeader(hdr); err != nil {
				return err
			}

			if !fi.Mode().IsDir() {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				if _, err := io.Copy(tw, file); err != nil {
					return err
				}
			}

			return nil
		})

		_ = pw.CloseWithError(errors.Join(err, tw.Close()))
	}()

	return pr
}
