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

var _ = Describe("Downloading releases from Tanzu Network", func() {
	steps := NewSteps()

	Scenario("download a tanzu network release from multiple files without a filter", func() {
		steps.Given("marman is built")
		steps.And("Tanzu Network is running")
		steps.And("there is a simple release")
		steps.And("there is a compound release")

		steps.When("marman tanzu-network-download -s stemcells-ubuntu-xenial is run")
		steps.And("it exits with error")
	})

	Scenario("download a tanzu network release without a filter and one file", func() {
		steps.Given("marman is built")
		steps.And("Tanzu Network is running")
		steps.And("there is a simple release")

		steps.When("marman tanzu-network-download -s stemcells-ubuntu-xenial is run")
		steps.And("it exits without error")
		steps.And("the 100.000 release is downloaded")
	})

	Scenario("download a tanzu network release with a filter", func() {
		steps.Given("marman is built")
		steps.And("Tanzu Network is running")
		steps.And("there is a simple release")
		steps.And("there is a compound release")

		steps.When("marman tanzu-network-download -s stemcells-ubuntu-xenial -f google is run")
		steps.And("it exits without error")
		steps.And("the latest GA release is downloaded")
	})

	Scenario("download a tanzu network release with a specific version", func() {
		steps.Given("marman is built")
		steps.And("Tanzu Network is running")
		steps.And("there is a simple release")
		steps.And("there is a compound release")

		steps.When("marman tanzu-network-download -s stemcells-ubuntu-xenial -v 100 is run")
		steps.And("it exits without error")
		steps.And("the 100.000 release is downloaded")
	})

	steps.Define(func(define Definitions) {
		var (
			marmanPath     string
			tanzuNetwork   *features.FakeTanzuNetwork
			commandSession *gexec.Session
		)

		define.Given(`^marman is built$`, func() {
			var err error
			marmanPath, err = gexec.Build("github.com/cf-platform-eng/marman/cmd/marman")
			Expect(err).NotTo(HaveOccurred())
		}, func() {
			gexec.CleanupBuildArtifacts()
		})

		define.When("^marman tanzu-network-download -s stemcells-ubuntu-xenial is run$", func() {
			command := exec.Command(marmanPath, "tanzu-network-download", "-s", "stemcells-ubuntu-xenial")
			command.Env = []string{
				fmt.Sprintf("TANZU_NETWORK_HOSTNAME=%s", tanzuNetwork.Host),
			}

			var err error
			commandSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		define.When("^marman tanzu-network-download -s stemcells-ubuntu-xenial -v 100 is run$", func() {
			command := exec.Command(marmanPath, "tanzu-network-download", "-s", "stemcells-ubuntu-xenial", "-v", "100")
			command.Env = []string{
				fmt.Sprintf("TANZU_NETWORK_HOSTNAME=%s", tanzuNetwork.Host),
			}

			var err error
			commandSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		define.When("^marman tanzu-network-download -s stemcells-ubuntu-xenial -f google is run$", func() {
			command := exec.Command(marmanPath, "tanzu-network-download", "-s", "stemcells-ubuntu-xenial", "-f", "google")
			command.Env = []string{
				fmt.Sprintf("TANZU_NETWORK_HOSTNAME=%s", tanzuNetwork.Host),
			}

			var err error
			commandSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
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

		define.Then(`^the latest GA release is downloaded$`, func() {
			steps.Run("the 123.456 release is downloaded")
		})

		define.Then(`^it exits without error$`, func() {
			Eventually(commandSession).Should(gexec.Exit(0))
		})

		define.Then(`^it exits with error$`, func() {
			Eventually(commandSession).Should(gexec.Exit())
			Expect(commandSession.ExitCode()).Should(BeNumerically(">", 0))
		})

		define.Given(`^Tanzu Network is running$`, func() {
			tanzuNetwork = features.NewFakeTanzuNetwork()
			tanzuNetwork.Start()
		}, func() {
			tanzuNetwork.Stop()
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
	})
})
