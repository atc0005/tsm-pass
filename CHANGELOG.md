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

[Unreleased]: https://github.com/atc0005/tsm-pass/compare/v0.1.4...HEAD
[v0.1.4]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.4
[v0.1.3]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.3
[v0.1.2]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/tsm-pass/releases/tag/v0.1.0
