// rsa-pkcs [-hash sha1] [-key id_rsa] -sign "foobar"
// rsa-pkcs [-hash sha1] [-pub id_rsa.pub] -verify "sign_data" "foobar"

package main

import (
	//_ "code.google.com/p/go.crypto/md4"
	//_ "code.google.com/p/go.crypto/ripemd160"
	"crypto"
	_ "crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	_ "crypto/sha1"
	_ "crypto/sha256"
	_ "crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var hashs = map[string]crypto.Hash{
	"md4":       crypto.MD4,
	"md5":       crypto.MD5,
	"sha1":      crypto.SHA1,
	"sha224":    crypto.SHA224,
	"sha256":    crypto.SHA256,
	"sha384":    crypto.SHA384,
	"sha512":    crypto.SHA512,
	"ripemd160": crypto.RIPEMD160,
}

func hash(hashFunc crypto.Hash, data []byte) []byte {
	h := hashFunc.New()
	h.Write(data)
	return h.Sum(nil)
}

func exit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func getPrivateKey(fname string) *rsa.PrivateKey {
	buf, errRead := ioutil.ReadFile(fname)
	if errRead != nil {
		exit(errRead)
	}
	block, _ := pem.Decode(buf)
	if block == nil {
		exit(errors.New("private key error"))
	}
	key, errParse := x509.ParsePKCS1PrivateKey(block.Bytes)
	if errParse != nil {
		exit(errParse)
	}
	return key
}

func sign(fname string, hashFunc crypto.Hash) {
	key := getPrivateKey(fname)
	data := []byte(strings.Join(flag.Args(), " "))
	sign, errSign := rsa.SignPKCS1v15(rand.Reader, key, hashFunc, hash(hashFunc, data))
	if errSign != nil {
		exit(errSign)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(sign))
}

func getPublicKey(fname string) *rsa.PublicKey {
	buf, errRead := ioutil.ReadFile(fname)
	if errRead != nil {
		exit(errRead)
	}
	block, _ := pem.Decode(buf)
	if block == nil {
		exit(errors.New("public key error"))
	}
	pub, errParse := x509.ParsePKIXPublicKey(block.Bytes)
	if errParse != nil {
		exit(errParse)
	}
	return pub.(*rsa.PublicKey)
}

func verify(fname, sign string, hashFunc crypto.Hash) {
	pub := getPublicKey(fname)
	data := []byte(strings.Join(flag.Args(), " "))
	signData, errDecode := base64.StdEncoding.DecodeString(sign)
	if errDecode != nil {
		exit(errDecode)
	}
	if err := rsa.VerifyPKCS1v15(pub, hashFunc, hash(hashFunc, data), signData); err != nil {
		fmt.Println("false")
		os.Exit(1)
	} else {
		fmt.Println("true")
		os.Exit(0)
	}
}

func getHashFunc(str string) (crypto.Hash, bool) {
	str = strings.ToLower(str)
	hashFunc, ok := hashs[strings.ToLower(str)]
	if ok && !hashFunc.Available() {
		exit(errors.New(fmt.Sprintf("hash type \"%s\" is not available", str)))
	}
	return hashFunc, ok
}

func main() {
	var (
		isSign   = flag.Bool("sign", false, "rsa sign with sha1")
		signStr  = flag.String("verify", "", "rsa verfiy")
		hashType = flag.String("hash", "sha1", "hash type, such as: md4, md5, sha1, sha224, sha256, sha384, sha512, ripemd160")
		key      = flag.String("key", "id_rsa", "rsa private key file")
		pub      = flag.String("pub", "id_rsa.pub", "rsa public key file")
	)
	flag.Parse()
	hashFunc, ok := getHashFunc(*hashType)
	if ok {
		switch {
		case *isSign:
			sign(*key, hashFunc)
		case *signStr != "":
			verify(*pub, *signStr, hashFunc)
		default:
			flag.Usage()
		}
	} else {
		flag.Usage()
	}
}
