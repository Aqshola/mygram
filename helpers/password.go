package helpers

import "golang.org/x/crypto/bcrypt"

func HashPass(p string) string {
	salt := 8
	password := []byte(p)
	hash, _ := bcrypt.GenerateFromPassword(password, salt)

	return string(hash)
}

func ComparePass(p string, h string) bool {
	pass, hash := []byte(p), []byte(h)

	err := bcrypt.CompareHashAndPassword(hash, pass)

	return err == nil
}
