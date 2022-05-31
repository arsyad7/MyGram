package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) string {
	salt := 8
	password := []byte(p)
	hashed, _ := bcrypt.GenerateFromPassword(password, salt)
	return string(hashed)
}

func ComparePass(hp, p string) bool {
	hash, pass := []byte(hp), []byte(p)
	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err == nil
}
