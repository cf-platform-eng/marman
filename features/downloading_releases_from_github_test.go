//go:build feature
// +build feature

package features_test

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	. "github.com/bunniesandbeatings/goerkin"
	"github.com/cf-platform-eng/marman/features"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Downloading releases from GitHub", func() {
	steps := NewSteps()

	Scenario("download release command", func() {
		steps.Given("marman is built")
		steps.And("GitHub is running")

		steps.When("marman github-download-release -o cf-platform-eng -r needs -f linux is run")
		steps.And("it exits without error")

		steps.Then("the latest GA release is downloaded")
		steps.And("GitHub is stopped")
	})

	Scenario("download release with a version", func() {
		steps.Given("marman is built")
		steps.And("GitHub is running")

		steps.When("marman github-download-release -o cf-platform-eng -r needs -v 1.2.3 -f linux is run")
		steps.And("it exits without error")

		steps.Then("the 1.2.3 release is downloaded")
		steps.And("GitHub is stopped")
	})

	Scenario("download release that is not in the first page of results", func() {
		steps.Given("marman is built")
		steps.And("GitHub is running")

		steps.When("marman github-download-release -o cf-platform-eng -r needs -v 1.0.0 -f linux is run")
		steps.And("it exits without error")

		steps.Then("the 1.0.0 release is downloaded")
		steps.And("GitHub is stopped")
	})

	steps.Define(func(define Definitions) {
		var (
			marmanPath     string
			gitHub         *features.FakeGitHub
			commandSession *gexec.Session
		)

		define.Given(`^marman is built$`, func() {
			var err error
			marmanPath, err = gexec.Build("github.com/cf-platform-eng/marman/cmd/marman")
			Expect(err).NotTo(HaveOccurred())
		}, func() {
			gexec.CleanupBuildArtifacts()
		})

		define.When("^marman github-download-release -o cf-platform-eng -r needs -f linux is run$", func() {
			command := exec.Command(marmanPath, "github-download-release", "-o", "cf-platform-eng", "-r", "needs", "-f", "linux")
			command.Env = []string{
				fmt.Sprintf("GITHUB_NETWORK_HOSTNAME=%s", gitHub.Host),
			}

			var err error
			commandSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		define.When(`^marman github-download-release -o cf-platform-eng -r needs -v ([0-9]\.[0-9]\.[0-9]) -f linux is run$`, func(version string) {
			command := exec.Command(marmanPath, "github-download-release", "-o", "cf-platform-eng", "-r", "needs", "-v", version, "-f", "linux")
			command.Env = []string{
				fmt.Sprintf("GITHUB_NETWORK_HOSTNAME=%s", gitHub.Host),
			}

			var err error
			commandSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		define.Then(`^the ([0-9]\.[0-9]\.[0-9]) release is downloaded$`, func(version string) {
			cwd, err := os.Getwd()
			Expect(err).NotTo(HaveOccurred())
			expectedReleasePath := path.Join(cwd, fmt.Sprintf("needs-%s-linux", version))

			info, err := os.Stat(expectedReleasePath)
			Expect(err).NotTo(HaveOccurred())
			Expect(info.IsDir()).To(BeFalse())

			err = os.Remove(expectedReleasePath)
			Expect(err).NotTo(HaveOccurred())
			expectedReleasePath = ""
		})

		define.Then(`^the latest GA release is downloaded$`, func() {
			steps.Run("the 2.0.0 release is downloaded")
		})

		define.Then(`^it exits without error$`, func() {
			Eventually(commandSession).Should(gexec.Exit(0))
		})

		define.Given(`^GitHub is running$`, func() {
			gitHub = features.NewFakeGitHub()
			gitHub.Start()
		})

		define.Given(`^GitHub is stopped$`, func() {
			gitHub.Stop()
		})
	})
})
