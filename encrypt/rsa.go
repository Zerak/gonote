package main

import (
	"crypto/aes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	pkey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		fmt.Printf("generate key err:%v\n", err)
	}
	fmt.Printf("escape:%v\n", time.Now().Sub(now))

	fmt.Printf("pkey:\nN:%v\nE:%v\nD:%v\nPrimes:%v\n \n", pkey.N, pkey.E, pkey.D, pkey.Primes)

	aes.NewCipher([]byte("d"))

	var publicKeyData = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCZ8Xy3Fq5xHnZOeKAbAS8wbx3H
7+OPM9yKBCqeQAtK0ofxmqFN+XF7uJnJICI9mH/2aZ7O8cCSPbmjvvqgplDa7NkW
Wjp86DDf9+LrhX1ELLMVCdciLNDX7B2igftV1ii+OISHfDPvbdfxfj+7KZEFu4FC
eic4u4XToLMmA7mWUwIDAQAB
-----END PUBLIC KEY-----
`
	block, _ := pem.Decode([]byte(publicKeyData))
	pubInterface, parseErr := x509.ParsePKIXPublicKey(block.Bytes)
	if parseErr != nil {
		fmt.Println("Load public key error")
		panic(parseErr)
	}

	hash := sha1.New()
	msg := []byte("helloworldjemyme")
	var pub *rsa.PublicKey
	pub = pubInterface.(*rsa.PublicKey)
	encryptedData, encryptErr := rsa.EncryptOAEP(hash, rand.Reader, pub, msg, nil)
	if encryptErr != nil {
		fmt.Println("Encrypt data error")
		panic(encryptErr)
	}
	encodedData := base64.URLEncoding.EncodeToString(encryptedData)
	fmt.Println(encodedData)
}
