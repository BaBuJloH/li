// 26. Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abCdefAaf"
	fmt.Println(IsUniqSymbols(s))
}

func IsUniqSymbols(s string) bool {
	// map обеспечит уникальность
	m := make(map[rune]bool)

	// для регистронезависимости переводим весь регистр строки к нижнему
	s = strings.ToLower(s)
	// цикл по рунам строки
	for _, c := range s {
		if m[c] {
			return false
		}

		m[c] = true
	}
	return true
}
