package marman

import (
	"github.com/cf-platform-eng/marman/downloadtile"
	pivnetClient "github.com/cf-platform-eng/marman/pivnet"
)

type DownloadSRTConfig struct {
	Version     string `short:"v" long:"version" default:"X" default-mask:"latest GA" description:"Semver constraint for picking a release version"`
	PivnetToken string `long:"pivnet-token" description:"Authentication token for PivNet" env:"PIVNET_TOKEN"`
}

func (cmd *DownloadSRTConfig) Execute(args []string) error {
	downloadTileCommand := downloadtile.Config{
		Slug:         "cf",
		File:         "srt-(.*)-(.*).pivotal$",
		Version:      cmd.Version,
		PivnetClient: pivnetClient.NewPivNetClient(cmd.PivnetToken),
	}
	return downloadTileCommand.DownloadTile()
}
