// Copyright 2020 Adam Chalkley
//
// https://github.com/atc0005/tsm-pass
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

// Generate Tivoli-compliant random passwords
package main

import (
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"math/big"
	mrand "math/rand"
	"os"

	"github.com/atc0005/tsm-pass/internal/config"
)

func generatePassword(length int, reqNums int, reqSpecialChars int) (string, error) {

	// Based on https://yourbasic.org/golang/generate-random-string/

	if reqNums+reqSpecialChars > length {
		return "", fmt.Errorf(
			"numbers + special characters greater than total password length",
		)
	}

	tsmSpecialChars := []byte(config.PasswordSpecialCharacters)
	digits := []byte(config.PasswordDigits)
	lowerLetters := []byte(config.PasswordLowerLetters)
	upperLetters := []byte(config.PasswordUpperLetters)
	allValidChars := append(tsmSpecialChars, digits...)
	allValidChars = append(allValidChars, lowerLetters...)
	allValidChars = append(allValidChars, upperLetters...)

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

	cfg, cfgErr := config.New()
	switch {
	case errors.Is(cfgErr, config.ErrVersionRequested):
		fmt.Println(config.Version())
		os.Exit(0)
	case cfgErr == nil:
		// do nothing for this one
	default:
		fmt.Printf("\nfailed to initialize application: %s\n", cfgErr)
		flag.Usage()
		os.Exit(1)
	}

	passwd, passwdErr := generatePassword(
		cfg.PassTotalChars(),
		cfg.PassMinDigits(),
		cfg.PassMinSpecialChars(),
	)

	if passwdErr != nil {
		panic(passwdErr)
	}

	fmt.Println(passwd)
}
