package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

var cost = 10

func SetCost(ecost int) {
	cost = ecost
}

func GetCost() int {
	return cost
}

func NewPassword(x string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(x), cost)
}

func NewPasswordString(x string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(x), cost)
	return string(b), err
}

func CheckPassword(x string, pswd []byte) (bool, error) {
	if err := bcrypt.CompareHashAndPassword(pswd, []byte(x)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			err = nil
		}
		return false, err
	}
	return true, nil
}

func CheckPasswordString(x string, pswd string) (bool, error) {
	return CheckPassword(x, []byte(pswd))
}
