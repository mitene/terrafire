package terrafire

type GithubClientMock struct {
	getSource func(owner string, repo string, ref string, subDir string, dest string) error
}

func (c *GithubClientMock) GetSource(owner string, repo string, ref string, subDir string, dest string) error {
	if c.getSource != nil {
		return c.getSource(owner, repo, ref, subDir, dest)
	}
	return nil
}
