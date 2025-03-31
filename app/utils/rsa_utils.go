package utils

import (
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"math/big"
)

func ConstructRsaPublicKey(modulus string, exponent string) (*rsa.PublicKey, error) {
	// Decode the base64-encoded modulus and exponent
	nBytes, err := base64.RawURLEncoding.DecodeString(modulus)
	if err != nil {
		return nil, fmt.Errorf("failed to decode modulus: %w", err)
	}
	eBytes, err := base64.RawURLEncoding.DecodeString(exponent)
	if err != nil {
		return nil, fmt.Errorf("failed to decode exponent: %w", err)
	}

	// Convert the modulus and exponent to big integers
	n := new(big.Int).SetBytes(nBytes)
	e := new(big.Int).SetBytes(eBytes)

	// Create the RSA public key
	pubKey := &rsa.PublicKey{
		N: n,
		E: int(e.Int64()), // Ensure exponent fits into an int
	}

	return pubKey, nil
}
