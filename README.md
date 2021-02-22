# MarMan

MarMan (Marketplace Manager), is a tool used by ISV-CI to download files from various marketplaces.

Currently, it supports downloading files from the [Tanzu Network](https://network.pivotal.io), or from releases on [Github](https://github.com).

## Commands

* `marman github-download-release` - Downloads a file from a GitHub release. May require a GITHUB_TOKEN.
* `marman tanzu-network-download` - Downloads a file off of Tanzu Network. Requires a PIVNET_TOKEN (refresh or legacy).
