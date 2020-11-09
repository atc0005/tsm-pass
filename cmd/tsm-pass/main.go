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

package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/google/uuid"
)

// const tsmValidPassCharsRegex string = `a-zA-Z0-9+._-&`
const defaultTSMPassMaxLength int = 63

// the original Python design was to use a 36 character UUID4 string and then
// mix in 36 TSM-compliant special characters, using a 1:1 mix-in as the
// script looped over the characters in the UUID string. After randomizing the
// results, the first 63 characters were used for the final password.
func generatePassword() (string, error) {

	var finalPasswd string

	tsmValidSpecialChars := map[int64]string{
		0: "&",
		1: "+",
		2: "-",
		3: "_",
		4: ".",
	}

	passwdBase := uuid.New().String()

	for _, c := range passwdBase {
		finalPasswd += string(c)

		// randomly return one of the result codes from the list above
		maxRandomNumber := big.NewInt(int64(len(tsmValidSpecialChars)))
		nBig, rngErr := rand.Int(rand.Reader, maxRandomNumber)
		if rngErr != nil {
			return "", fmt.Errorf("unable to generate random number: %w", rngErr)
		}
		randomResultIdx := nBig.Int64()

		finalPasswd += tsmValidSpecialChars[randomResultIdx]
	}

	return finalPasswd[0 : defaultTSMPassMaxLength-1], nil

}

func main() {
	passwd, err := generatePassword()
	if err != nil {
		panic(err)
	}
	fmt.Println(passwd)
}
