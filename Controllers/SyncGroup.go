package Controllers

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mutex sync.Mutex
var wg sync.WaitGroup

func WaitGrp1() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementCounter()
	}
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func incrementCounter() {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	counter++
	time.Sleep(time.Millisecond * 1000) // Simulating some work
}

func WaitGrp() {
	go printNumbers()
	go printLetters()
	time.Sleep(time.Second) // Let the goroutines run for a while
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range []string{"A", "B", "C", "D", "E"} {
		fmt.Printf("%s ", letter)
		time.Sleep(100 * time.Millisecond)
	}
}

func SelectFun() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Hello"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "World"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}
}
