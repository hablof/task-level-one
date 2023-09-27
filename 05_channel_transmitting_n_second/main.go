package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	expectedArgMsg       = "please specify work duration"
	mustPositiveDuration = "duration must be positive"
)

func main() {

	// Избрал реализацию установки времени работы
	// путём передачи аргумента командной строки
	// первый аргумент (с индексом 0) -- всегда имя исполняемого файла
	// второй (с индексом 1) -- ожидаемый
	// если аргумент не передан -- делать нечего
	if len(os.Args) < 2 {
		fmt.Println(expectedArgMsg)
		os.Exit(1)
	}

	// аргумент должен быть числом...
	workDuration, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// ... и числом  строго положительным
	if workDuration <= 0 {
		fmt.Println(mustPositiveDuration)
		os.Exit(1)
	}

	// И, хотя, в данной конкретной задаче нет необходимости завершать работу воркеров
	// т.к. в любом случае горутины воркеров завершатся вслед за горутиной main,
	// правильнее будет явно остановить горутины воркеров.
	// Завершать работу горутин будем путём закрытия канала из пишущей горутины

	// основной канал, через который будут поступать данные
	mainChannel := make(chan int)
	wg := &sync.WaitGroup{}

	wg.Add(1)

	go func() {
		defer wg.Done()

		fmt.Println("reader worker started")

		// читаем из канала с помощью for range
		// такой способ чтения гарантирует,
		// что из канала будут прочитаны все данные,
		// записанные до его закрытия
		for i := range mainChannel {
			fmt.Printf("recived value: %d\n", i)
		}

		fmt.Println("reader worker terminating")
	}()

	t := time.NewTimer(time.Second * time.Duration(workDuration))

	counter := 0
	// цикл основной горутины, пишущий в канал
	for {
		select {
		// использование default гарантирует, что при получении сигнала <-t.C
		// select не провалится в ветку записи "mainChannel <- randString()"
		// такое было бы возможно (и не однократно) при использовании конструкции
		// "case mainChannel <- counter"
		default:
			fmt.Printf("sending value: %d...\n", counter)
			// пишем в канал
			mainChannel <- counter
			// и увеличиваем counter на 1,
			// тем самым демонстрируя "последовательность" отправляемых значений
			counter++
			// чтобы в глазах не рябило...
			time.Sleep(100 * time.Millisecond)

		case <-t.C:
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
