package main

import (
	"fmt"
	"sync"
)

//func main() {
//	data := []byte{12, 255, 129, 2, 1, 2, 255, 130, 0, 1, 12,
//		0, 0, 17, 255, 130, 0, 2, 6, 72, 101, 108, 108,
//		111, 44, 5, 119, 111, 114, 108, 100}
//
//	var buf bytes.Buffer
//	buf.Write(data)
//	dec := gob.NewDecoder(&buf)
//	// запись с верху симетрична записи с низу по функционалу тк write запизывает данные в переменную buf
//	// так же запись bytes.Buffer.Copy симетрична записи Read(указать куда считывать данные из баффера)
//	//buf := bytes.NewBuffer(data)
//	//dec := gob.NewDecoder(buf)
//
//	dec := gob.NewDecoder(&buf)
//	m := make([]string, 0)
//	if err := dec.Decode(&m); err != nil {
//		panic(err)
//	}
//	fmt.Println(m)
//
//}

//func main() {
//	ch := make(chan int)
//	var wg sync.WaitGroup
//	wg.Add(1)
//
//	go func() {
//		v := <-ch
//		fmt.Println(v)
//		wg.Done()
//	}()
//
//	ch <- 7
//	wg.Wait()
//}

// ниже суть в том что мы после прохождения цикла закрываем канал но из закрытого канала можно считывать
// нуливый значения для типа данных канала и из-за того что мы теперь можем считывать из канала у нас срабатывает кейс done
// и он обрабатывается 1 раз дже если wg.Add(1) находится в цикле и 9 раз добавляет условие (но задается значение и не добавляется там), а выодит у нас постоянно много цифр тк
// каждая горутина передает свое значение в вечный цикл и получается так что если убрать close(done) мы просто получим функцию без выхода.
//var done = make(chan struct{})
//
//func worker(wg *sync.WaitGroup, i int) {
//	for {
//		select {
//		case <-done:
//			fmt.Println("Завершаем", i)
//			wg.Done()
//			return
//		default:
//			fmt.Println(i)
//		}
//		time.Sleep(50 * time.Millisecond)
//	}
//}
//
//func main() {
//	var wg sync.WaitGroup
//
//	// создаём горутины
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go worker(&wg, i)
//	}
//	time.Sleep(1 * time.Second)
//	// сообщаем горутинам о завершении работы
//	close(done)
//	// ждём завершения всех горутин
//	wg.Wait()
//}

//func main() {
//	var wg sync.WaitGroup
//	chIn := make(chan int)
//	chOut := make(chan int)
//	quit := make(chan struct{})
//
//	go func() {
//		for i := 0; i <= 15; i++ {
//			wg.Add(1)
//			chIn <- i
//		}
//		wg.Wait()
//		close(quit)
//	}()
//	go func() {
//		for {
//			select {
//			case <-quit:
//				return
//			default:
//				x := <-chIn
//				chOut <- x * 2
//			}
//		}
//	}()
//	go func() {
//		for {
//			fmt.Printf("%d ", <-chOut)
//			wg.Done()
//		}
//	}()
//	<-quit
//}

//func main() {
//	chIn := make(chan int)
//	chOut := make(chan int)
//	quit := make(chan struct{})
//
//	go func() {
//		for i := 0; i <= 15; i++ {
//			chIn <- i
//		}
//		close(chIn)
//	}()
//	go func() {
//		for x := range chIn {
//			chOut <- x * 2
//		}
//		close(chOut)
//	}()
//	go func() {
//		for x := range chOut {
//			fmt.Printf("%d ", x)
//		}
//		quit <- struct{}{}
//	}()
//	<-quit
//}

//func process(in1, in2 <-chan int, out chan<- int) {
//loop:
//	for {
//		select {
//		case x, ok := <-in1:
//			if !ok {
//				break loop
//			}
//			out <- x * 2
//		case x, ok := <-in2:
//			if !ok {
//				break loop
//			}
//			out <- x * 3
//		}
//	}
//	close(out)
//}

