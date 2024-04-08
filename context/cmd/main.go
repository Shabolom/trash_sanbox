package main

import (
	"context"
	"fmt"
	"time"
)

func doSomething(ctx context.Context) {
	fmt.Println("Начало работы")
	for i := 0; i < 5; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Прервали работу")
			return
		default:
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}
	fmt.Println("Конец работы")

	//fmt.Println("Начало работы")
	//for i := 0; i < 5; i++ {
	//	if ctx.Err() != nil {
	//		fmt.Println("Прервали работу")
	//		return
	//	}
	//	fmt.Println(i)
	//	time.Sleep(1 * time.Second)
	//}
	//fmt.Println("Конец работы")
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	// через 1,5 секунды вызываем cancel для отмены операции
	// если cancel вызовется два раза, это не приведёт к ошибке
	time.AfterFunc(1500*time.Millisecond, cancel)

	doSomething(ctx)
}
