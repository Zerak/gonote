package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	"gonote/apple/apple"
)

var (
	// replace your configs here
	secret = []byte(`-----BEGIN PRIVATE KEY-----
MIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgQdJUt8xmtSoZNPHw
OWBYfQIHbkH3DfZV/jCOMaEpbRWgCgYIKoZIzj0DAQehRANCAATz5T0kaJfDJwc/
Z2yA0i+SA1JqAisoAYdNhn6s9d+pB9THiTziDsdTM2CXTld+oweQImZEGQ2qukQ4
FzKIMciI
-----END PRIVATE KEY-----`)
	keyId    = "U677AM9A48"
	teamId   = "745BFEQQ26"
	clientId = "com.qun.qun"
	//clientId    = "com.chaoqun.mobi.qunqun"
	redirectUrl = "https://chaoqun.mobi"
)

// create client_secret
func GetAppleSecret() string {
	token := &jwt.Token{
		Header: map[string]interface{}{
			"alg": "ES256",
			"kid": keyId,
		},
		Claims: jwt.MapClaims{
			"iss": teamId,
			"iat": time.Now().Unix(),
			// constraint: exp - iat <= 180 days
			"exp": time.Now().Add(24 * time.Hour).Unix(),
			"aud": "https://appleid.apple.com",
			"sub": clientId,
		},
		Method: jwt.SigningMethodES256,
	}

	ecdsaKey, _ := AuthKeyFromBytes([]byte(secret))
	ss, _ := token.SignedString(ecdsaKey)
	return ss
}

// create private key for jwt sign
func AuthKeyFromBytes(key []byte) (*ecdsa.PrivateKey, error) {
	var err error

	// Parse PEM block
	var block *pem.Block
	if block, _ = pem.Decode(key); block == nil {
		return nil, errors.New("token: AuthKey must be a valid .p8 PEM file")
	}

	// Parse the key
	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(block.Bytes); err != nil {
		return nil, err
	}

	var pkey *ecdsa.PrivateKey
	var ok bool
	if pkey, ok = parsedKey.(*ecdsa.PrivateKey); !ok {
		return nil, errors.New("token: AuthKey must be of type ecdsa.PrivateKey")
	}

	return pkey, nil
}

// do http request
func HttpRequest(method, addr string, params map[string]string) ([]byte, int, error) {
	form := url.Values{}
	for k, v := range params {
		form.Set(k, v)
	}

	var request *http.Request
	var err error
	if request, err = http.NewRequest(method, addr, strings.NewReader(form.Encode())); err != nil {
		return nil, 0, err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var response *http.Response
	if response, err = http.DefaultClient.Do(request); nil != err {
		return nil, 0, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, 0, err
	}
	return data, response.StatusCode, nil
}

func main() {
	code := "c26ac66d01e7b4e289680ab3190a80bdf.0.mzs.GsyP-ym6cpGwHJX1z91mww"
	var resp apple.ValidationResponse
	c := apple.New()
	s, _ := apple.GenerateClientSecret(string(secret), teamId, clientId, keyId)
	c.VerifyAppToken(context.Background(), apple.AppValidationTokenRequest{
		ClientID:     clientId,
		ClientSecret: s,
		Code:         code,
	}, &resp)
	fmt.Println(resp)

	set, err := jwk.FetchHTTP("https://appleid.apple.com/auth/keys", jwk.WithHTTPClient(http.DefaultClient))
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, key := range set.Keys {
		fmt.Println(key)

		k, ok := key.Get("n")
		fmt.Println(k, ok)
		pubKey := k.(*rsa.PublicKey)

		token, err := jwt.Parse(code, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return "", fmt.Errorf("rsa")
			}
			return pubKey, nil
		})

		fmt.Println(token, err)
	}
}
