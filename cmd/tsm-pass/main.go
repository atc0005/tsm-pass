// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/tsm-pass
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

// Purpose:
//
//   Generate Tivoli-compliant random passwords
//
// Password requirements:
//
// Tivoli passwords can be between 0 and 64 characters, but this is
// controlled by the server administrator(s).
// http://publib.boulder.ibm.com/infocenter/tivihelp/v1r1/topic/com.ibm.itsmcw.doc/anrwgd55429.htm#stmpwd
//
// Other sources say 63 characters, so it is probably better to use that as the
// limit.
// http://publib.boulder.ibm.com/infocenter/tsminfo/v6/topic/com.ibm.itsm.client.doc/r_cmd_setpassword.html
//
// Tivoli passwords are _not_ case-sensitive and can be be composed of these
// valid characters:
//
// a-z     Any letter, a through z, upper or lower-case
// 0-9     Any number, 0 through 9
// +       Plus
// .       Period
// _       Underscore
// -       Hyphen
// &       Ampersand
//
// History:
//
// The original Python design was to use a 36 character UUID4 string and then
// mix in 36 TSM-compliant special characters, using a 1:1 mix-in as the
// script looped over the characters in the UUID string. After randomizing the
// results, the first 63 characters were used for the final password.
//
// This implementation builds a map of all valid characters and then
// randomizes which characters from that map will be used to generate the
// final password string. This produces a string which differs significantly
// from earlier output, but has a higher degree of randomness than before.

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// const tsmValidPassCharsRegex string = `a-zA-Z0-9+._-&`
const defaultTSMPassMaxLength int = 63

func generatePassword(length int) (string, error) {

	// Based on https://yourbasic.org/golang/generate-random-string/

	tsmSpecialChars := "&+-_."
	digits := "0123456789"
	letters := "abcdefghijklmnopqrstuvwxyz"
	allValidChars := []rune(tsmSpecialChars + digits + letters)

	var password strings.Builder

	for i := 0; i <= length; i++ {
		maxRandomNumber := big.NewInt(int64(len(allValidChars)))
		nBig, rngErr := rand.Int(rand.Reader, maxRandomNumber)
		if rngErr != nil {
			return "", fmt.Errorf("unable to generate random number: %w", rngErr)
		}
		randomResultIdx := nBig.Int64()

		password.WriteRune(allValidChars[randomResultIdx])
	}

	return password.String(), nil

}

func main() {
	passwd, err := generatePassword(defaultTSMPassMaxLength)
	if err != nil {
		panic(err)
	}
	fmt.Println(passwd)
}
