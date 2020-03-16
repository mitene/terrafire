package terrafire

import (
	"fmt"
	"os"
	"strconv"
)

type Reporter interface {
	Report(PlanResults) error
}

type ReporterGithub struct {
	github GithubClient
}

func NewReporterGithub(github GithubClient) Reporter {
	return &ReporterGithub{
		github: github,
	}
}

func (r *ReporterGithub) Report(planResults PlanResults) error {
	owner, ok := os.LookupEnv("TERRAFIRE_REPORT_GITHUB_OWNER")
	if !ok {
		return fmt.Errorf("TERRAFIRE_REPORT_GITHUB_OWNER is not set")
	}

	repo, ok := os.LookupEnv("TERRAFIRE_REPORT_GITHUB_REPO")
	if !ok {
		return fmt.Errorf("TERRAFIRE_REPORT_GITHUB_REPO is not set")
	}

	issue, ok := os.LookupEnv("TERRAFIRE_REPORT_GITHUB_ISSUE")
	if !ok {
		return fmt.Errorf("TERRAFIRE_REPORT_GITHUB_ISSUE is not set")
	}
	issueNumber, err := strconv.Atoi(issue)
	if err != nil {
		return err
	}

	err = r.github.CreateComment(owner, repo, issueNumber, r.formatBody(planResults))
	if err != nil {
		return err
	}

	return nil
}

func (r *ReporterGithub) formatBody(results PlanResults) string {
	body := ""

	for _, result := range results {
		body = body + fmt.Sprintf("```\n%s\n```", result.Body)
	}

	return body
}
