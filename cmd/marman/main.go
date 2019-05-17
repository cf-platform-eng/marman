package main

import (
	"fmt"
	"github.com/cf-platform-eng/marman/downloadrelease"
	"os"

	"code.cloudfoundry.org/lager"

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
		fmt.Println("Could not add download-release command")
		os.Exit(1)
	}

	downloadStemcellOpts := &downloadstemcell.Config{
		Logger: lager.NewLogger("download-stemcell"),
	}
	_, err = parser.AddCommand(
		"download-stemcell",
		"Download stemcell",
		"Download stemcell from PivNet",
		downloadStemcellOpts,
	)
	if err != nil {
		fmt.Println("Could not add download-stemcell command")
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
		fmt.Println("Could not add download-tile command")
		os.Exit(1)
	}

	_, err = parser.AddCommand(
		"version",
	    "print version",
        "print marman version",
        &version.VersionOpt{})
	if err != nil {
		fmt.Println("Could not add version command")
		os.Exit(1)
	}

	_, err = parser.Parse()
	if err != nil {
		// TODO: look into printing a usage on bad commands
		os.Exit(1)
	}
}
