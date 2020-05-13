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
	"github.com/pivotal-cf/go-pivnet"
)

var _ = Describe("Downloading stemcells", func() {
	steps := NewSteps()

	Scenario("downloading stemcells", func() {
		steps.Given("marman is built")
		steps.And("Tanzu Network is running")
		steps.And("there is a stemcell")

		steps.When("marman download-stemcell -o ubuntu-xenial -i google -v 123 is run")

		steps.Then("it exits without error")
		steps.And("the stemcell is downloaded")
	})

	Scenario("download the light stemcell, if there is an option", func() {
		steps.Given("marman is built")
		steps.And("Tanzu Network is running")
		steps.And("there is both a light and a heavy stemcell")

		steps.When("marman download-stemcell -o ubuntu-xenial -i google -v 123 is run")

		steps.Then("it exits without error")
		steps.And("the light stemcell is downloaded")
	})

	Scenario("download latest GA version", func() {
		steps.Given("marman is built")
		steps.And("Tanzu Network is running")
		steps.And("there are multiple stemcells")

		steps.When("marman download-stemcell -o ubuntu-xenial -i google is run")

		steps.Then("it exits without error")
		steps.And("the latest GA stemcell is downloaded")
	})

	steps.Define(func(define Definitions) {
		var (
			expectedStemcellFilePath string
			marmanPath               string
			tanzuNetwork             *features.FakeTanzuNetwork
			commandSession           *gexec.Session
		)

		define.Given(`^marman is built$`, func() {
			var err error
			marmanPath, err = gexec.Build("github.com/cf-platform-eng/marman/cmd/marman")
			Expect(err).NotTo(HaveOccurred())
		}, func() {
			gexec.CleanupBuildArtifacts()
		})

		define.Given(`^Tanzu Network is running$`, func() {
			tanzuNetwork = features.NewFakeTanzuNetwork()
			tanzuNetwork.Start()
		}, func() {
			tanzuNetwork.Stop()
		})

		define.Given(`^there is a stemcell$`, func() {
			tanzuNetwork.Releases = append(tanzuNetwork.Releases, pivnet.Release{
				ID:      12345,
				Version: "123",
			})
			tanzuNetwork.ProductFiles = append(tanzuNetwork.ProductFiles, pivnet.ProductFile{
				ID:           6789,
				AWSObjectKey: "/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
				Links: &pivnet.Links{
					Download: map[string]string{
						"href": "/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
					},
				},
			})
		})

		define.Given(`^there are multiple stemcells$`, func() {
			tanzuNetwork.Releases = append(tanzuNetwork.Releases, pivnet.Release{
				ID:      11111,
				Version: "100",
			})
			tanzuNetwork.Releases = append(tanzuNetwork.Releases, pivnet.Release{
				ID:      12345,
				Version: "123",
			})
			tanzuNetwork.ProductFiles = append(tanzuNetwork.ProductFiles, pivnet.ProductFile{
				ID:           6789,
				AWSObjectKey: "/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
				Links: &pivnet.Links{
					Download: map[string]string{
						"href": "/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
					},
				},
			})
		})

		define.Given(`^there is both a light and a heavy stemcell$`, func() {
			tanzuNetwork.Releases = append(tanzuNetwork.Releases, pivnet.Release{
				ID:      12345,
				Version: "123",
			})
			tanzuNetwork.ProductFiles = append(tanzuNetwork.ProductFiles, pivnet.ProductFile{
				ID:           6789,
				AWSObjectKey: "/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
				Links: &pivnet.Links{
					Download: map[string]string{
						"href": "/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
					},
				},
			})
			tanzuNetwork.ProductFiles = append(tanzuNetwork.ProductFiles, pivnet.ProductFile{
				ID:           7890,
				AWSObjectKey: "/path/to/light-bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
				Links: &pivnet.Links{
					Download: map[string]string{
						"href": "/path/to/light-bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
					},
				},
			})
		})

		define.When(`^marman download-stemcell -o ubuntu-xenial -i google is run$`, func() {
			command := exec.Command(marmanPath, "download-stemcell", "-o", "ubuntu-xenial", "-i", "google")
			command.Env = []string{
				fmt.Sprintf("TANZU_NETWORK_HOSTNAME=%s", tanzuNetwork.Host),
			}

			var err error
			commandSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		define.When(`^marman download-stemcell -o ubuntu-xenial -i google -v 123 is run$`, func() {
			command := exec.Command(marmanPath, "download-stemcell", "-o", "ubuntu-xenial", "-i", "google", "-v", "123")
			command.Env = []string{
				fmt.Sprintf("TANZU_NETWORK_HOSTNAME=%s", tanzuNetwork.Host),
			}

			var err error
			commandSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		define.Then(`^it exits without error$`, func() {
			Eventually(commandSession).Should(gexec.Exit(0))
		})

		define.Then(`^the stemcell is downloaded$`, func() {
			if expectedStemcellFilePath == "" {
				cwd, err := os.Getwd()
				Expect(err).NotTo(HaveOccurred())

				expectedStemcellFilePath = path.Join(cwd, "bosh-stemcell-123.456-google-super-duper-stemcell.tgz")
			}

			info, err := os.Stat(expectedStemcellFilePath)
			Expect(err).NotTo(HaveOccurred())
			Expect(info.IsDir()).To(BeFalse())
		}, func() {
			err := os.Remove(expectedStemcellFilePath)
			Expect(err).NotTo(HaveOccurred())
			expectedStemcellFilePath = ""
		})

		define.Then(`^the light stemcell is downloaded$`, func() {
			cwd, err := os.Getwd()
			Expect(err).NotTo(HaveOccurred())

			expectedStemcellFilePath = path.Join(cwd, "light-bosh-stemcell-123.456-google-super-duper-stemcell.tgz")
			steps.Run("the stemcell is downloaded")
		})

		define.Then(`^the latest GA stemcell is downloaded$`, func() {
			steps.Run("the stemcell is downloaded")
		})
	})
})
