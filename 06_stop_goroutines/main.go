package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("🕓 Завершились все горутины, которые мы ждали")

	wg := &sync.WaitGroup{}
	termSig := make(chan struct{})
	closingChn := make(chan struct{})
	ctx, cf := context.WithCancel(context.Background())

	wg.Add(1)
	// эта горутина завершается сама
	go func() {
		defer wg.Done()
		defer fmt.Println("✅ Я завершаюсь самостоятельно...")
	}()

	wg.Add(1)
	// Эта горутина завершится отловив панику
	go func() {
		defer wg.Done()
		defer fmt.Println("✅ Я пережила панику")
		defer func() {
			msg := recover()
			fmt.Printf("😱 panic message was: %v\n", msg)
		}()
		panic("spicial panic")
	}()

	wg.Add(1)
	// эта горутина завершается по получению значения из канала
	go func() {
		defer wg.Done()
		defer fmt.Println("✅ Прочитано значение из канала. Завершаюсь...")

		// конструкция for {select{}} здесь не обязательна
		for {
			select {
			case <-termSig:
				return

			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	wg.Add(1)
	// эта горутина завершается по закрытию канала с использованием конструкции for range
	go func() {
		defer wg.Done()
		defer fmt.Println("✅ Канал закрылся. Завершаюсь... (for range)")

		for range closingChn {

		}
	}()

	wg.Add(1)
	// эта горутина завершается по закрытию канала с использованием конструкции "val, ok := <- ch"
	go func() {
		defer wg.Done()
		defer fmt.Println("✅ Канал закрылся. Завершаюсь... (val, ok := <- ch)")

		for {
			if _, ok := <-closingChn; !ok {
				return
			}
		}
	}()

	wg.Add(1)
	// эта горутина завершается по отмене контекста (хотя, технически -- это то же самое, что и закрытие канала)
	go func() {
		defer wg.Done()
		defer fmt.Println("✅ Контекст отменён. Завершаюсь...")

		// конструкция for {select{}} здесь не обязательна
		for {
			select {
			case <-ctx.Done():
				return

			default:
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// эта горутина завершится влед за горутиной main
	// чтобы это сработало, не будем добавлять wg Add(1)
	// Выражения в defer этой горутины не выполнятся
	go func() {
		defer fmt.Println("❌ Эту надпись никто не увидит в терминале")
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("🐛 Горутина всё ещё жива")
		}
	}()

	time.Sleep(500 * time.Millisecond)

	// начинаем завершать горутины...
	close(closingChn)
	termSig <- struct{}{}
	cf()

	wg.Wait()
}
