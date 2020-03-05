package terrafire

import (
	"context"
	"io"
	"net/http"
	"os"

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

	return resp.Body, nil
}
