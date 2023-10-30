// 19.	Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	question := "Mediis tempestatibus placidus"
	fmt.Printf("%s\n%s", question, ReverseString(question))
}

func ReverseString(s string) string {
	// функция для определения длины строки в рунах. Символы не из ASCII могут представлять более 1 байта
	strLen := utf8.RuneCountInString(s)

	// создаем срез рун, определяем длину в рунах
	reverseString := make([]rune, strLen)
	ind := 1
	// цикл с range - не по байтам, а по символам
	for _, c := range s {
		reverseString[strLen-ind] = c
		ind++
	}
	return string(reverseString)
}
