package main

import (
	"os"
	"bufio"
	"fmt"
	"encoding/base64"
	"crypto/sha256"
	"code.google.com/p/go.crypto/pbkdf2"
	"github.com/howeyc/gopass"
	"github.com/atotto/clipboard"
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

func getMasterHash() []byte{
	hasher := sha256.New()
	
	fmt.Printf("Master password: ")
	hasher.Write(gopass.GetPasswd())
	masterHash := hasher.Sum(nil)
		
	fmt.Printf("Checksum: -> " + base64.StdEncoding.EncodeToString(masterHash[:3]) + " <-\n")
	return masterHash
}

func getId() []byte {
	bio := bufio.NewReader(os.Stdin)
	
	fmt.Printf("Get password for ID (empty to quit): ")
	id, err := bio.ReadBytes('\n')
	check(err)
		
	return id[:len(id)-1] // remove '\n' from the end
}

func generatePassword(masterHash, id []byte) []byte {		
	hasher := sha256.New()	
	hasher.Write(id)
	
	return pbkdf2.Key(masterHash, hasher.Sum(nil), 100000, 15, sha256.New)
}

func main() {
	masterHash := getMasterHash()
	
	for {
		id := getId()
		if len(id) == 0 {
			break
		}
		
		password := generatePassword(masterHash, id)		
		
		clipboard.WriteAll(base64.StdEncoding.EncodeToString(password))
		fmt.Printf("Password copied to clipboard.\n")
		
		clear(password)
	}
	
	clear(masterHash)
	
}