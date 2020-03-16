package terrafire

type GithubClientMock struct {
	getSource     func(owner string, repo string, ref string, subDir string, dest string) error
	createComment func(owner string, repo string, issue int, body string) error
}

func (c *GithubClientMock) GetSource(owner string, repo string, ref string, subDir string, dest string) error {
	if c.getSource != nil {
		return c.getSource(owner, repo, ref, subDir, dest)
	}
	return nil
}

func (c *GithubClientMock) CreateComment(owner string, repo string, issue int, body string) error {
	if c.createComment != nil {
		return c.createComment(owner, repo, issue, body)
	}
	return nil
}
