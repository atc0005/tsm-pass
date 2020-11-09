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
)

// const tsmValidPassCharsRegex string = `a-zA-Z0-9+._-&`
const defaultTSMPassMaxLength int = 63

// newTSMCharMap returns a map of int64 to TSM compliant password character.
// This is intended to be used with rand.Int to produce random password
// strings.
func newTSMCharMap() map[int64]string {

	// https://www.ascii-code.com/
	tsmCharMap := make(map[int64]string)

	// 5 special characters (literals)
	tsmValidSpecialChars := []string{
		"&", "+", "-", "_", ".",
	}
	for idx, c := range tsmValidSpecialChars {
		tsmCharMap[int64(idx)] = c
	}

	// 10 numbers (ASCII 48-57)
	extendCharMap(tsmCharMap, 48, 57)

	// 26 case-insensitive, lowercase letters (97-122)
	extendCharMap(tsmCharMap, 97, 122)

	return tsmCharMap

}

// extendCharMap is a helper function to accept a map of int64 to string and
// start/end decimal codes which represent a range of ASCII characters. The
// map is populated with string values representing ASCII characters in that
// range (e.g., 0-9 or a-z).
func extendCharMap(cMap map[int64]string, start int, end int) {
	mapIdx := int64(len(cMap))
	for i := start; i <= end; i++ {
		cMap[mapIdx] = string(i)
		mapIdx++
	}
}

func generatePassword(length int) (string, error) {

	var password string

	tsmValidChars := newTSMCharMap()

	for i := 0; i <= length; i++ {
		// randomly return one of the TSM-compliant characters
		maxRandomNumber := big.NewInt(int64(len(tsmValidChars)))
		nBig, rngErr := rand.Int(rand.Reader, maxRandomNumber)
		if rngErr != nil {
			return "", fmt.Errorf("unable to generate random number: %w", rngErr)
		}
		randomResultIdx := nBig.Int64()

		password += tsmValidChars[randomResultIdx]
	}

	return password, nil

}

func main() {
	passwd, err := generatePassword(defaultTSMPassMaxLength)
	if err != nil {
		panic(err)
	}
	fmt.Println(passwd)
}
