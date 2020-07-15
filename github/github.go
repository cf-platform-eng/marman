package github

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"

	"github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
)

//go:generate counterfeiter Client
type Client interface {
	ListReleases(owner, repo string, opt *github.ListOptions) ([]*github.RepositoryRelease, *github.Response, error)
	DownloadReleaseAsset(owner, repo string, id int64) (io.ReadCloser, string, error)
}

type GitHubClient struct {
	Client  *github.Client
	context context.Context
}

func NewGitHubClient(token, gitHubBaseURL string) (*GitHubClient, error) {
	ctx := context.Background()
	var (
		httpClient   *http.Client
		gitHubClient *github.Client
		err          error
	)

	if token != "" {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		httpClient = oauth2.NewClient(ctx, ts)
	}

	gitHubClient = github.NewClient(httpClient)
	gitHubClient.BaseURL, err = url.Parse(gitHubBaseURL)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse Github BaseURL %s", gitHubBaseURL)
	}

	return &GitHubClient{
		Client:  gitHubClient,
		context: ctx,
	}, nil
}

func (c *GitHubClient) ListReleases(owner, repo string, opt *github.ListOptions) ([]*github.RepositoryRelease, *github.Response, error) {
	releases, response, err := c.Client.Repositories.ListReleases(c.context, owner, repo, opt)
	return releases, response, err
}

func (c *GitHubClient) DownloadReleaseAsset(owner, repo string, id int64) (io.ReadCloser, string, error) {
	return c.Client.Repositories.DownloadReleaseAsset(c.context, owner, repo, id)
}
