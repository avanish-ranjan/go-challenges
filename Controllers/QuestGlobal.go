package Controllers

import "fmt"

//odd number and even number with two function

func PickFunc() {
	inp := 9
	chOdd := make(chan bool)
	chEven := make(chan bool)

	go PrintOddNum(inp, chOdd, chEven) // Start with Odd numbers
	go PrintEvenNum(inp, chOdd, chEven)

	chOdd <- true // Signal to start printing Odd numbers
	select {}     // Block main routine to keep Goroutines running
}

func PrintOddNum(inp int, chOdd, chEven chan bool) {
	for i := 1; i <= inp; i++ {
		<-chOdd // Wait for signal to print Odd number
		if i%2 != 0 {
			fmt.Println("Odd Number: ", i)
		}
		chEven <- true // Signal Even number Goroutine
	}
}

func PrintEvenNum(inp int, chOdd, chEven chan bool) {
	for i := 1; i <= inp; i++ {
		<-chEven // Wait for signal to print Even number
		if i%2 == 0 {
			fmt.Println("Even Number: ", i)
		}
		if i < inp {
			chOdd <- true // Signal Odd number Goroutine
		}
	}
}
