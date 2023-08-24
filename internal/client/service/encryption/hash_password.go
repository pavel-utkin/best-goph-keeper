package encryption

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) (string, error) {
	h := hmac.New(sha256.New, []byte(password))
	_, err := h.Write([]byte(password))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
