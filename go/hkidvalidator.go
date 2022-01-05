package hkidvalidator

import (
	"bytes"
)

func checkHKID(hkid string) bool {
	l := len(hkid)
	b := bytes.ToUpper([]byte(hkid))
	var checkdigit int

	if b[l-1] == 'A' {
		checkdigit = 10
	} else {
		checkdigit = int(hkid[l-1] - '0')
	}

	var sum int
	switch l {
	case 8:
		// Ugly prepend.
		b = append([]byte{' '}, b...)
	case 9:
		break
	default:
		return false
	}

	if ((b[0] == ' ') || isChar(b[0])) && isChar(b[1]) && isDigits(b[2:8]) {
		sum = 0
		for i, v := range b[:8] {
			if v == ' ' {
				sum += 36 * 9
			} else if isChar(v) {
				sum += (int(v-'A') + 10) * (9 - i)
			} else {
				sum += int(v-'0') * (9 - i)
			}
		}
	} else {
		return false
	}

	m := sum % 11
	if m == 0 {
		return checkdigit == 0
	} else {
		return checkdigit == 11-m
	}
}

func isChar(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isDigits(bs []byte) bool {
	for _, v := range bs {
		if !isDigit(v) {
			return false
		}
	}
	return true
}
