package main

import (
	"context"
	"fmt"
	"time"
)

// В предыдущем задании вы получили время из строки.
// Теперь нужно понять, отличается ли время, которое вам передали, от текущего.
// Сравните текущее время с полученным результатом, используя методы Equal(), After() и Before().
func main() {
	currentTimeStr := "2021-09-19T15:59:41+03:00"
	// допишите код
	// ...
	currentTime, err := time.Parse(time.RFC3339, currentTimeStr)
	if err != nil {
		fmt.Println(err)
	}

	now := time.Now()
	if !now.Equal(currentTime) {
		fmt.Println("время не соответствует")
	}

	tick()
}

// Округлите текущее время до начала дня (полуночи), используя метод Truncate().
func trunc() {
	var today time.Time
	// допишите код
	// ...
	today = time.Now().Truncate(24 * time.Hour)
	fmt.Println(today)
}

// Здесь конструкция now.In(loc) приводит время к нужному часовому поясу.
func timeZone() {
	layout := "02.01.06 15:04:05 -07 MST"
	now := time.Now()
	fmt.Println(now.Format(layout))
	loc, _ := time.LoadLocation("Europe/Moscow")
	fmt.Println(now.In(loc).Format(layout))
}

// Разработчик Андрей родился 26 ноября 1993 года.
// Посчитайте количество дней до его 100-летия — относительно сегодняшнего дня.
// возвращает количество дней отнасительно даты рождения
func andrey() {
	born := time.Date(1993, time.November, 26, 0, 0, 0, 0, time.Local)
	die := time.Date(2093, time.November, 26, 0, 0, 0, 0, time.Local)
	days := die.Sub(born).Hours() / 24

	fmt.Println(days)
}

// возвращает количество дней отнасительно сегодня
func ad() {
	birthday := time.Date(2093, time.November, 26, 0, 0, 0, 0, time.Local)
	days := int(time.Until(birthday).Hours() / 24)
	// альтернативный вариант
	// days := int(duration / time.Hour / 24)
	fmt.Println(days)
}

func after() {
	time.AfterFunc(1*time.Second, func() {
		fmt.Println("Hi from AfterFunc")
	})
	fmt.Println("Hi")
	// ожидаем 2 секунды, чтобы успела запуститься функция в AfterFunc
	time.Sleep(2 * time.Second)
	fmt.Println("Goodbye")
}

// Используя Ticker, напишите программу,
// которая десять раз с интервалом в две секунды выведет разницу в секундах между текущим временем и временем запуска программы.
// Лучше выводить только целую часть секунд.
func tick() {
	start := time.Now()
	ticker := time.NewTicker(2 * time.Second)
	for i := 0; i < 10; i++ {
		t := <-ticker.C
		fmt.Println(int(t.Sub(start).Seconds()))
	}
	context.Background()
}
