package main

import (
	"os"
	"bufio"
	"fmt"
	"encoding/base64"
	"crypto/sha256"
	"code.google.com/p/go.crypto/pbkdf2"
	"github.com/howeyc/gopass"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

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
	fmt.Printf("Master password: ")
	masterPassword := gopass.GetPasswd()
	
	bio := bufio.NewReader(os.Stdin)
	
	fmt.Printf("Generate password for ID: ")
	id, err := bio.ReadBytes('\n')
	check(err)
	id = id[:len(id)-1] // remove '\n' from the end
	
	password := generatePassword(masterPassword, id)
		
	// need to figure out better output mechanism than this:
	fmt.Printf("Password for " + string(id) + ": " + base64.StdEncoding.EncodeToString(password) + "\n")
	
	clear(masterPassword)
	clear(password)
}