package Controllers

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func SumOfArray() {
	arr := []int{3, 5, 6}
	ch := make(chan int)
	go Sum(arr, ch)
	go close(ch)

	fmt.Println("ch: ", ch)
	resfin := 0
	for partSum := range ch {
		resfin += partSum
	}
	fmt.Println("resfin: ", resfin)
}

func Sum(arr []int, result chan int) {
	res := 0
	for _, v := range arr {
		res += v

	}
	result <- res

}

func calculateSum(arr []int, resultChan chan int) {
	sum := 0
	for _, value := range arr {
		sum += value
	}

	// Send the partial sum to the channel
	resultChan <- sum
}

func MFun() {
	// Example array
	myArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create a channel to receive the partial sums
	resultChan := make(chan int)

	// Split the array into two halves
	mid := len(myArray) / 2
	firstHalf := myArray[:mid]
	secondHalf := myArray[mid:]

	// Launch goroutines to calculate the sum concurrently for each half
	go calculateSum(firstHalf, resultChan)
	go calculateSum(secondHalf, resultChan)

	// Close the channel after both goroutines finish
	close(resultChan)

	// Receive partial sums from the channel and combine them
	finalSum := 0
	for partialSum := range resultChan {
		finalSum += partialSum
	}

	// Print the final sum
	fmt.Println("Final Sum:", finalSum)
}

// buffered and unbuffered channels

func myChan() {
	ch, ch2, ch3 := make(chan string), make(chan []int), make(chan int)
	inp := []int{2, 4, 6}
	go PrintString(ch)
	go SumNum(ch2, ch3)
	ch2 <- inp
	reslt := <-ch
	res2 := <-ch3

	time.Sleep(4 * time.Second)
	fmt.Println(reslt)
	fmt.Println("Result2: ", res2)
}

func PrintString(ch chan string) {
	fmt.Println("Go Routine 1")
	ch <- "Avanish"

}
func SumNum(ch chan []int, ch2 chan int) {
	fmt.Println("Go Routine 2")
	inp := <-ch
	var result int
	for i, _ := range inp {
		result += inp[i]
	}
	ch2 <- result
}

func Chanf() {
	ch := make(chan int, 3)
	go PrintData(ch)
	time.Sleep(3 * time.Second)
}

func PrintData(ch chan int) {
	ch <- 2
	ch <- 6
	ch <- 8
	fmt.Println("3 values added!")
	go ConsumeData(ch)
	go ConsumeData(ch)
	ch <- 9
	fmt.Println("4 values added!")

	ch <- 10
	fmt.Println("5 values added!")
}
func ConsumeData(ch chan int) {
	fmt.Println("Values consumed1: ", <-ch)
}

//bi directional channel using buffered

func ChannelFunc() {
	ch := make(chan int, 3)

	go sender(ch)

	// Start receiver goroutine
	go receiver(ch)

	// Wait for goroutines to finish
	time.Sleep(7 * time.Second)
}
func sender(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i // Sending data to channel
		time.Sleep(time.Second)
	}
	close(ch) // Closing the channel after sending all values
}

func receiver(ch <-chan int) {
	for val := range ch { // Receiving data from channel using range
		fmt.Println("Received:", val)
	}
	fmt.Println("Channel closed")
}

func MyChan() {
	inc := IncrementFunc()
	fmt.Println(inc())
	fmt.Println(inc())

}
func IncrementFunc() func() int {
	count := 0

	return func() int {
		count++
		return count
	}

}

func wordCounter(text string, wordChan chan<- string) {
	myWords := strings.Fields(text)
	for _, words := range myWords {
		wordChan <- words
	}
	close(wordChan)
}

func RunFuncn() {
	text := "my name is avanish  avanish is my name"
	ch := make(chan string)
	go wordCounter(text, ch)

	myMap := make(map[string]int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for word := range ch {
			myMap[word]++
		}
	}()
	wg.Wait()
	for i, v := range myMap {
		fmt.Printf("word '%s' is '%d' times. \n", i, v)
	}
}
