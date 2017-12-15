package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"

	"golang.org/x/crypto/ssh"
)

func generateKey() string {
	privateKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return fmt.Sprintf("Error generating key: %v", err)
	}

	pub, err := ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return fmt.Sprintf("Error generating ssh key: %v", err)
	}

	return ssh.FingerprintSHA256(pub)
}
func main() {
	for i := 0; i < 1000; i++ {
		print(".")
		generateKey()
	}
	fmt.Println("Done")

}
