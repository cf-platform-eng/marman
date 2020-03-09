package downloadstemcell

import (
	"errors"
	"fmt"

	"github.com/cf-platform-eng/marman/downloadtile"
)

type Config struct {
	OS          string `short:"o" long:"os" required:"true" description:"Stemcell OS name" choice:"ubuntu-trusty" choice:"ubuntu-xenial"`
	Version     string `short:"v" long:"version" default:"X" default-mask:"latest GA" description:"Stemcell version"`
	IAAS        string `short:"i" long:"iaas" required:"true" description:"Specific stemcell IaaS to download"`
	PivnetToken string `long:"pivnet-token" description:"Authentication token for PivNet" env:"PIVNET_TOKEN"`
	Downloader  Downloader
}

//go:generate counterfeiter Downloader
type Downloader interface {
	DownloadFromPivnet(slug, file, version, pivnetToken string) error
}

func stemcellOSToSlug(os string) (string, error) {
	if os == "" {
		return "", errors.New("missing stemcell os")
	}
	switch os {
	case "ubuntu-trusty":
		return "stemcells-ubuntu", nil
	case "ubuntu-xenial":
		return "stemcells-ubuntu-xenial", nil
	}
	return "", fmt.Errorf("invalid stemcell os: %s", os)
}

func stemcellFileFilter(version, iaas string, light bool) string {
	if light {
		return fmt.Sprintf("light-bosh-stemcell-%s-%s-.*\\.tgz$", version, iaas)
	}
	return fmt.Sprintf("bosh-stemcell-%s-%s-.*\\.tgz$", version, iaas)
}

func (cmd *Config) Execute(args []string) error {
	slug, err := stemcellOSToSlug(cmd.OS)
	if err != nil {
		return err
	}

	err = cmd.Downloader.DownloadFromPivnet(slug, stemcellFileFilter(cmd.Version, cmd.IAAS, false), cmd.Version, cmd.PivnetToken)
	if errors.As(err, &downloadtile.TooManyFilesError{}) {
		err = cmd.Downloader.DownloadFromPivnet(slug, stemcellFileFilter(cmd.Version, cmd.IAAS, true), cmd.Version, cmd.PivnetToken)
	}
	return err
}
