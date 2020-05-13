// +build feature

package features_test

import (
	"os/exec"

	. "github.com/bunniesandbeatings/goerkin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

var _ = Describe("Marman", func() {
	steps := NewSteps()

	Scenario("version command reports version", func() {
		steps.Given("marman is built with a version")

		steps.When("marman version is run")

		steps.Then("it exits without error")
		steps.And("it prints the version")
	})

	steps.Define(func(define Definitions) {
		var (
			marmanPath     string
			commandSession *gexec.Session
		)

		define.Given(`^marman is built with a version$`, func() {
			var err error
			marmanPath, err = gexec.Build(
				"github.com/cf-platform-eng/marman/cmd/marman",
				"-ldflags",
				"-X github.com/cf-platform-eng/marman/version.Version=1.2.3",
			)
			Expect(err).NotTo(HaveOccurred())
		}, func() {
			gexec.CleanupBuildArtifacts()
		})

		define.When(`^marman version is run$`, func() {
			versionCommand := exec.Command(marmanPath, "version")
			var err error
			commandSession, err = gexec.Start(versionCommand, GinkgoWriter, GinkgoWriter)
			Expect(err).NotTo(HaveOccurred())
		})

		define.Then(`^it exits without error$`, func() {
			Eventually(commandSession).Should(gexec.Exit(0))
		})

		define.Then(`it prints the version`, func() {
			Eventually(commandSession.Out).Should(Say("marman version: 1.2.3"))
		})
	})
})
