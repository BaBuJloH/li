// 11.	Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func main() {
	array1 := []int{54, 21, 0, 12, 65, 99}
	array2 := []int{6, 45, 65, 87, 21}

	// создаем импровизированное множество в виде карты
	set1 := createSet(array1)
	set2 := createSet(array2)

	resultSet := setIntersection(set1, set2)

	// вывод в виде "множества"
	setOutput(resultSet)
}

// метод создает карту с уникальными ключами (множество)
func createSet(arr []int) map[int]bool {
	set := make(map[int]bool)
	for _, val := range arr {
		set[val] = true
	}
	return set
}

func setIntersection(set1, set2 map[int]bool) map[int]bool {
	set := make(map[int]bool)
	// сохраняем ключи из первого множества, которые имеются во втором множестве
	for key := range set1 {
		if set2[key] {
			set[key] = true
		}
	}
	return set
}

func setOutput(set map[int]bool) {

	for key := range set {
		fmt.Printf("%v\n", key)
	}
}
