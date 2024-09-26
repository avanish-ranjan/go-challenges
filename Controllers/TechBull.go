package Controllers

import (
	"fmt"
	"sync"
)

// print 1-50 numbers in chunk of 10 numbers use 5 go routines
func PrintNumInChunk() {
	wg := &sync.WaitGroup{}

	ch := make(chan []int)
	numsTotal := make([]int, 50)
	for i := 0; i < 50; i++ {
		numsTotal[i] = i + 1
	}
	wg.Add(5)
	go GenerateNum(ch, wg)
	ch <- numsTotal[0:10]
	go GenerateNum(ch, wg)
	ch <- numsTotal[10:20]
	go GenerateNum(ch, wg)
	ch <- numsTotal[20:30]
	go GenerateNum(ch, wg)
	ch <- numsTotal[30:40]
	go GenerateNum(ch, wg)
	ch <- numsTotal[40:50]
	wg.Wait()
}

func GenerateNum(ch chan []int, wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println("Numbers in chunk: ", <-ch)
}

//swap number without using third variable
func SwapNum() (int, int) {
	a := 6
	b := 8
	a = a + b
	b = a - b
	a = a - b
	return a, b
}

func SwapNumbers() {
	a, b := SwapNum()
	fmt.Printf("value of a is %d, value of b is %d", a, b)
}
