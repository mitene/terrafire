package terrafire

import (
	"archive/zip"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
)

type GithubClient interface {
	GetSource(owner string, repo string, ref string) (io.ReadCloser, error)
}

type GithubClientImpl struct {
	client *github.Client
}

func NewGithubClient() GithubClient {
	token, ok := os.LookupEnv("GITHUB_ACCESS_TOKEN")
	var tc *http.Client
	if ok {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc = oauth2.NewClient(ctx, ts)
	}

	return &GithubClientImpl{
		client: github.NewClient(tc),
	}
}

func (c *GithubClientImpl) GetSource(owner string, repo string, ref string) (io.ReadCloser, error) {
	opt := github.RepositoryContentGetOptions{
		Ref: ref,
	}

	url, _, err := c.client.Repositories.GetArchiveLink(context.Background(), owner, repo, github.Zipball, &opt, true)
	if err != nil {
		return nil, err
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	resp.Header.Write(os.Stdout)

	return resp.Body, nil
}

func (*GithubClientImpl) extract(src io.Reader, dest string) error {
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
