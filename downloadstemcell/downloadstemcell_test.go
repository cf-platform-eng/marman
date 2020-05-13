package downloadstemcell_test

import (
	"errors"

	"github.com/cf-platform-eng/marman/downloadstemcell"
	"github.com/cf-platform-eng/marman/downloadstemcell/downloadstemcellfakes"
	"github.com/cf-platform-eng/marman/downloadtile"
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
		cmd.Version = "123"
		cmd.IAAS = "google"
		cmd.PivnetToken = "secret-token"
		cmd.TanzuNetHost = "https://network.tanzu.vmware.com"

		err := cmd.Execute([]string{})
		Expect(err).ToNot(HaveOccurred())
		Expect(downloader.DownloadFromPivnetCallCount()).To(Equal(1))
		slug, file, version, tanzuNetHost, pivnetToken := downloader.DownloadFromPivnetArgsForCall(0)
		Expect(slug).To(Equal("stemcells-ubuntu-xenial"))
		Expect(file).To(Equal("bosh-stemcell-123[\\d.]*-google-.*\\.tgz$"))
		Expect(version).To(Equal("123"))
		Expect(pivnetToken).To(Equal("secret-token"))
		Expect(tanzuNetHost).To(Equal("https://network.tanzu.vmware.com"))
	})

	Context("results include both a lite and a heavy stemcell", func() {
		BeforeEach(func() {
			cmd.OS = "ubuntu-xenial"
			cmd.Version = "123.4"
			cmd.IAAS = "google"
			cmd.PivnetToken = "secret-token"
		})

		Context("has a matching light stemcell", func() {
			It("retries with light stemcells", func() {
				downloader.DownloadFromPivnetReturnsOnCall(0, downloadtile.TooManyFilesError{})
				downloader.DownloadFromPivnetReturnsOnCall(1, nil)

				By("not returning an error", func() {
					err := cmd.Execute([]string{})
					Expect(err).NotTo(HaveOccurred())
				})

				By("trying with a filter that matches too many files", func() {
					_, file, _, _, _ := downloader.DownloadFromPivnetArgsForCall(0)
					Expect(file).To(Equal("bosh-stemcell-123.4[\\d.]*-google-.*\\.tgz$"))
				})

				By("trying the light version", func() {
					_, file, _, _, _ := downloader.DownloadFromPivnetArgsForCall(1)
					Expect(file).To(Equal("light-bosh-stemcell-123.4[\\d.]*-google-.*\\.tgz$"))
				})
			})
		})

		Context("has no matching light stemcell", func() {
			It("returns an error", func() {
				exitError := downloadtile.NoMatchError{Filter: "google"}
				downloader.DownloadFromPivnetReturnsOnCall(0, downloadtile.TooManyFilesError{})
				downloader.DownloadFromPivnetReturnsOnCall(1, exitError)

				By("returning a slug not matched error", func() {
					err := cmd.Execute([]string{})
					Expect(err).To(HaveOccurred())
					Expect(err).To(Equal(exitError))
				})

				By("trying with a filter that matches too many files", func() {
					_, file, _, _, _ := downloader.DownloadFromPivnetArgsForCall(0)
					Expect(file).To(Equal("bosh-stemcell-123.4[\\d.]*-google-.*\\.tgz$"))
				})

				By("trying the light version", func() {
					_, file, _, _, _ := downloader.DownloadFromPivnetArgsForCall(1)
					Expect(file).To(Equal("light-bosh-stemcell-123.4[\\d.]*-google-.*\\.tgz$"))
				})
			})

		})
	})

	Context("download from pivnet returns an error", func() {
		BeforeEach(func() {
			cmd.OS = "ubuntu-xenial"
			cmd.Version = "123.4"
			cmd.IAAS = "google"
			cmd.PivnetToken = "secret-token"
		})

		It("returns an error", func() {
			downloadFromPivnetError := errors.New("some error")
			downloader.DownloadFromPivnetReturnsOnCall(0, downloadFromPivnetError)

			By("returning a slug not matched error", func() {
				err := cmd.Execute([]string{})
				Expect(err).To(HaveOccurred())
				Expect(err).To(Equal(downloadFromPivnetError))
			})
		})
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
