package utils

import "golang.org/x/crypto/bcrypt"

// ====================================================================
// UNTUK HASHING
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// ====================================================================
// COMPARE HASHING
func CheckPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
