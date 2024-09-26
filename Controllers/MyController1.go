package Controllers

import "fmt"

type MySt struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type MyStr2 struct {
	MySt
}

func MyChanFunc() {

}

func Print() {
	arr := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	for _, v := range arr {
		for _, j := range v {
			fmt.Println(j)
		}
	}
}
