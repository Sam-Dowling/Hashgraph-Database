package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
)

func GenKey() *rsa.PrivateKey {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return privatekey
}

func CalcSignature(privKey *rsa.PrivateKey, data string) string {
	h := sha256.New()
	h.Write([]byte(data))
	d := h.Sum(nil)
	sig, err := rsa.SignPKCS1v15(rand.Reader, privKey, crypto.SHA256, d)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return base64.StdEncoding.EncodeToString(sig)
}

func VerifySignature(pubKey rsa.PublicKey, signature string, data string) bool {
	h := sha256.New()
	h.Write([]byte(data))
	d := h.Sum(nil)
	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ok := rsa.VerifyPKCS1v15(&pubKey, crypto.SHA256, d, sigBytes)

	return ok == nil

}
