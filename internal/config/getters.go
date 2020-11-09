// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/tsm-pass
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

// ShowVersion returns the user-provided choice of displaying the application
// version and exiting or the default value for this choice.
func (c Config) ShowVersion() bool {
	return c.showVersion
}

// PassMinDigits returns the user-specified number of minimum digits to be
// used in the generated password, or the default value if not provided.
func (c Config) PassMinDigits() int {
	return c.minDigits
}

// PassMinSpecialChars returns the user-specified number of minimum special
// characters to be used in the generated password, or the default value if
// not provided.
func (c Config) PassMinSpecialChars() int {
	return c.minSpecialChars
}

// PassTotalChars returns the user-specified total number of characters to be
// used in the generated password, or the default value if not provided.
func (c Config) PassTotalChars() int {
	return c.totalChars
}
