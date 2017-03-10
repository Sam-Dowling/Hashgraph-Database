package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"
)

type signature struct {
	R *big.Int
	S *big.Int
}

func GenKey() *ecdsa.PrivateKey {
	pubkeyCurve := elliptic.P256()
	privatekey, err := ecdsa.GenerateKey(pubkeyCurve, rand.Reader)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return privatekey
}

func CalcSignature(privKey *ecdsa.PrivateKey, data string) string {

	h := sha1.New()
	io.WriteString(h, data)
	signhash := h.Sum(nil)

	r, s, err := ecdsa.Sign(rand.Reader, privKey, signhash)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	out, err := json.Marshal(signature{r, s})
	if err != nil {
		panic(err)
	}
	return string(out)
}

func VerifySignature(s string, pubKey ecdsa.PublicKey, data string) bool {
	sig := signature{}
	if err := json.Unmarshal([]byte(s), &sig); err != nil {
		panic(err)
	}

	h := sha1.New()
	io.WriteString(h, data)
	signhash := h.Sum(nil)
	return ecdsa.Verify(&pubKey, signhash, sig.R, sig.S)
}
