package config

import (
	"encoding/base64"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Base64ncryptionGenerate encrypt user password
func Base64ncryptionGenerate(pssword string) string {
	bytePwd := []byte(pssword)
	encode := base64.StdEncoding.EncodeToString(bytePwd)
	return encode
}

// Base64Compare compare passwords for consistency
func Base64Compare(password string, encode string) bool {
	bytepwd := []byte(encode)
	decode := base64.StdEncoding.EncodeToString(bytepwd)

	if result := strings.Compare(password, decode); result == 0 {
		return true
	}

	return false
}

// SaltHashGenerate encrypt user passwords
func SaltHashGenerate(password string) (string, error) {
	hex := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(hex, 10)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// SaltHashCompare compare passwords for consistency
func SaltHashCompare(digest []byte, password string) bool {
	hex := []byte(password)

	if err := bcrypt.CompareHashAndPassword(digest, hex); err == nil {
		return true
	}

	return false
}
