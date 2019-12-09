package userservice

import "unicode"

func CheckStrongPassword(pwd string) string {
	if len(pwd) < 6 {
		return "password needs at least length of 6"
	}
	var sd, sl, su = SumDigit(pwd), SumLower(pwd), SumUpper(pwd)
	if sd < 2 {
		return "password needs at least digit count of 2"
	}
	if sl == 0 || su == 0 {
		return "both lower and upper alphabet must be in the password"
	}
	return ""
}

func SumDigit(s string) int {
	var t int
	for i := range s {
		if unicode.IsDigit(rune(s[i])) {
			t++
		}
	}
	return t
}

func SumLower(s string) int {
	var t int
	for i := range s {
		if unicode.IsLower(rune(s[i])) {
			t++
		}
	}
	return t
}

func SumUpper(s string) int {
	var t int
	for i := range s {
		if unicode.IsUpper(rune(s[i])) {
			t++
		}
	}
	return t
}
