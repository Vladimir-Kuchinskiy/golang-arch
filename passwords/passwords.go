package passwords

import (
	"crypto/hmac"
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error while generating bcrypt hash from password: %w", err)
	}

	return bs, nil
}

func ComparePassword(password string, hashedPassword []byte) error {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		return fmt.Errorf("invalid password: %w", err)
	}

	return nil
}

func SignMessage(msg []byte, key []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(msg)

	if err != nil {
		return nil, fmt.Errorf("error in signMessage while hasing message: %w", err)
	}

	signature := h.Sum(nil)

	return signature, nil
}

func CheckSign(msg, sig []byte, key []byte) (bool, error) {
	newSign, err := SignMessage(msg, key)
	if err != nil {
		return false, fmt.Errorf("error in checkSign while getting signature of message: %w", err)
	}

	same := hmac.Equal(newSign, sig)
	return same, nil
}
