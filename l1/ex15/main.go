// 15. К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.


package main

import "fmt"

// var justString string

// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

//	func main() {
//	  someFunc()
//	}
func createHugeString() []byte {
	hugeString := `Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. 
	Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. 
	Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. 
	Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

	// createHugeString возвращает слайс байт (type []byte)
	/*
		В таком случае, нужно обратить внимание:
			1. createHugeString возвращает срез []byte, который указывает на массив. Нарезка среза не делает копию базового массива.
				Поскольку срез ссылается на исходный массив, пока срез хранится, то сборщик мусора не может освободить массив; (!)
				ведь 100 полезных байтов сохраняют все содержимое в памяти.
			2. Нужно убедиться, что в слайсе байт hugeStringByte есть > 100 элементов, поскольку срез берем от 0 до 100 байта.Хотя, если по условию обещают, что строка очень большая, то это не проблема.
			3. Нужно учитывать, что изменение элементов среза изменяет элементы исходного среза.
			4. Если при добавлении элемента в срез длина увеличивается на единицу и тем самым превышает заявленный объем, необходимо
				предоставить новый объем (в этом случае текущий объем обычно удваивается). Что может привести к большим затратам по памяти

		Чтобы решить проблему 1, можно скопировать интересные данные в новый срез перед возвратом:
	*/
	hugeStringByte := []byte(hugeString)
	justStringByte := hugeStringByte[:100]        // оставил срез, для того, чтобы указать диапазон данных, например [23:28]
	newSlice := make([]byte, len(justStringByte)) // создаем пустой срез с заданной длиной
	copy(newSlice, hugeStringByte)                // копируем и возвращаемся из функции
	return newSlice
}

func main() {
	// Теперь, когда функция createHugeString() вернула копию полезных данных, то сборщик мусора может освободить массив
	fmt.Printf("Task15: %c\n", createHugeString())
}
