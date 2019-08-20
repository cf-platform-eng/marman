package downloadstemcell

import (
	"errors"
	"fmt"
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

func (cmd *Config) Execute(args []string) error {
	slug, err := stemcellOSToSlug(cmd.OS)
	if err != nil {
		return err
	}
	return cmd.Downloader.DownloadFromPivnet(slug, cmd.IAAS, cmd.Version, cmd.PivnetToken)
}
