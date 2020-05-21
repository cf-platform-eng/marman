package pivnet

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDownloadtile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Pivnet Suite")
}
