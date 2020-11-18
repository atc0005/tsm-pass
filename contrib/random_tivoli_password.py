#!/usr/bin/env python

# Copyright 2020 Adam Chalkley
#
# https://github.com/atc0005/tsm-pass
#
# Licensed under the MIT License. See LICENSE file in the project root for
# full license information.

# Purpose:
#
#   Generates random passwords by generating UUID and inserting random
#   "special" characters that comply with Tivoli password requirements
#
# Password requirements:
#
# * See project README

from __future__ import print_function

import uuid
from random import shuffle

# Grab a UUID and convert to a string. We'll use this as our "base" string that
# we will further manipulate
new_password_base = str(uuid.uuid4())

# These are the only special characters allowed by Tivoli
special_characters = [
    '&', '+', '-', '_', '.',
    '!', '@', '#', '$', '%',
    '^','*', '=', '`', '(',
    ')', '|', '{', '}', '[',
    ']', ':', ';', '<', '>',
    ',', '?', '/', '~',
]

# Going to dynamically build this
new_password_char_list = []
new_password = ''

for char in new_password_base:

    # Spin the bottle
    shuffle(special_characters)

    # Grab the first special character from the shuffled list
    new_password_char_list.append(char + special_characters[0])

# One last shuffle
shuffle(new_password_char_list)

for char in new_password_char_list:
    new_password += char

# Replace any leading '-' characters with underscores instead and then
# recompose the password string
new_password = new_password[0].replace('-', '_') + new_password[1:]

# No matter the current length, we only want the first 63 characters to comply
# with Tivoli password length restrictions
print(new_password[0:63])
