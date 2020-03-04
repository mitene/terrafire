package terrafire

import (
	"archive/zip"
	"context"
	"github.com/google/go-github/v29/github"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// TODO: GET CLEAN!!!
func GetSource(owner string, repo string, ref string) error {
	client := github.NewClient(nil)
	url, _, err := client.Repositories.GetArchiveLink(context.Background(), owner, repo, github.Zipball, nil, true)
	if err != nil {
		return err
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return err
	}

	tmpfile, err := ioutil.TempFile("", "")
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name()) // clean up

	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		return err
	}

	tmpfile.Close()

	Unzip(tmpfile.Name(), "out")

	return nil
}

func Unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
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
			buf := make([]byte, f.UncompressedSize)
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