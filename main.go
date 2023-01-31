package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/sha3"
)

func main() {
	curve := elliptic.P384()
	privatekey, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		return
	}
	publickey := append(x.Bytes(), y.Bytes()...)
	address := sha3.Sum256(publickey)
	fmt.Println("Private key:", privatekey)
	fmt.Println("Public key(X,Y):", x, y)
	fmt.Println("Address:", address)
}
