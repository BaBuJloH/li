// 14. Разработать программу, которая в рантайме способна определить
// тип переменной: int, string, bool, channel из переменной типа interface{}.

package main

import "fmt"

type Definer interface {
}

func getType(d Definer) string {
	return fmt.Sprintf("Type - %T", d)
}

func main() {
	var typeDef Definer

	typeDef = 100
	fmt.Println(getType(typeDef))

	typeDef = "golang"
	fmt.Println(getType(typeDef))

	typeDef = true
	fmt.Println(getType(typeDef))

	typeDef = make(chan int)
	fmt.Println(getType(typeDef))
}
