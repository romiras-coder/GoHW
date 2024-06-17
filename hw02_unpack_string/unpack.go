package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

// Структура для подсчета символов и их повторений
//  типа {"a": 1, "b": 4, "c": 0}
type charCounter struct {
	charItem string // символ
	cntItem  int    // кол-во
}

// Парсинг строки на струкруты
func generateItem(data []rune) ([]charCounter, error) {
	result := make([]charCounter, 0)
	for i := 0; i < len(data); i++ {

		// Если первый символ число - возвращаем ошибку
		if unicode.IsDigit(data[i]) {
			return nil, ErrInvalidString
		}

		count := 1
		char := string(data[i])

		if i < len(data)-1 && unicode.IsDigit(data[i+1]) {
			count = int(data[i+1]) - 0x30
			i++
		}
		result = append(result, charCounter{charItem: char, cntItem: count})
	}
	return result, nil
}

// Unpack Распаковка строки
func Unpack(packingString string) (string, error) {

	tokens, err := generateItem([]rune(packingString))

	if err != nil {
		return "", ErrInvalidString
	}

	var builder strings.Builder
	for _, t := range tokens {
		builder.WriteString(strings.Repeat(t.charItem, t.cntItem))
	}
	return builder.String(), nil
}
