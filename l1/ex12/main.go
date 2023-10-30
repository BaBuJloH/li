// 12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

// Функция удалит повторяющиеся элементы
func set(s []string) []string {
	var res []string
	setOfString := make(map[string]bool, len(s))

	for _, val := range s {
		setOfString[val] = true
	}
	for key := range setOfString {
		res = append(res, key)
	}
	return res
}

func main() {

	array := []string{"cat", "cat", "dog", "cat", "tree"}

	setArr := set(array)

	fmt.Printf("Arr: %v\nSet: %v\n", array, setArr)
}
