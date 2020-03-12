package terrafire

type Reporter interface {
	Report(PlanResult) error
}

type ReporterGithub struct {
	github GithubClient
}

func NewReporterGithub(github GithubClient) Reporter {
	return &ReporterGithub{
		github: github,
	}
}

func (*ReporterGithub) Report(planResult PlanResult) error {

	return nil
}
