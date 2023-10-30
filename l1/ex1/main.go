/*Дана структура Human (с произвольным набором полей и методов). Реализовать встраивание
методов в структуре Action от родительской структуры Human (аналог наследования).*/

package main

import "fmt"

type Human struct {
	Name   string
	Gender string
	Age    int
}

func (h Human) info() string {
	return fmt.Sprintf("%s %s %d", h.Name, h.Gender, h.Age)
}

type Action struct {
	Human      //композиция структур
	Profession string
	Hobby      string
}

func main() {

	profile := Action{
		Human:      Human{Name: "John", Gender: "Man", Age: 30},
		Profession: "Manager",
		Hobby:      "Football",
	}

	//вызов встроенного в структуру Action метода структуры Human
	fmt.Printf("This is %s, his works as a %s and likes %s\n", profile.info(), profile.Profession, profile.Hobby)

	//вызов метода встроенного типа
	fmt.Println(profile.Age)

}
