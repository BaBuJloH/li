/*
Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают
произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

Корректно завершить работу программы, гарантируя, что все ожидающие задачи и соединения будут правильно закрыты.

	Невыполнение этого требования может привести к потере данных или другим нежелательным последствиям.
*/
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) { // функцию handleSignal для обработки полученного сигнала
	fmt.Println("Received signal:", sig) // Здесь выполняются задачи по очистке
}

func main() {
	sigChan := make(chan os.Signal, 1) // создаем сигнальный канал, уведомляем его о сигналах SIGINT и SIGTERM
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() { // Ждем сигнала
		sig := <-sigChan  // Обрабатываем сигнал
		handleSignal(sig) //Как только сигнал получен, мы вызываем функцию handleSignal и выполняем необходимую очистку
		os.Exit(0)
	}()

	var N int
	fmt.Println("Введите количество воркеров")
	fmt.Scanf("%d", &N)

	chanInput := make(chan int)

	// создание горутин с порядковым номером и значением
	for i := 0; i < N; i++ {
		go worker(i, chanInput)
	}

	// запись данных в канал из главного потока
	for {
		chanInput <- (rand.Intn(1000))
		time.Sleep(time.Second)
	}
}

func worker(workerNum int, in <-chan int) {

	for {
		num := <-in
		fmt.Printf("Воркер %d, Вывод: %d\n", workerNum, num)
	}
}
