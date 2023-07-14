# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/tsm-pass/issues) for any
deviations that you spot; I'm still learning!.

## Types of changes

The following types of changes will be recorded in this file:

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- placeholder

## [v0.2.1] - 2023-07-14

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.19.11
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.8` to `1.19.11`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.10.4` to `go-ci-oldstable-build-v0.11.4`
- (GH-136) Update vuln analysis GHAW to remove on.push hook

### Fixed

- (GH-132) Disable depguard linter
- (GH-136) Restore local CodeQL workflow

## [v0.2.0] - 2023-04-10

### Overview

- Add support for generating DEB, RPM packages
- Build improvements
- Generated binary changes
  - filename patterns
  - compression (~ 66% smaller)
  - executable metadata
- built using Go 1.19.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-118) Generate RPM/DEB packages using nFPM
- (GH-119) Add version details to Windows executables

### Changed

- (GH-117) Switch to semantic versioning (semver) compatible versioning
  pattern
- (GH-120) Makefile: Compress binaries & use fixed filenames
- (GH-121) Makefile: Refresh recipes to add "standard" set, new
  package-related options
- (GH-122) Build dev/stable releases using go-ci Docker image

## [v0.1.13] - 2023-04-10

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions workflow updates
- built using Go 1.19.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- (GH-111) Add Go Module Validation, Dependency Updates jobs

### Changed

- Dependencies
  - `Go`
    - `1.19.4` to `1.19.8`
- (GH-113) Drop `Push Validation` workflow
- (GH-114) Rework workflow scheduling
- (GH-116) Remove `Push Validation` workflow status badge

### Fixed

- (GH-124) Update vuln analysis GHAW to use on.push hook

## [v0.1.12] - 2022-12-12

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.4
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.1` to `1.19.4`
- (GH-103) Refactor GitHub Actions workflows to import logic

### Fixed

- (GH-106) Fix Makefile Go module base path detection

## [v0.1.11] - 2022-09-22

### Overview

- Dependency updates
- GitHub Actions Workflows updates
- built using Go 1.19.1
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.13` to `1.19.1`
  - `github/codeql-action`
    - `v2.1.22` to `v2.1.24`
- (GH-94) Update project to Go 1.19
- (GH-95) Update Makefile and GitHub Actions Workflows
- (GH-96) Add CodeQL GitHub Actions Workflow

## [v0.1.10] - 2022-09-01

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.17.13
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.9` to `1.17.13`

### Fixed

- (GH-90) Update lintinstall Makefile recipe
- (GH-93) Fix/workaround linting errors

## [v0.1.9] - 2022-05-06

### Overview

- Dependency updates
- built using Go 1.17.9
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.7` to `1.17.9`

## [v0.1.8] - 2022-03-02

### Overview

- Dependency updates
- CI / linting improvements
- built using Go 1.17.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.17.6` to `1.17.7`
  - `actions/checkout`
    - `v2.4.0` to `v3`
  - `actions/setup-node`
    - `v2.5.1` to `v3`

- (GH-74) Expand linting GitHub Actions Workflow to include `oldstable`,
  `unstable` container images
- (GH-75) Switch Docker image source from Docker Hub to GitHub Container
  Registry (GHCR)

### Fixed

- (GH-77) var-declaration: should omit type string from declaration of var
  (revive)

## [v0.1.7] - 2022-01-25

### Overview

- Dependency updates
- built using Go 1.17.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.12` to `1.17.6`
    - (GH-70) Update go.mod file, canary Dockerfile to reflect current
      dependencies

## [v0.1.6] - 2021-12-29

### Overview

- Dependency updates
- built using Go 1.16.12
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.10` to `1.16.12`
  - `actions/setup-node`
    - `v2.4.1` to `v2.5.1`

## [v0.1.5] - 2021-11-10

### Overview

- Dependency updates
- built using Go 1.16.10
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.8` to `1.16.10`
  - `actions/checkout`
    - `v2.3.4` to `v2.4.0`
  - `actions/setup-node`
    - `v2.4.0` to `v2.4.1`

## [v0.1.4] - 2021-09-27

### Overview

- Dependency updates
- built using Go 1.16.8
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.7` to `1.16.8`

## [v0.1.3] - 2021-08-08

### Overview

- Dependency updates
- built using Go 1.16.7
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.16.6` to `1.16.7`
  - `actions/setup-node`
    - updated from `v2.2.0` to `v2.4.0`

## [v0.1.2] - 2021-07-19

### Overview

- Built using Go 1.16.6
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Added

- Add "canary" Dockerfile to track stable Go releases, serve as a reminder to
  generate fresh binaries

### Changed

- Swap out GoDoc badge for pkg.go.dev badge

- Dependencies
  - `Go`
    - `1.15.5` to `1.16.6`
  - `actions/setup-node`
    - `v2.1.2` to `v2.2.0`
    - update `node-version` value to always use latest LTS version instead of
      hard-coded version

### Fixed

- cmd/tsm-pass/main.go:37:19: appendAssign: append result not assigned to the
  same slice (gocritic)

## [v0.1.1] - 2020-11-18

### Overview

- Built using Go 1.15.5
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

- Update password restrictions, minor cleanup.

### Changed

- Extend list of supported characters to reflect IBM Spectrum Protect v6.3.3+
- Update config initialization
- Minor code cleanup

### Fixed

- GoDoc: Missing line breaks between features list entries

## [v0.1.0] - 2020-11-10

### Added

Initial release!

This release provides an early version of a Go CLI application that can be
used to generate a password string compatible with Tivoli/TSM/Spectrum
Protect. An older Python script previously used for this purpose is included
for reference or alternative use.

- Statically linked binary release
  - Built using Go 1.15.3
  - Windows
    - x86
    - x64
  - Linux
    - x86
    - x64

[Unreleased]: https://github.com/atc0005/tsm-pass/compare/v0.2.1...HEAD
[v0.2.1]: https://github.com/atc0005/tsm-pass/releases/tag/v0.2.1
[v0.2.0]: https://github.com/atc0005/tsm-pass/releases/tag/v0.2.0
[v0.1.13]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.13
[v0.1.12]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.12
[v0.1.11]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.11
[v0.1.10]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.10
[v0.1.9]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.9
[v0.1.8]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.8
[v0.1.7]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.7
[v0.1.6]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.6
[v0.1.5]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.5
[v0.1.4]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.4
[v0.1.3]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.3
[v0.1.2]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.0
