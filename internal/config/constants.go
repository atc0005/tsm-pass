// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/tsm-pass
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

// Shared, password related constants for easy access.
// #nosec G101 (https://github.com/securego/gosec)
const (
	PasswordSpecialCharacters string = ".!@#$%^&*_-+=`()|{}[]:;<>,?/~"
	PasswordDigits            string = "0123456789"
	PasswordLowerLetters      string = "abcdefghijklmnopqrstuvwxyz"
	PasswordUpperLetters      string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)
