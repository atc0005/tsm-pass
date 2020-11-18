// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/tsm-pass
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// Updated via Makefile builds. Setting placeholder value here so that
// something resembling a version string will be provided for non-Makefile
// builds.
var version string = "x.y.z"

// ErrVersionRequested indicates that the user requested application version
// information
var ErrVersionRequested = errors.New("version information requested")

// Config represents the application configuration as specified via
// command-line flags.
type Config struct {

	// The minimum number of digits that will be used when generating a new
	// TSM-compatible password.
	minDigits int

	// The minimum number of (compatible) special characters that will be used
	// when generating a new TSM-compatible password.
	minSpecialChars int

	// The total number of characters to use when generating a new
	// TSM-compatible password.
	totalChars int

	// showVersion is a flag indicating whether the user opted to display only
	// the version string and then immediately exit the application
	showVersion bool
}

// Usage is a custom override for the default Help text provided by the flag
// package. Here we prepend some additional metadata to the existing output.
var Usage = func() {
	fmt.Fprintln(flag.CommandLine.Output(), "\n"+Version()+"\n")
	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
}

// Version emits application name, version and repo location.
func Version() string {
	return fmt.Sprintf("%s %s (%s)", myAppName, version, myAppURL)
}

// New is a factory function for instantiating new application configurations.
// This function handles processing flag values, including validation of
// provided values.
func New() (*Config, error) {

	var config Config

	config.handleFlagsConfig()

	// Return immediately if user just wants version details
	if config.ShowVersion() {
		return nil, ErrVersionRequested
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return &config, nil

}

// validate verifies all Config struct fields have been provided acceptable
// values.
func (c Config) validate() error {

	if c.PassMinDigits()+c.PassMinSpecialChars() > c.PassTotalChars() {
		return fmt.Errorf(
			"min digits (%d) + min special chars (%d): %d chars; "+
				"required chars greater than total chars limit (%d)",
			c.PassMinDigits(),
			c.PassMinSpecialChars(),
			c.PassMinDigits()+c.PassMinSpecialChars(),
			c.totalChars,
		)
	}

	if c.PassTotalChars() > defaultTotalChars {
		fmt.Printf(
			"\nWARNING: Password length of %d requested.\n"+
				"TSM limits the total password length "+
				"to %d characters by default.\n"+
				"Check with your TSM server admins for the current "+
				"maximum password length.\n\n",
			c.PassTotalChars(),
			defaultTotalChars,
		)
	}

	// Optimist
	return nil

}
