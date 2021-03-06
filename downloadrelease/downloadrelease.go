package downloadrelease

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"code.cloudfoundry.org/lager"
	"github.com/cf-platform-eng/marman"
	github2 "github.com/cf-platform-eng/marman/github"
	"github.com/google/go-github/v25/github"
	"github.com/pkg/errors"
)

type Config struct {
	Owner   string `short:"o" long:"owner" description:"Repository owner" required:"true"`
	Repo    string `short:"r" long:"repo" description:"Repository name" required:"true"`
	Version string `short:"v" long:"version" description:"Release version (default: latest GA)"`
	Filter  string `short:"f" long:"filter" description:"Filter to specific asset"`

	GithubToken   string `long:"github-token" description:"Authentication token for GitHub" env:"GITHUB_TOKEN"`
	GithubNetHost string `long:"github-host" description:"Base URL host for Github, (e.g. https://enterprise.example.com:8443/)" env:"GITHUB_NETWORK_HOSTNAME" default:"https://api.github.com/"`

	Logger       lager.Logger
	GithubClient github2.Client
	Downloader   marman.Downloader
}

func assetsToString(assets []github.ReleaseAsset, filter string) string {
	builder := strings.Builder{}
	for _, asset := range assets {
		matched, _ := regexp.MatchString(filter, *asset.Name)
		if matched {
			builder.WriteString(fmt.Sprintf("\n    %s", *asset.Name))
		}
	}
	return builder.String()
}

func (cmd *Config) findRelease(releases []*github.RepositoryRelease) *github.RepositoryRelease {
	for _, release := range releases {
		if cmd.Version == "" && !*release.Prerelease {
			return release
		} else if cmd.Version == *release.Name {
			return release
		}
	}
	return nil
}

func (cmd *Config) DownloadRelease() error {
	listOpt := &github.ListOptions{
		Page: 1,
	}
	var chosenRelease *github.RepositoryRelease

	for chosenRelease == nil {
		releases, response, err := cmd.GithubClient.ListReleases(cmd.Owner, cmd.Repo, listOpt)
		if err != nil {
			githubError, ok := err.(*github.ErrorResponse)
			if ok && githubError.Message == "Not Found" {
				if cmd.GithubToken == "" {
					return fmt.Errorf("could not find %s/%s. If this repository is private, try again with a GitHub token", cmd.Owner, cmd.Repo)
				} else {
					return fmt.Errorf("could not find %s/%s", cmd.Owner, cmd.Repo)
				}
			}
			return errors.Wrapf(err, "failed to get the list of releases for %s/%s", cmd.Owner, cmd.Repo)
		}

		if len(releases) == 0 {
			return fmt.Errorf("no releases found for %s/%s", cmd.Owner, cmd.Repo)
		}

		chosenRelease = cmd.findRelease(releases)
		if chosenRelease == nil {
			if response.NextPage > listOpt.Page {
				listOpt = &github.ListOptions{
					Page: response.NextPage,
				}
			} else {
				return fmt.Errorf("no releases found for %s/%s with version %s", cmd.Owner, cmd.Repo, cmd.Version)
			}
		}
	}

	if len(chosenRelease.Assets) == 0 {
		return fmt.Errorf("no release assets found for %s/%s", cmd.Owner, cmd.Repo)
	}

	chosenAsset := github.ReleaseAsset{}
	for _, asset := range chosenRelease.Assets {
		matched, err := regexp.MatchString(cmd.Filter, *asset.Name)
		if err != nil {
			return errors.Wrapf(err, "failed to apply filter: %s", cmd.Filter)
		}

		if matched {
			if chosenAsset.ID != nil {
				return fmt.Errorf("multiple assets found. Please use a filter:%s", assetsToString(chosenRelease.Assets, cmd.Filter))
			}
			chosenAsset = asset
		}
	}

	if chosenAsset.ID == nil {
		return fmt.Errorf("no release assets found for %s/%s with given filter: %s", cmd.Owner, cmd.Repo, cmd.Filter)
	}

	filename := path.Base(chosenAsset.GetBrowserDownloadURL())
	body, redirectURL, err := cmd.GithubClient.DownloadReleaseAsset(cmd.Owner, cmd.Repo, *chosenAsset.ID)
	if err != nil {
		return errors.Wrapf(err, "failed to download release asset %s from %s/%s", *chosenAsset.Name, cmd.Owner, cmd.Repo)
	}

	fmt.Printf("Downloading %s from %s/%s %s...\n", filename, cmd.Owner, cmd.Repo, *chosenRelease.Name)

	if body == nil && redirectURL != "" {
		err = cmd.Downloader.DownloadFromURL(filename, redirectURL)
	} else {
		err = cmd.Downloader.DownloadFromReader(filename, body)
	}
	if err != nil {
		return errors.Wrapf(err, "failed to save release asset %s from %s/%s", *chosenAsset.Name, cmd.Owner, cmd.Repo)
	}
	return nil
}

func (cmd *Config) Execute(args []string) error {
	cmd.Downloader = &marman.MarmanDownloader{}
	var err error
	cmd.GithubClient, err = github2.NewGitHubClient(cmd.GithubToken, cmd.GithubNetHost)

	if err != nil {
		return errors.Wrapf(err, "failed to create GitHubClient for host %s", cmd.GithubNetHost)
	}

	return cmd.DownloadRelease()
}
