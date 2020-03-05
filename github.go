package terrafire

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
)

// TODO: GET CLEAN!!!
func GetSource(owner string, repo string, ref string, dest string) error {
	client := newGithubClient()
	opt := github.RepositoryContentGetOptions{
		Ref: ref,
	}
	url, _, err := client.Repositories.GetArchiveLink(context.Background(), owner, repo, github.Zipball, &opt, true)
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

	Unzip(tmpfile.Name(), dest)

	return nil
}

func newGithubClient() *github.Client {
	token, ok := os.LookupEnv("GITHUB_ACCESS_TOKEN")
	var tc *http.Client
	if ok {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc = oauth2.NewClient(ctx, ts)
	}

	return github.NewClient(tc)
}
