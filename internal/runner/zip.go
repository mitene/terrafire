package runner

import (
	"archive/zip"
	"fmt"
	"github.com/mitene/terrafire/internal/utils"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Zip(w io.Writer, dir string) error {
	zw := zip.NewWriter(w)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		src, err := os.Open(path)
		if err != nil {
			return err
		}
		defer utils.LogDefer(src.Close)

		name, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		header := &zip.FileHeader{
			Name:   name,
			Method: zip.Deflate,
		}
		header.SetMode(info.Mode())
		dst, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}

		_, err = io.Copy(dst, src)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return zw.Close()
}

func Unzip(r io.Reader, path string) error {
	buf, err := utils.TempFile()
	if err != nil {
		return err
	}
	defer utils.TempClean(buf.Name())

	_, err = io.Copy(buf, r)
	if err != nil {
		return err
	}

	err = buf.Sync()
	if err != nil {
		return err
	}

	_, err = buf.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	fi, err := os.Stat(buf.Name())
	if err != nil {
		return err
	}

	zr, err := zip.NewReader(buf, fi.Size())
	if err != nil {
		return err
	}

	for _, f := range zr.File {
		err = func() error {
			src, err := f.Open()
			if err != nil {
				return err
			}
			defer utils.LogDefer(src.Close)

			name := filepath.Clean(f.Name)
			if strings.HasPrefix(name, "..") || strings.HasPrefix(name, "/") {
				return fmt.Errorf("invalid file path: %s", name)
			}

			fp := filepath.Join(path, f.Name)

			err = os.MkdirAll(filepath.Dir(fp), 0755)
			if err != nil {
				return err
			}

			dst, err := os.Create(fp)
			if err != nil {
				return err
			}
			defer utils.LogDefer(dst.Close)

			err = dst.Chmod(f.Mode())
			if err != nil {
				return err
			}

			_, err = io.Copy(dst, src)
			if err != nil {
				return err
			}

			return nil
		}()
		if err != nil {
			return err
		}
	}

	return nil
}
