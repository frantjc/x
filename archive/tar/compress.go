package xtar

import (
	"archive/tar"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func Compress(dir string) io.ReadCloser {
	pr, pw := io.Pipe()

	go func() {
		tw := tar.NewWriter(pw)
		defer tw.Close()

		if err := filepath.WalkDir(dir, func(path string, di fs.DirEntry, err error) error {
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
		}); err != nil {
			_ = pw.CloseWithError(err)
		}
	}()

	return pr
}
