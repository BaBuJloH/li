// 25. Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Start sleeping...")
	Sleep(5)
	fmt.Println("Wake up!")
}

func Sleep(s int) {
	<-time.After(time.Duration(s) * time.Second)
}