//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//	chout := make(chan int)
//
//	go func() {
//		for i := 0; i <= 20; i++ {
//			select {
//			case ch1 <- i:
//			case ch2 <- i:
//			}
//		}
//		close(ch1)
//		close(ch2)
//	}()
//	go process(ch1, ch2, chout)
//	for i := range chout {
//		fmt.Printf("%d ", i)
//	}
//}

//func count() {
//	var counter int64
//
//	var wg sync.WaitGroup
//
//	// горутины увеличивают значение счётчика
//	for i := 0; i < 25; i++ {
//		wg.Add(1)
//		go func() {
//			for e := 0; e < 2000; e++ {
//				counter++
//			}
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//	fmt.Printf("%d ", counter)
//}
//
//func main() {
//	// делаем несколько попыток
//	for i := 0; i < 5; i++ {
//		count()
//	}
//}

// та же функция но с использованием атомарных операторов
//func count() {
//	var counter int64
//
//	var wg sync.WaitGroup
//
//	// горутины увеличивают значение счётчика
//	for i := 0; i < 25; i++ {
//		wg.Add(1)
//		go func() {
//			for i := 0; i < 2000; i++ {
//				atomic.AddInt64(&counter, 1)
//			}
//			wg.Done()
//		}()
//	}
//	wg.Wait()
//	fmt.Printf("%d ", atomic.LoadInt64(&counter))
//}
//
//func main() {
//	// делаем несколько попыток
//	for i := 0; i < 5; i++ {
//		count()
//	}
//}

//var counter atomic.Int64
//
//func worker(wg *sync.WaitGroup) {
//	for i := 0; i < 10000; i++ {
//		counter.Add(1)
//	}
//	wg.Done()
//}
//
//func main() {
//	var wg sync.WaitGroup
//	for i := 0; i < 20; i++ {
//		wg.Add(1)
//		go worker(&wg)
//	}
//	wg.Wait()
//
//	// программа должна выводить 200000
//	fmt.Println(counter.Load())
//}

//func main() {
//	ch := generator("Hello")
//	for msg := range ch {
//		fmt.Println(msg)
//	}
//}
//
//// Тут ваш генератор
//func generator(msg string) chan string {
//	ch := make(chan string) // важно делать так а иначе канал не инициализируется а создается просто тип
//
//	go func() {
//		defer close(ch)
//		for i, _ := range msg {
//			ch <- fmt.Sprintf("%s: %d", msg, i)
//		}
//	}()
//
//	return ch
//}

