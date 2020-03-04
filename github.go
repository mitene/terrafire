package terrafire

import (
	"context"
	"fmt"
	"github.com/google/go-github/v29/github"
)

func GetSource(owner string, repo string, ref string) error {
	client := github.NewClient(nil)
	url, _, err := client.Repositories.GetArchiveLink(context.Background(), owner, repo, github.Zipball, nil, true)
	if err != nil {
		return err
	}
	fmt.Println(url)
	return nil
}
