package main

import (
	"context"
	"fmt"
	"time"
)

func dooSomething(ctx context.Context) {
	fmt.Println("Начало работы")
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("Конец работы")
}

func main() {
	// если указать интервал 1500 миллисекунд, то doSomething
	// успеет выполниться полностью
	ctx, cancel := context.WithTimeout(context.Background(), 4000*time.Millisecond)

	time.AfterFunc(time.Second, cancel)

	defer cancel()
	dooSomething(ctx)

	// проверяем, как завершился контекст
	switch ctx.Err() {
	case context.Canceled:
		fmt.Println("Прервали работу")
	case context.DeadlineExceeded:
		fmt.Println("Истекло время работы")
	}
}
