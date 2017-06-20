package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// MyCustomClaim is a customized claim type
type MyCustomClaim struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func main() {

	claim := &MyCustomClaim{Foo: "hello"}
	claim.ExpiresAt = time.Now().Add(time.Second * 86400).Unix()

	fmt.Printf("Claim: %v\n", claim)

	// Test JWT with HMAC
	key := []byte("hmackey")

	// Encode
	tokenString, err := EncodeJWTWithHMAC(claim, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("HMAC JWT Token: %s\n", tokenString)

	// Decode
	decodedClaim := &MyCustomClaim{}
	err = DecodeJWTWithHMAC(tokenString, key, decodedClaim)
	if err != nil {
		panic(err)
	}

	// Validate
	err = decodedClaim.Valid()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded Claim: %v\n", decodedClaim)

	// Test JWT with RSA
	privKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}

	// Encode
	tokenString, err = EncodeJWTWithRSA(claim, privKey)
	if err != nil {
		panic(err)
	}
	fmt.Printf("RSA JWT Token: %s\n", tokenString)

	// Decode
	decodedClaim = &MyCustomClaim{}
	err = DecodeJWTWithRSA(tokenString, &privKey.PublicKey, decodedClaim)
	if err != nil {
		panic(err)
	}

	// Validate
	err = decodedClaim.Valid()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decoded Claim: %v\n", decodedClaim)
}

// EncodeJWTWithHMAC encodes a claim into a token string with HMAC key
func EncodeJWTWithHMAC(claim jwt.Claims, key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(key)
	return tokenString, err
}

// DecodeJWTWithHMAC decodes a token string into a claim with HMAC key
func DecodeJWTWithHMAC(tokenString string, key []byte, claim jwt.Claims) error {
	_, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	return err
}

// EncodeJWTWithRSA encodes a claim into a token string with RSA private key
func EncodeJWTWithRSA(claim jwt.Claims, privKey *rsa.PrivateKey) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claim)
	tokenString, err := token.SignedString(privKey)
	return tokenString, err
}

// DecodeJWTWithRSA decodes a token string into a claim with RSA public key
func DecodeJWTWithRSA(tokenString string, pubKey *rsa.PublicKey, claim jwt.Claims) error {
	_, err := jwt.ParseWithClaims(tokenString, claim, func(token *jwt.Token) (interface{}, error) {
		return pubKey, nil
	})
	return err
}
