// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/tsm-pass
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

// App specific shared values
const (
	myAppName string = "tsm-pass"
	myAppURL  string = "https://github.com/atc0005/tsm-pass"
)

// Help text used with flags
const (
	versionFlagHelp         string = "Whether to display application version and then immediately exit application."
	minDigitsFlagHelp       string = "The minimum number of digits that will be used when generating a new TSM-compatible password."
	minSpecialCharsFlagHelp string = "The minimum number of (compatible) special characters that will be used when generating a new TSM-compatible password."
	totalCharsFlagHelp      string = "The total number of characters to use when generating a new TSM-compatible password. See Password Requirements in the README for more information."
)

// Default flag settings if not overridden by user input
const (
	defaultMinDigits             int  = 10
	defaultMinSpecialChars       int  = 25
	defaultTotalChars            int  = 63
	defaultDisplayVersionAndExit bool = false
)

// Shared, password related constants for easy access.
// #nosec G101 (https://github.com/securego/gosec)
const (
	PasswordSpecialCharacters string = ".!@#$%^&*_-+=`()|{}[]:;<>,?/~"
	PasswordDigits            string = "0123456789"
	PasswordLowerLetters      string = "abcdefghijklmnopqrstuvwxyz"
	PasswordUpperLetters      string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)
