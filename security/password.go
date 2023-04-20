package security

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(registredPassword string, givenPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(registredPassword), []byte(givenPassword))
	message := ""
	passCheck := true
	if err != nil {
		message = "Login or password incorect"
		passCheck = false
	}
	return passCheck, message
}
