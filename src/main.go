package main

import (
	"log"
	"encoding/base64"
	"crypto/sha256"
	"code.google.com/p/go.crypto/pbkdf2"
)

func clear(b []byte) {
    for i := 0; i < len(b); i++ {
        b[i] = 0;
    }
}

func generatePassword(masterPassword, id []byte) []byte {
	defer clear(masterPassword)
	
	hasher := sha256.New()
	
	hasher.Write(masterPassword)
	key := hasher.Sum(nil)
	
	hasher.Reset()
	hasher.Write(id)
	salt := hasher.Sum(nil)
		
	return pbkdf2.Key(key, salt, 100000, 15, sha256.New)
}

func main() {
	masterPassword := []byte("tesstpa!#)ssword")
	id := []byte("github")
	
	log.Print(string(id) + ": " + base64.StdEncoding.EncodeToString(generatePassword(masterPassword, id)))
}