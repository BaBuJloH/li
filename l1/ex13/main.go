// 13. Поменять местами два числа без создания временной переменной.

package main

import "fmt"

func main() {

	a := 41
	b := 13

	fmt.Printf("Изначально а = %d, b = %d\n", a, b)

	a, b = b, a
	fmt.Printf("После свапа а = %d, b = %d\n", a, b)

}
