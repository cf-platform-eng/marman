# MarMan

MarMan (Marketplace Manager), is a tool used by ISV-CI to download files from various marketplaces.

Currently, it supports downloading files from the [Pivotal Network](https://network.pivotal.io), or from releases on [Github](https://github.com).

## Commands

* `marman download-tile` - Downloads a file off of pivnet. Requires a PIVNET_TOKEN (refresh or legacy).
* `marman download-release` - Downloads a file from a GitHub release. May require a GITHUB_TOKEN.
