package terrafire

import (
	"archive/zip"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
)

type GithubClient interface {
	GetSource(owner string, repo string, ref string, subDir string, dest string) error
	CreateComment(owner string, repo string, issue int, body string) error
}

type GithubClientImpl struct {
	client *github.Client
}

func NewGithubClient() GithubClient {
	token, ok := os.LookupEnv("GITHUB_TOKEN")
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

func (c *GithubClientImpl) GetSource(owner string, repo string, ref string, subDir string, dest string) error {
	opt := github.RepositoryContentGetOptions{
		Ref: ref,
	}

	url, _, err := c.client.Repositories.GetArchiveLink(context.Background(), owner, repo, github.Zipball, &opt, true)
	if err != nil {
		return err
	}

	resp, err := http.Get(url.String())
	if err != nil {
		return err
	}

	err = c.extract(resp.Body, subDir, dest)
	if err != nil {
		return err
	}

	return nil
}

func (c *GithubClientImpl) CreateComment(owner string, repo string, issue int, body string) error {
	_, _, err := c.client.Issues.CreateComment(context.Background(), owner, repo, issue, &github.IssueComment{
		Body: &body,
	})
	return err
}

func (*GithubClientImpl) extract(src io.Reader, subDir string, dest string) error {
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
		filename := filepath.Join(strings.Split(f.Name, string(os.PathSeparator))[1:]...)
		if !strings.HasPrefix(filename, subDir) {
			continue
		}

		if filename, err = filepath.Rel(subDir, filename); err != nil {
			return err
		}

		path := filepath.Join(dest, filename)

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(path, f.Mode())
			if err != nil {
				return err
			}
		} else {
			err = func() error {
				rc, err := f.Open()
				if err != nil {
					return err
				}
				defer rc.Close()

				d, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, f.Mode())
				if err != nil {
					return err
				}
				defer d.Close()

				if _, err = io.Copy(d, rc); err != nil {
					return err
				}
				return nil
			}()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
