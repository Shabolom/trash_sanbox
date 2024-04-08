package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	theMine := [5]string{"rock", "ore", "ore", "rock", "ore"}
	oreChannel := make(chan string)
	minedOreChan := make(chan string)
	// Разведчик
	go func(mine [5]string) {
		for _, item := range mine {
			if item == "ore" {
				oreChannel <- item //передаем данные в oreChannel
			}
		}
	}(theMine)
	// Добытчик
	go func() {
		for i := 0; i < 3; i++ {
			foundOre := <-oreChannel //чтение из канала oreChannel
			fmt.Println("From Finder: ", foundOre)
			minedOreChan <- "minedOre" //передаем данные в minedOreChan
		}
	}()
	// Переработчик
	go func() {
		for i := 0; i < 3; i++ {
			minedOre := <-minedOreChan //чтение данных из minedOreChan
			fmt.Println("From Miner: ", minedOre)
			fmt.Println("From Smelter: Ore is smelted")
		}
		Elapsed(start)
	}()
	<-time.After(time.Second * 2) // Все еще можете игнорировать
}

func Elapsed(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Время работы программы: %s\n", elapsed)
	return
}
