package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var b strings.Builder
	inputLen := len(input)
	for i := 0; i < inputLen-1; i++ {
		current, next := rune(input[i]), rune(input[i+1])
		if unicode.IsDigit(current) {
			if unicode.IsDigit(next) || i == 0 {
				return "", ErrInvalidString
			}
			continue
		}
		if unicode.IsDigit(next) {
			digit, err := strconv.Atoi(string(next))
			if err != nil {
				return "", errors.New("error while doing Atoi")
			}
			b.WriteString(strings.Repeat(string(current), digit))
			continue
		}
		b.WriteRune(current)
	}
	if inputLen > 0 && !unicode.IsDigit(rune(input[len(input)-1])) {
		b.WriteByte(input[len(input)-1])
	}
	return b.String(), nil
}
