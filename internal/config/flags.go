// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/tsm-pass
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

import "flag"

func (c *Config) handleFlagsConfig() {

	flag.BoolVar(&c.showVersion, "v", defaultDisplayVersionAndExit, versionFlagHelp+" (shorthand)")
	flag.BoolVar(&c.showVersion, "version", defaultDisplayVersionAndExit, versionFlagHelp)

	flag.IntVar(&c.minDigits, "md", defaultMinDigits, minDigitsFlagHelp+" (shorthand)")
	flag.IntVar(&c.minDigits, "min-digits", defaultMinDigits, minDigitsFlagHelp)

	flag.IntVar(&c.minSpecialChars, "ms", defaultMinSpecialChars, minSpecialCharsFlagHelp+" (shorthand)")
	flag.IntVar(&c.minSpecialChars, "min-specials", defaultMinSpecialChars, minSpecialCharsFlagHelp)

	flag.IntVar(&c.totalChars, "l", defaultTotalChars, totalCharsFlagHelp+" (shorthand)")
	flag.IntVar(&c.totalChars, "length", defaultTotalChars, totalCharsFlagHelp)

	// Allow our function to override the default Help output
	flag.Usage = Usage

	// parse flag definitions from the argument list
	flag.Parse()

}
