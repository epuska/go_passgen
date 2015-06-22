go_passgen
==========

A simple password generator implemented in Go.

Passwords are generated from a master password and id (e.g. "github.com"). Algorithm is based on SHA-256 hashes and 100k rounds PBKDF2.

Can be used as a replacement for a password manager. Doesn't save anything to disk.

