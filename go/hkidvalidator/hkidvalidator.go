package hkidvalidator

import (
	"bytes"
)

func checkLengthAndFormat(hkid string) []byte {
	var s []byte
	if len(hkid) == 8 {
		s = append([]byte{' '}, []byte(hkid)...)
	} else if len(hkid) == 9 {
		if isChar(hkid[0]) {
			s = []byte(hkid)
		}
	} else {
		return s
	}

	if isChar(s[1]) && isDigits(s[2:8]) && (isDigit(s[8]) || s[8] == 'A') {
		return bytes.ToUpper(s)
	} else {
		return []byte{}
	}
}

func checkDigit(hkid []byte) int {
	if hkid[8] == 'A' {
		return 10
	} else {
		return int(hkid[8] - '0')
	}
}

func checkSum(hkid []byte) int {
	s := 0
	for i, v := range hkid[:8] {
		if v == ' ' {
			s += 36 * 9
		} else if isChar(v) {
			s += int(v-'A'+10) * (9 - i)
		} else {
			s += int(v-'0') * (9 - i)
		}
	}
	return s
}

func CheckHKID(hkid string) bool {
	id := checkLengthAndFormat(hkid)
	if len(id) == 0 {
		return false
	}
	cd := checkDigit(id)
	s := checkSum(id)
	m := s % 11

	if m == 0 {
		return cd == 0
	} else {
		return cd == 11-m
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
