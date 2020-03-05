package terrafire

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Unzip(src io.Reader, dest string) error {

	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name()) // clean up

	_, err = io.Copy(tmpfile, src)
	if err != nil {
		return err
	}

	if err = tmpfile.Close(); err != nil {
		return err
	}

	r, err := zip.OpenReader(tmpfile.Name())
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		if f.FileInfo().IsDir() {
			path := filepath.Join(dest, f.Name)
			os.MkdirAll(path, f.Mode())
		} else {
			buf := make([]byte, f.UncompressedSize64)
			_, err = io.ReadFull(rc, buf)
			if err != nil {
				return err
			}

			path := filepath.Join(dest, f.Name)
			if err = ioutil.WriteFile(path, buf, f.Mode()); err != nil {
				return err
			}
		}
	}

	return nil
}
