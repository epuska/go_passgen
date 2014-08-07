go_passgen
==========

Simple password generator implemented in Go.

Passwords are generated from a master password and id (e.g. "github.com"). No need to save anything to disk and only need to remember the master password.

Algorithm based on SHA-256 hashes and 100k rounds PBKDF2.
