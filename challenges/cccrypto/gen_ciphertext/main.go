package main

import (
	"encoding/base64"
	"fmt"

	cccrypto "cccrypto/lib"
)

func main() {

	desOne := cccrypto.New(796952)
	cc := []byte("4971 9660 3200 7972")

	encrypted := desOne.Encrypt(cc)
	fmt.Printf("Encrypted: %s\n", encrypted)

	encryptedBytes, _ := base64.StdEncoding.DecodeString(encrypted)
	decrypted := desOne.Decrypt(encryptedBytes)
	fmt.Printf("Decrypted is: %s\n", decrypted)
}
