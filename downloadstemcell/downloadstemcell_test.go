package downloadstemcell_test

import (
	"github.com/cf-platform-eng/marman/downloadstemcell"
	"github.com/cf-platform-eng/marman/downloadstemcell/downloadstemcellfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Download Stemcell", func() {
	var (
		downloader *downloadstemcellfakes.FakeDownloader
		cmd        *downloadstemcell.Config
	)

	BeforeEach(func() {
		downloader = &downloadstemcellfakes.FakeDownloader{}
		cmd = &downloadstemcell.Config{
			Downloader: downloader,
		}
	})

	It("downloads the file", func() {
		cmd.OS = "ubuntu-xenial"
		cmd.Version = "123.4"
		cmd.IAAS = "google"
		cmd.PivnetToken = "secret-token"

		err := cmd.Execute([]string{})
		Expect(err).ToNot(HaveOccurred())
		Expect(downloader.DownloadFromPivnetCallCount()).To(Equal(1))
		slug, file, version, pivnetToken := downloader.DownloadFromPivnetArgsForCall(0)
		Expect(slug).To(Equal("stemcells-ubuntu-xenial"))
		Expect(file).To(Equal("google"))
		Expect(version).To(Equal("123.4"))
		Expect(pivnetToken).To(Equal("secret-token"))
	})

	Context("Missing stemcell OS argument", func() {
		BeforeEach(func() {
			cmd.OS = ""
		})

		It("returns an error", func() {
			err := cmd.Execute([]string{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("missing stemcell os"))
		})
	})

	Context("Invalid stemcell OS argument", func() {
		BeforeEach(func() {
			cmd.OS = "ibm-os/2"
		})

		It("returns an error", func() {
			err := cmd.Execute([]string{})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("invalid stemcell os: ibm-os/2"))
		})
	})
})
