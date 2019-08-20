package downloadtile

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/Masterminds/semver"

	pivnetClient "github.com/cf-platform-eng/marman/pivnet"
	"github.com/pivotal-cf/go-pivnet"
	"github.com/pkg/errors"
)

type Config struct {
	Slug         string `short:"s" long:"slug" required:"true" description:"Product slug name"`
	File         string `short:"f" long:"file" description:"RegEx pattern to select the specific file to download"`
	Version      string `short:"v" long:"version" default:"X" default-mask:"latest GA" description:"Semver constraint for picking a release version"`
	PivnetClient pivnetClient.Client
	PivnetToken  string `long:"pivnet-token" description:"Authentication token for PivNet" env:"PIVNET_TOKEN"`
}

func filesToString(files []pivnet.ProductFile, filter string) string {
	builder := strings.Builder{}
	for _, file := range files {
		matched, _ := regexp.MatchString(filter, file.AWSObjectKey)
		if matched {
			builder.WriteString(fmt.Sprintf("\n    %s", file.AWSObjectKey))
		}
	}
	return builder.String()

}

func (cmd *Config) FindFile(productFiles []pivnet.ProductFile, id int) (*pivnet.ProductFile, error) {
	var (
		err         error
		productFile pivnet.ProductFile
	)
	productFile.ID = 0
	for _, fileUnderConsideration := range productFiles {
		matched, _ := regexp.MatchString(cmd.File, fileUnderConsideration.AWSObjectKey)
		if matched {
			if productFile.ID == 0 {
				productFile = fileUnderConsideration
			} else {
				err = fmt.Errorf("too many matching files found with the given file filter \"%s\"%s", cmd.File, filesToString(productFiles, cmd.File))
			}
		}
	}

	if productFile.ID == 0 {
		err = errors.Errorf("unable to find the tile with name %s", cmd.File)
	}

	return &productFile, err
}

func (cmd *Config) DownloadTile() error {
	versionConstraint, err := semver.NewConstraint(cmd.Version)
	if err != nil {
		return errors.Wrapf(err, "tile version is not valid semver")
	}

	release, err := cmd.PivnetClient.FindReleaseByVersionConstraint(cmd.Slug, versionConstraint)
	if err != nil {
		return errors.Wrapf(err, "could not list releases for slug %s", cmd.Slug)
	}

	productFiles, err := cmd.PivnetClient.ListFilesForRelease(cmd.Slug, release.ID)
	if err != nil {
		return errors.Wrapf(err, "could not list files for release %d on slug %s", release.ID, cmd.Slug)
	}

	productFile, err := cmd.FindFile(productFiles, release.ID)
	if err != nil {
		return err
	}

	err = cmd.PivnetClient.AcceptEULA(cmd.Slug, release.ID)
	if err != nil {
		return errors.Wrapf(err, "could not accept the eula for slug %s", cmd.Slug)
	}

	err = cmd.PivnetClient.DownloadFile(cmd.Slug, release.ID, productFile)

	return err
}

func (cmd *Config) DownloadFromPivnet(slug, file, version, pivnetToken string) error {
	cmd.Slug = slug
	cmd.File = file
	cmd.Version = version
	cmd.PivnetClient = pivnetClient.NewPivNetClient(pivnetToken)
	return cmd.DownloadTile()
}

func (cmd *Config) Execute(args []string) error {
	cmd.PivnetClient = pivnetClient.NewPivNetClient(cmd.PivnetToken)
	return cmd.DownloadTile()
}
