package Controllers

import (
	"fmt"
)

type MyInh interface {
	SumTotal()
}

type Details struct {
	Price    int
	NoOfItem int
}

func (d Details) SumTotal() int {
	return d.Price * d.NoOfItem
}

func MyMainFunc() {
	data := Details{Price: 49, NoOfItem: 3}

	fmt.Printf("Total Sum is %d : ", data.SumTotal())
}

func DuplicateInArray() {
	myArr := []int{1, 2, 3, 4, 5, 6, 2, 4, 5, 8}
	myD := map[int]int{}
	for _, v := range myArr {
		myD[v]++
	}
	duplicate := []int{}
	for i, v := range myD {
		if myD[v] > 1 {
			duplicate = append(duplicate, i)
		}
	}
	fmt.Println(duplicate)
}

// func httpReq() {
// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", "", bytes.NewBuffer([]byte(`"name":"Avanish","mob":"7999042575"`)))
// 	res, err := client.Do(req)

// }

func CounterFunc() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
func PrintFunc() {
	r := CounterFunc()
	fmt.Println(r())
	fmt.Println(r())
	fmt.Println(r())
	fmt.Println(r())
}
