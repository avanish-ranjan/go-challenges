package Controllers

import "fmt"

//to write a func to take array of int and return the sum of the all even numbers

func MyFuncn() {
	myArr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Sum is : ", SumOfEvenNum(myArr))
}
func SumOfEvenNum(data []int) int {
	sum := 0
	for _, v := range data {
		if v%2 == 0 {
			sum = sum + v
		}
	}
	return sum
}

//func to sort array of int in asc

func SortMyArr() {
	myArray := []int{1, 3000, 15, 42, 300, 28}
	fmt.Println("Original Array: ", myArray)
	for i := 0; i < len(myArray)-1; i++ {
		for j := 0; j < len(myArray)-i-1; j++ {
			if myArray[j] > myArray[j+1] {
				myArray[j], myArray[j+1] = myArray[j+1], myArray[j]
			}
		}
	}
	fmt.Println("Sorted Array of int: ", myArray)
}
