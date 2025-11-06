package auth

import (
	"crypto/sha1"
	"crypto/subtle"
	"encoding/hex"
)

func sha1Bytes(input string) [sha1.Size]byte {
	return sha1.Sum([]byte(input))
}

func ValidateSignature(expectedHex, hashedMerchantSecret, octoPaymentUUID, status string) bool {
	data := hashedMerchantSecret + octoPaymentUUID + status
	actualHash := sha1Bytes(data)

	expectedBytes, err := hex.DecodeString(expectedHex)
	if err != nil || len(expectedBytes) != sha1.Size {
		return false
	}

	return subtle.ConstantTimeCompare(actualHash[:], expectedBytes) == 1
}
