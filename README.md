<!-- omit in toc -->
# tsm-pass

Generate passwords compatible with Tivoli/TSM/Spectrum Protect

[![Latest Release](https://img.shields.io/github/release/atc0005/tsm-pass.svg?style=flat-square)](https://github.com/atc0005/tsm-pass/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/atc0005/tsm-pass.svg)](https://pkg.go.dev/github.com/atc0005/tsm-pass)
[![go.mod Go version](https://img.shields.io/github/go-mod/go-version/atc0005/tsm-pass)](https://github.com/atc0005/tsm-pass)
[![Lint and Build](https://github.com/atc0005/tsm-pass/actions/workflows/lint-and-build.yml/badge.svg)](https://github.com/atc0005/tsm-pass/actions/workflows/lint-and-build.yml)
[![Project Analysis](https://github.com/atc0005/tsm-pass/actions/workflows/project-analysis.yml/badge.svg)](https://github.com/atc0005/tsm-pass/actions/workflows/project-analysis.yml)

<!-- omit in toc -->
## Table of Contents

- [Project home](#project-home)
- [Overview](#overview)
- [Password requirements](#password-requirements)
  - [Length](#length)
  - [case-sensitive (or not)](#case-sensitive-or-not)
  - [Valid characters](#valid-characters)
- [Features](#features)
- [Changelog](#changelog)
- [Requirements](#requirements)
  - [Building source code](#building-source-code)
  - [Running](#running)
- [Installation](#installation)
  - [From source](#from-source)
  - [Using release binaries](#using-release-binaries)
- [Configuration](#configuration)
  - [Command-line Arguments](#command-line-arguments)
- [Examples](#examples)
  - [Help output](#help-output)
  - [Default settings](#default-settings)
  - [Default length, custom minimum digits, special characters](#default-length-custom-minimum-digits-special-characters)
  - [Exceeding max length](#exceeding-max-length)
- [License](#license)
- [References](#references)

## Project home

See [our GitHub repo][repo-url] for the latest code, to file an issue or
submit improvements for review and potential inclusion into the project.

## Overview

This repo provides a Go CLI application that can be used to generate a
password string compatible with IBM Spectrum Protect. An older Python
script previously used for this purpose is included for reference or
alternative use.

Historically, IBM Spectrum Protect was known as `Tivoli` or `TSM`. Because the
current maintainer has used the product for years by either of `Tivoli` or
`TSM`, those names are still heavily referenced both here and in internal
documentation.

## Password requirements

The following requirements apply to IBM Spectrum Protect server versions
between 6.3.3 and 8.1.10 (latest version at the time of this writing).

### Length

Tivoli passwords can be up to 63 characters in length.

### case-sensitive (or not)

- if *not* using a LDAP directory server
  - passwords are stored in the IBM Spectrum Protect server database
  - passwords are *not* case-sensitive
- if using a LDAP directory server
  - passwords are stored in the LDAP directory server
    - or whatever underlying system that the LDAP directory server uses
  - passwords are case-sensitive
  - passwords are subject to more restrictions that can be imposed by LDAP
    policies

### Valid characters

Passwords can be be composed of these valid characters:

| Characters | Explanation                                            |
| ---------- | ------------------------------------------------------ |
| `a-z`      | Any letter, a through z, lower-case                    |
| `A-Z`      | Any letter, a through z, upper-case (*see LDAP notes*) |
| `0-9`      | Any number, 0 through 9                                |
| `+`        | Plus                                                   |
| `.`        | Period                                                 |
| `_`        | Underscore                                             |
| `-`        | Hyphen                                                 |
| `&`        | Ampersand                                              |
| `*`        | Asterisk                                               |
| `^`        | Caret - circumflex                                     |
| &#96;      | Grace accent                                           |
| `!`        | Exclamation mark                                       |
| `@`        | At symbol                                              |
| `#`        | Number ("pound")                                       |
| `$`        | Dollar                                                 |
| `%`        | Percent (per cent) sign                                |
| `=`        | Equals                                                 |
| `(`        | Open parenthesis (or open bracket)                     |
| `)`        | Close parenthesis (or close bracket)                   |
| `          | `                                                      | Vertical bar |
| `{`        | Opening brace                                          |
| `}`        | Closing brace                                          |
| `[`        | Opening bracket                                        |
| `]`        | Closing bracket                                        |
| `:`        | Colon                                                  |
| `;`        | Semicolon                                              |
| `<`        | Less than (or open angled bracket)                     |
| `>`        | Greater than (or close angled bracket)                 |
| `,`        | Comma                                                  |
| `?`        | Question mark                                          |
| `/`        | Slash                                                  |
| `~`        | Equivalency sign - tilde                               |

Server administrators set final password policy (e.g., required password
length). As previously noted, these requirements change if using a LDAP
directory server. Reach out to your IT support team to learn what specific
requirements have been configured for your site.

## Features

- Optional number of required total characters
- Optional number of required special characters
- Optional number of required digits

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application.

## Requirements

The following is a loose guideline. Other combinations of Go and operating
systems for building and running tools from this repo may work, but have not
been tested.

### Building source code

- Go
  - see this project's `go.mod` file for *preferred* version
  - this project tests against [officially supported Go
    releases][go-supported-releases]
    - the most recent stable release (aka, "stable")
    - the prior, but still supported release (aka, "oldstable")
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

### Running

- Windows 10
- Ubuntu Linux 18.04+
- Red Hat Enterprise Linux 7+

## Installation

### From source

1. [Download][go-docs-download] Go
1. [Install][go-docs-install] Go
   - NOTE: Pay special attention to the remarks about `$HOME/.profile`
1. Clone the repo
   1. `cd /tmp`
   1. `git clone https://github.com/atc0005/tsm-pass`
   1. `cd tsm-pass`
1. Install dependencies (optional)
   - for Ubuntu Linux
     - `sudo apt-get install make gcc`
   - for CentOS Linux
     - `sudo yum install make gcc`
   - for Windows
     - Emulated environments (*easier*)
       - Skip all of this and build using the default `go build` command in
         Windows (see below for use of the `-mod=vendor` flag)
       - build using Windows Subsystem for Linux Ubuntu environment and just
         copy out the Windows binaries from that environment
       - If already running a Docker environment, use a container with the Go
         tool-chain already installed
       - If already familiar with LXD, create a container and follow the
         installation steps given previously to install required dependencies
     - Native tooling (*harder*)
       - see the StackOverflow Question `32127524` link in the
         [References](references.md) section for potential options for
         installing `make` on Windows
       - see the mingw-w64 project homepage link in the
         [References](references.md) section for options for installing `gcc`
         and related packages on Windows
1. Build binaries
   - for the current operating system, explicitly using bundled dependencies
         in top-level `vendor` folder
     - `go build -mod=vendor ./cmd/tsm-pass/`
   - for all supported platforms (where `make` is installed)
      - `make all`
   - for use on Windows
      - `make windows`
   - for use on Linux
     - `make linux`
1. Copy the newly compiled binary from the applicable `/tmp` subdirectory path
   (based on the clone instructions in this section) below and deploy where
   needed.
   - if using `Makefile`
     - look in `/tmp/tsm-pass/release_assets/tsm-pass/`
   - if using `go build`
     - look in `/tmp/tsm-pass/`

### Using release binaries

1. Download the [latest release][repo-url] binaries
1. Deploy
   - Place `tsm-pass` in a location of your choice
     - e.g., `/usr/local/bin/tsm-pass`

## Configuration

### Command-line Arguments

- Flags marked as **`required`** must be set via CLI flag or environment
  variable.
- Flags *not* marked as required are for settings where a useful default is

| Option               | Required | Default | Repeat | Possible       | Description                                                                                                                                        |
| -------------------- | -------- | ------- | ------ | -------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `h`, `help`          | No       | `false` | No     | `h`, `help`    | Show Help text along with the list of supported flags.                                                                                             |
| `v`, `version`       | No       | `false` | No     | `v`, `version` | Toggles emission of branding details with plugin status details. This output is disabled by default.                                               |
| `md`, `min-digits`   | No       | `10`    | No     | `1+`           | The minimum number of digits that will be used when generating a new TSM-compatible password.                                                      |
| `ms`, `min-specials` | No       | `25`    | No     | `1+`           | The minimum number of (compatible) special characters that will be used when generating a new TSM-compatible password.                             |
| `l`, `length`        | No       | `25`    | No     | `3+`           | The total number of characters to use when generating a new TSM-compatible password. See Password Requirements in the README for more information. |

## Examples

### Help output

```ShellSession
$ ./release_assets/tsm-pass/tsm-pass-74cdc0a-linux-amd64 --help

tsm-pass 74cdc0a (https://github.com/atc0005/tsm-pass)

Usage of ./release_assets/tsm-pass/tsm-pass-74cdc0a-linux-amd64:
  -l int
        The total number of characters to use when generating a new TSM-compatible password. (shorthand) (default 63)
  -length int
        The total number of characters to use when generating a new TSM-compatible password. (default 63)
  -md int
        The minimum number of digits that will be used when generating a new TSM-compatible password. (shorthand) (default 10)
  -min-digits int
        The minimum number of digits that will be used when generating a new TSM-compatible password. (default 10)
  -min-specials int
        The minimum number of (compatible) special characters that will be used when generating a new TSM-compatible password. (default 25)
  -ms int
        The minimum number of (compatible) special characters that will be used when generating a new TSM-compatible password. (shorthand) (default 25)
  -v    Whether to display application version and then immediately exit application. (shorthand)
  -version
        Whether to display application version and then immediately exit application.
```

### Default settings

```ShellSession
$ ./release_assets/tsm-pass/tsm-pass-74cdc0a-linux-amd64
&-4o7.&1+__w-.+9+mr4-ul7&9+0+.66_h&9qk&.+3q4+6__+&.3&.939r__oyf
```

### Default length, custom minimum digits, special characters

```ShellSession
$ ./release_assets/tsm-pass/tsm-pass-74cdc0a-linux-amd64 --min-specials 25 --min-digits 10
._&u7172-+&s+s&xkvp3.47h+345_.21_3j6&t_+&9s0.3_.-_-9++077s-_egc
```

### Exceeding max length

As of this writing, the password limit is believed to be 63 characters. When
requesting a longer password than that, a WARNING is emitted.

```ShellSession
$ ./release_assets/tsm-pass/tsm-pass-74cdc0a-linux-amd64 --length 100 --min-specials 25 --min-digits 10

WARNING: Password length of 100 requested.
TSM limits the total password length to 63 characters by default.
Check with your TSM server admins for the current maximum password length.

._6.j&.-u07ox44_qs2_.+xabey++73rx5xnb&z+.pb04x0&dw.sm&&vzg-ld&1y7w1_4+w_-1&&p0._.-&-kl.+9b&5.7fhmxee
```

## License

```license
MIT License

Copyright (c) 2020 Adam Chalkley

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

## References

- Heavily inspired v0.1.0 version
  - <https://yourbasic.org/golang/generate-random-string/>
- Example from Todd McLeod (initial ideas)
  - <https://play.golang.org/p/hqrdJ29qGzA>
  - <https://twitter.com/Todd_McLeod/status/1144028304061562880>

- TSM/Tivoli/Spectrum Protect
  - Password requirements
    - recent
      - <https://www.ibm.com/support/knowledgecenter/SSEQVQ_8.1.10/client/r_cmd_setpassword.html>
    - historical
      - <http://publib.boulder.ibm.com/infocenter/tivihelp/v1r1/topic/com.ibm.itsmcw.doc/anrwgd55429.htm#stmpwd>
      - <http://publib.boulder.ibm.com/infocenter/tsminfo/v6/topic/com.ibm.itsm.client.doc/r_cmd_setpassword.html>

- <https://www.ascii-code.com/>

<!-- Footnotes here  -->

[repo-url]: <https://github.com/atc0005/tsm-pass>  "This project's GitHub repo"

[go-docs-download]: <https://golang.org/dl>  "Download Go"

[go-docs-install]: <https://golang.org/doc/install>  "Install Go"

[go-supported-releases]: <https://go.dev/doc/devel/release#policy> "Go Release Policy"

<!-- []: PLACEHOLDER "DESCRIPTION_HERE" -->
