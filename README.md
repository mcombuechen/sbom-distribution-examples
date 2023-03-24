# SBOM Distribution Examples

A test repository to play around with SBOM ditribution.

## Features of GitHub Action

* Run on every GitHub release
* generate SBOM for many ecosystems (npm, pypi, maven, ...)
* generate SBOM with many generators (snyk, syft, cdxgen, ...)
* generate SBOM in many formats (SPDX, CycloneDX)
* Post documents to Release Artifacts

## TODOs

* get hold of contextual information about release (ID, version, ...)
* move logic into `action` directory
* post SHA256 of SBOM as artifact
