package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

const (
	expectedArgMsg            = "please specify number of workers as argument"
	mustPositiveNumberWorkers = "number of workers must be positive"
)

var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString() string {
	b := make([]byte, 10+rand.Intn(10))
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func main() {

	// Избрал реализацию выбора количества воркеров
	// путём передачи аргумента командной строки
	// первый аргумент (с индексом 0) -- всегда имя исполняемого файла
	// второй (с индексом 1) -- ожидаемый
	// если аргумент не передан -- делать нечего
	if len(os.Args) < 2 {
		fmt.Println(expectedArgMsg)
		os.Exit(1)
	}

	// аргумент должен быть числом...
	numberWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ... и числом  строго положительным
	if numberWorkers <= 0 {
		fmt.Println(mustPositiveNumberWorkers)
		os.Exit(1)
	}

	// И, хотя, в данной конкретной задаче нет необходимости завершать работу воркеров
	// т.к. в любом случае горутины воркеров завершатся вслед за горутиной main,
	// правильнее будет явно остановить горутины воркеров.
	// Завершать работу горутин будем путём закрытия канала из пишущей горутины

	// основной канал, через который будут поступать данные
	mainChannel := make(chan string)
	wg := &sync.WaitGroup{}

	// создаём горутины воркеров в количестве N
	for i := 0; i < numberWorkers; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			fmt.Printf("worker %d starting\n", i)

			// читаем из канала с помощью for range
			// такой способ чтения гарантирует,
			// что из канала будут прочитаны все данные,
			// записанные до его закрытия
			for s := range mainChannel {
				fmt.Printf("worker %d printing: %s\n", i, s)
			}

			fmt.Printf("worker %d terminating\n", i)
			return
		}(i)
	}

	// канал, который должен передать сигнал завершения (Ctrl+C)
	terminationChannel := make(chan os.Signal, 1)
	signal.Notify(terminationChannel, os.Interrupt)

	// цикл основной горутины, пишущий в канал
	for {
		select {
		// использование default гарантирует, что при получении сигнала <-terminationChannel
		// select не провалится в ветку записи "mainChannel <- randString()"
		// такое было бы возможно (и не однократно) при использовании конструкции
		// "case mainChannel <- randString():"
		default:
			// пишем в канал
			mainChannel <- randString()
			// чтобы в глазах не рябило...
			time.Sleep(50 * time.Millisecond)

		case <-terminationChannel:
			fmt.Println("*** recived terminating signal ***")

			// получение сигнала завершения провоцируем завершение читающих горутин
			// путём закрытия канала с данными
			close(mainChannel)

			// Ждём завершения горутин
			wg.Wait()

			return
		}
	}
}
