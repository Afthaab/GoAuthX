package auth

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	time    = 1
	memory  = 32 * 1024
	threads = 4
	keylen  = 32
	saltlen = 16
)

func GenerateSalt() ([]byte, error) {
	salt := make([]byte, saltlen)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func HashPassword(password string) (string, error) {
	salt, err := GenerateSalt()
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keylen)

	return fmt.Sprintf("%s$%s", base64.StdEncoding.EncodeToString(salt), base64.StdEncoding.EncodeToString(hash)), nil
}

func VerifyHasedPassword(encodedHasedPassword string, password string) (bool, error) {
	parts := strings.Split(encodedHasedPassword, "$")
	if len(parts) != 2 {
		return false, errors.New("invalid hash format")
	}

	salt, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, errors.New("invalid salt encoding")
	}

	storedHash, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, errors.New("invalid hash encoding")
	}

	newHash := argon2.IDKey(storedHash, salt, time, memory, threads, keylen)

	if subtle.ConstantTimeCompare(newHash, storedHash) != 1 {
		return false, errors.New("invalid password")
	}
	return true, nil
}
