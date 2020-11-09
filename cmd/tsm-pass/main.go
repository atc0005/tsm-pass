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
	mrand "math/rand"
)

const defaultTSMPassMaxLength int = 63
const requiredNumbers int = 10
const requiredSpecialCharacters int = 25

func generatePassword(length int, reqNums int, reqSpecialChars int) (string, error) {

	// Based on https://yourbasic.org/golang/generate-random-string/

	if reqNums+reqSpecialChars > length {
		return "", fmt.Errorf(
			"numbers + special characters greater than total password length",
		)
	}

	tsmSpecialChars := []byte("&+-_.")
	digits := []byte("0123456789")
	letters := []byte("abcdefghijklmnopqrstuvwxyz")
	allValidChars := append(tsmSpecialChars, digits...)
	allValidChars = append(allValidChars, letters...)

	password := make([]byte, length)
	var passIdx int

	getByte := func(b []byte) (byte, error) {
		maxRandNum := big.NewInt(int64(len(b)))

		nBig, rngErr := rand.Int(rand.Reader, maxRandNum)
		if rngErr != nil {
			return 0, fmt.Errorf("unable to generate random number: %w", rngErr)
		}

		return b[nBig.Int64()], nil
	}

	for i := 0; i < reqNums; i++ {
		// fmt.Fprintln(os.Stderr, i)
		val, err := getByte(digits)
		if err != nil {
			return "", err
		}
		password[passIdx] = val
		passIdx++
	}

	for i := 0; i < reqSpecialChars; i++ {
		// fmt.Fprintln(os.Stderr, i)
		val, err := getByte(tsmSpecialChars)
		if err != nil {
			return "", err
		}
		password[passIdx] = val
		passIdx++
	}

	remainingChars := length - reqNums - reqSpecialChars
	for i := 0; i < remainingChars; i++ {
		// fmt.Fprintln(os.Stderr, i)
		val, err := getByte(allValidChars)
		if err != nil {
			return "", err
		}
		password[passIdx] = val
		passIdx++
	}

	mrand.Shuffle(len(password), func(i int, j int) {
		password[i], password[j] = password[j], password[i]
	})

	// if password has leading dash, replace with underscore (GH-15)
	if password[0] == '-' {
		password[0] = '_'
	}

	return string(password), nil

}

func main() {
	passwd, err := generatePassword(
		defaultTSMPassMaxLength,
		requiredNumbers,
		requiredSpecialCharacters,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(passwd)
}
