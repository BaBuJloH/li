// 18. Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"log"
	"sync"
)

type Counter struct {
	num int
	sync.Mutex
}

func (c *Counter) Inc() { // блокируем любой доступ к c.sum на время инкремента
	c.Lock()
	defer c.Unlock()
	c.num++
}

func (c *Counter) Value() int {
	return c.num
}

func main() {
	cnt := &Counter{
		num: 0,
	}

	finish := make(chan struct{})

	go Do(cnt, finish)

	select {
	case <-finish:
		log.Printf("Work done with count: %d", cnt.Value())
	}
}

func Do(cnt *Counter, finish chan struct{}) {
	wg := sync.WaitGroup{}

	//
	for i := 0; i < 20; i++ {
		wg.Add(1)

		//
		go func(num int, cnt *Counter, wg *sync.WaitGroup) {
			defer wg.Done()

			fmt.Printf("Worker %d starting\n", num)
			cnt.Inc()
			fmt.Printf("Worker %d done\n", num)
		}(i, cnt, &wg)
	}

	wg.Wait()
	finish <- struct{}{}
	close(finish)
}
