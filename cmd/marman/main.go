package main

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/lager"
	"github.com/cf-platform-eng/marman/downloadrelease"

	"github.com/cf-platform-eng/marman"

	"github.com/cf-platform-eng/marman/downloadstemcell"
	"github.com/cf-platform-eng/marman/downloadtile"
	"github.com/cf-platform-eng/marman/version"

	"github.com/jessevdk/go-flags"
)

var config marman.Config
var parser = flags.NewParser(&config, flags.Default)

func main() {
	downloadReleaseOpts := &downloadrelease.Config{
		Logger: lager.NewLogger("download-release"),
	}
	_, err := parser.AddCommand(
		"download-release",
		"Download release",
		"Download release from GitHub",
		downloadReleaseOpts,
	)
	if err != nil {
		fmt.Printf("Could not add download-release command: %s\n", err.Error())
		os.Exit(1)
	}

	downloadStemcellOpts := &downloadstemcell.Config{
		Downloader: &downloadtile.Config{},
	}
	_, err = parser.AddCommand(
		"download-stemcell",
		"Download stemcell",
		"Download stemcell from PivNet",
		downloadStemcellOpts,
	)
	if err != nil {
		fmt.Printf("Could not add download-stemcell command: %s\n", err.Error())
		os.Exit(1)
	}

	downloadPKSOpts := &downloadtile.Config{
		Slug: "pivotal-container-service",
		File: ".pivotal$",
	}
	_, err = parser.AddCommand(
		"download-pks",
		"Download pks",
		"Download pks tile from PivNet",
		downloadPKSOpts,
	)
	if err != nil {
		fmt.Printf("Could not add download-pks command: %s\n", err.Error())
		os.Exit(1)
	}

	downloadSRTOpts := &marman.DownloadSRTConfig{}
	_, err = parser.AddCommand(
		"download-srt",
		"Download SRT",
		"Download SRT tile from PivNet",
		downloadSRTOpts,
	)
	if err != nil {
		fmt.Printf("Could not add download-srt command: %s\n", err.Error())
		os.Exit(1)
	}

	downloadPASOpts := &marman.DownloadPASConfig{}
	_, err = parser.AddCommand(
		"download-pas",
		"Download PAS",
		"Download PAS tile from PivNet",
		downloadPASOpts,
	)
	if err != nil {
		fmt.Printf("Could not add download-pas command: %s\n", err.Error())
		os.Exit(1)
	}

	downloadTileOpts := &downloadtile.Config{}
	_, err = parser.AddCommand(
		"download-tile",
		"Download tile",
		"Download tile from PivNet",
		downloadTileOpts,
	)
	if err != nil {
		fmt.Printf("Could not add download-tile command: %s\n", err.Error())
		os.Exit(1)
	}

	_, err = parser.AddCommand(
		"version",
		"print version",
		"print marman version",
		&version.VersionOpt{})
	if err != nil {
		fmt.Printf("Could not add version command: %s\n", err.Error())
		os.Exit(1)
	}

	_, err = parser.Parse()
	if err != nil {
		// TODO: look into printing a usage on bad commands
		os.Exit(1)
	}
}