//// add принимает на вход сигнальный канал для прекращения работы и канал с входными данными для работы,
//// а возвращает канал, в который будет отправляться результат вычислений.
//// На фоне будет запущена горутина, выполняющая вычисления до момента закрытия doneCh.
//func add(doneCh chan struct{}, inputCh chan int) chan int {
//	// канал с результатом
//	addRes := make(chan int)
//
//	// горутина, в которой добавляем к значению из inputCh единицу и отправляем результат в addRes
//	go func() {
//		// закрываем канал, когда горутина завершается
//		defer close(addRes)
//
//		// берём из канала inputCh значения, которые надо изменить
//		for data := range inputCh {
//			result := data + 1
//
//			select {
//			// если канал doneCh закрылся, выходим из горутины
//			case <-doneCh:
//				return
//			// если doneCh не закрыт, отправляем результат вычисления в канал результата
//			case addRes <- result:
//			}
//		}
//	}()
//	// возвращаем канал для результатов вычислений
//	return addRes
//}
//
//// multiply принимает на вход сигнальный канал для прекращения работы и канал с входными данными для работы,
//// а возвращает канал, в который будет отправляться результат вычислений.
//// На фоне будет запущена горутина, выполняющая вычисления до момента закрытия doneCh.
//func multiply(doneCh chan struct{}, inputCh chan int) chan int {
//	// канал с результатом
//	multiplyRes := make(chan int)
//
//	// горутина, в которой значение из inputCh умножаем на 2 и возвращаем в канал multiplyRes
//	go func() {
//		// закрываем канал multipleRes, когда горутина завершается
//		defer close(multiplyRes)
//
//		// берем из канала inputCh значения, которые надо изменить
//		for data := range inputCh {
//			// изменяем данные
//			result := data * 2
//
//			select {
//			// если канал doneCh закрылся, выходим из горутины
//			case <-doneCh:
//				return
//			// если doneCh не закрыт, отправляем результат вычисления в канал результата
//			case multiplyRes <- result:
//			}
//		}
//	}()
//
//	// возвращаем канал для результатов вычислений
//	return multiplyRes
//}
//
//// generator возвращает канал с данными
//func generator(doneCh chan struct{}, input []int) chan int {
//	// канал, в который будем отправлять данные из слайса
//	inputCh := make(chan int)
//
//	// горутина, в которой отправляем в канал  inputCh данные
//	go func() {
//		// как отправители закрываем канал, когда всё отправим
//		defer close(inputCh)
//
//		// перебираем все данные в слайсе
//		for _, data := range input {
//			select {
//			// если doneCh закрыт, сразу выходим из горутины
//			case <-doneCh:
//				return
//			// если doneCh не закрыт, кидаем в канал inputCh данные data
//			case inputCh <- data:
//			}
//		}
//	}()
//
//	// возвращаем канал для данных
//	return inputCh
//}
//
//func main() {
//	// ваши данные в слайсе
//	input := []int{1, 2, 3, 4, 5, 6, 7, 8}
//
//	// канал для сигнала к выходу из горутины
//	doneCh := make(chan struct{})
//	// при завершении программы закрываем канал doneCh, чтобы все горутины тоже завершились
//	defer close(doneCh)
//
//	// получаем канал с данными с помощью генератора
//	inputCh := generator(doneCh, input)
//
//	// ваш конвейер, сначала работает add,  потом multiply
//	resultCh := multiply(doneCh, add(doneCh, inputCh))
//
//	// выводим результат
//	for res := range resultCh {
//		fmt.Println(res)
//	}
//}

// Реализуйте функцию fanIn(), чтобы объединениить результатов функций square() и gen().
// Используйте функции gen() и square() из кода предыдущего задания. Основная функция уже реализована.
// После запуска программы, вывод должен быть таким:

func main() {
	inCh := gen(2, 3)
	ch1 := square(inCh)
	ch2 := square(inCh)
	for n := range fanIn(ch1, ch2) {
		fmt.Println(n)
	}
}

func gen(nums ...int) chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for _, n := range nums {
			outCh <- n
		}
	}()

	return outCh
}

func square(inCh chan int) chan int {
	outCh := make(chan int)
	go func() {
		defer close(outCh)
		for n := range inCh {
			outCh <- n * n
		}
	}()

	return outCh
}

// fanIn принимает несколько каналов, в которых итоговые значения
func fanIn(chs ...chan int) chan int {
	var wg sync.WaitGroup
	outCh := make(chan int)

	// определяем функцию output для каждого канала в chs
	// функция output копирует значения из канала с в канал outCh, пока с не будет закрыт
	output := func(c chan int) {
		for n := range c {
			outCh <- n
		}
		wg.Done()
	}

	// добавляем в группу столько горутин, сколько каналов пришло в fanIn
	wg.Add(len(chs))
	// перебираем все каналы, которые пришли и отправляем каждый в отдельную горутину
	for _, c := range chs {
		go output(c)
	}

	// запускаем горутину для закрытия outCh после того, как все горутины отработают
	go func() {
		wg.Wait()
		close(outCh)
	}()

	// возвращаем общий канал
	return outCh
}
