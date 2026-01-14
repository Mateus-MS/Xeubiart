package internal_security

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidPassword = errors.New("invalid password hash")
)

type PasswordHash struct {
	hash string
}

func (ph PasswordHash) GetHash() string {
	return ph.hash
}

func HashPassword(password string) (PasswordHash, error) {
	if password == "" {
		return PasswordHash{}, ErrInvalidPassword
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return PasswordHash{hash: string(bytes)}, err
}

func CheckPassword(hashedPassword PasswordHash, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword.GetHash()), []byte(plainPassword))
	return err == nil
}
