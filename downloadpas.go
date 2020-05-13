package marman

import (
	"github.com/cf-platform-eng/marman/downloadtile"
	pivnetClient "github.com/cf-platform-eng/marman/pivnet"
)

type DownloadPASConfig struct {
	Version      string `short:"v" long:"version" default:"X" default-mask:"latest GA" description:"Semver constraint for picking a release version"`
	TanzuNetHost string `long:"pivnet-host" description:"Host for Tanzu Network" env:"TANZU_NETWORK_HOSTNAME" default:"https://network.pivotal.io"`
	PivnetToken  string `long:"pivnet-token" description:"Authentication token for PivNet" env:"PIVNET_TOKEN"`
}

func (cmd *DownloadPASConfig) Execute(args []string) error {
	downloadTileCommand := downloadtile.Config{
		Slug:         "cf",
		File:         "cf-(.*)-(.*).pivotal$",
		Version:      cmd.Version,
		PivnetClient: pivnetClient.NewPivNetClient(cmd.TanzuNetHost, cmd.PivnetToken),
	}
	return downloadTileCommand.DownloadTile()
}
