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
		steps.And("there is a simple release")

		steps.When("marman download-stemcell -o ubuntu-xenial -i google -v 100 is run")

		steps.Then("it exits without error")
		steps.And("the 100.000 release is downloaded")
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

			stemcell := pivnet.ProductFile{
				ID:           6789,
				AWSObjectKey: "/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
				Links: &pivnet.Links{
					Download: map[string]string{
						"href": "/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
					},
				},
			}

			if tanzuNetwork.ReleaseFiles == nil {
				tanzuNetwork.ReleaseFiles = map[string][]pivnet.ProductFile{}
			}

			tanzuNetwork.ReleaseFiles["12345"] = []pivnet.ProductFile{stemcell}
		})

		define.Given(`^there is both a light and a heavy stemcell$`, func() {
			tanzuNetwork.Releases = append(tanzuNetwork.Releases, pivnet.Release{
				ID:      12345,
				Version: "123",
			})

			heavy123 := pivnet.ProductFile{
				ID:           6789,
				AWSObjectKey: "/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
				Links: &pivnet.Links{
					Download: map[string]string{
						"href": "/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
					},
				},
			}

			light123 := pivnet.ProductFile{
				ID:           7890,
				AWSObjectKey: "/path/to/light-bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
				Links: &pivnet.Links{
					Download: map[string]string{
						"href": "/path/to/light-bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
					},
				},
			}

			if tanzuNetwork.ReleaseFiles == nil {
				tanzuNetwork.ReleaseFiles = map[string][]pivnet.ProductFile{}
			}

			tanzuNetwork.ReleaseFiles["12345"] = []pivnet.ProductFile{heavy123, light123}
		})

		define.Given(`^there is a simple release$`, func() {
			tanzuNetwork.Releases = append(tanzuNetwork.Releases, pivnet.Release{
				ID:      11111,
				Version: "100",
			})

			google100 := pivnet.ProductFile{
				ID:           6789,
				AWSObjectKey: "/path/to/bosh-stemcell-100.000-google-super-duper-stemcell.tgz",
				Links: &pivnet.Links{
					Download: map[string]string{
						"href": "/path/to/bosh-stemcell-100.000-google-super-duper-stemcell.tgz",
					},
				},
			}

			if tanzuNetwork.ReleaseFiles == nil {
				tanzuNetwork.ReleaseFiles = map[string][]pivnet.ProductFile{}
			}

			tanzuNetwork.ReleaseFiles["11111"] = []pivnet.ProductFile{google100}
		})

		define.Given(`^there is a compound release$`, func() {
			tanzuNetwork.Releases = append(tanzuNetwork.Releases, pivnet.Release{
				ID:      12345,
				Version: "123",
			})

			vsphere123 := pivnet.ProductFile{
				ID:           6790,
				AWSObjectKey: "/path/to/bosh-stemcell-123.456-vsphere-super-duper-stemcell.tgz",
				Links: &pivnet.Links{
					Download: map[string]string{
						"href": "/path/to/bosh-stemcell-123.456-vsphere-super-duper-stemcell.tgz",
					},
				},
			}

			google123 := pivnet.ProductFile{
				ID:           6791,
				AWSObjectKey: "/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
				Links: &pivnet.Links{
					Download: map[string]string{
						"href": "/path/to/bosh-stemcell-123.456-google-super-duper-stemcell.tgz",
					},
				},
			}

			if tanzuNetwork.ReleaseFiles == nil {
				tanzuNetwork.ReleaseFiles = map[string][]pivnet.ProductFile{}
			}

			tanzuNetwork.ReleaseFiles["12345"] = []pivnet.ProductFile{vsphere123, google123}
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

		define.When(`^marman download-stemcell -o ubuntu-xenial -i google -v ([0-9\.]+) is run$`, func(version string) {
			command := exec.Command(marmanPath, "download-stemcell", "-o", "ubuntu-xenial", "-i", "google", "-v", version)
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

			err = os.Remove(expectedStemcellFilePath)
			Expect(err).NotTo(HaveOccurred())
			expectedStemcellFilePath = ""
		})

		define.Then(`^the ([0-9\.]+) release is downloaded$`, func(version string) {
			cwd, err := os.Getwd()
			Expect(err).NotTo(HaveOccurred())
			expectedReleasePath := path.Join(cwd, fmt.Sprintf("bosh-stemcell-%s-google-super-duper-stemcell.tgz", version))

			info, err := os.Stat(expectedReleasePath)
			Expect(err).NotTo(HaveOccurred())
			Expect(info.IsDir()).To(BeFalse())

			err = os.Remove(expectedReleasePath)
			Expect(err).NotTo(HaveOccurred())
			expectedReleasePath = ""
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
