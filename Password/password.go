package Password

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) string {
	hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		fmt.Println("failed to hash the password")

	}
	return string(hashpassword)
}

func Verifypassword(givenpassword string, storedpassword string) bool {
	return true
}
