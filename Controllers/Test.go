package Controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func MyFun() {
	a := [3]string{"a", "b", "c"}
	b := a[0:1]

	fmt.Println(b)

	b[0] = "hello"

	fmt.Println(b)
	fmt.Println(a)
}

//b
//a
//a,b,c
func MyFunc() {
	mySlice := make([]int, 8)
	mySlice = append(mySlice, 1)

	fmt.Println(mySlice)
	fmt.Println("length: ", len(mySlice))
	fmt.Println("capacity: ", cap(mySlice))
}

//1
//length: 2
//capacity: 7

// func IsPolidrom() {
// 	value := []int{1, 2, 3, 2, 1}
// 	finalVal := value
// 	for i, j := 0, (len(value))-1; i < j; i, j = i+1, j-1 {
// 		finalVal[i], finalVal[j] = finalVal[j], finalVal[i]
// 	}

// }
func isPalindrome(arr []int) bool {
	length := len(arr)
	for i := 0; i < length/2; i++ {
		if arr[i] != arr[length-1-i] {
			return false
		}
	}
	return true
}

/*Write a function that takes a map with string keys and interface{} values and returns a new map containing only the key-value pairs where the value is of a specific type(int)*/

// func Test1() {
// 	data := map[string]interface{}{
// 		"one":   1,
// 		"two":   "twoo",
// 		"three": 3,
// 		"four":  4.0,
// 		"five":  true,
// 		"six":   "sixx",
// 	}
// 	res := MyFuncA(data)
// 	fmt.Println("Result: ", res)
// }
// func MyFuncA(request map[string]interface{}) map[string]interface{} {
// 	var resVal []interface{}
// 	for _, v := range request {
// 		if fmt.Sprintf("%t", v) == "int" {
// 			resVal = append(resVal, v)
// 		}
// 	}

// 	return resVal
// }

//two tables Deprt d/employee e
//Emp,emiId,EmpNNAme,deptId
//Dept  deptName, deptId

//names with departmentName
//empDept

//SELECT  e.EmpName, d.DeptName From Employee e JOINS DepartMent d ON d.deptId = e.deptId ;
// SELECT e.empId, e.EmpName, e.deptId, d.deptName fROM Employee e Join DepartMent d on de.deptId = e.deptId GROUP BY empDept;
/////////-------------------------------------------------////////////////////////
func MaxNum() {
	data := []int{4, 5, 6, 5, 4}
	maxNum := 0
	for _, v := range data {
		if v > maxNum {
			maxNum = v
		}
	}
	fmt.Println("maxNum: ", maxNum)
}

func IsPalinDromeStr() bool {
	orgStr := "Jahaj"
	newStr := strings.ToLower(orgStr)
	for i := 0; i < (len(newStr) / 2); i++ {
		for newStr[i] != newStr[len(newStr)-1-i] {
			return false
		}
	}
	return true
}
func RunFunc0() {
	fmt.Println("Is Palindrome : ", IsPalinDromeInt())
}

func IsPalinDromeInt() bool {
	orgD := []int{4, 5, 6, 5, 0}
	for i := 0; i < (len(orgD) / 2); i++ {
		for orgD[i] != orgD[len(orgD)-1-i] {
			return false
		}
	}
	return true
}

func StringMan() {
	originalData := `
	{
		"boardInfo": {
			"data": "Random Data"
		},
		"things": [
			{
				"data": "Things Data 1"
			},
			{
				"data": "Things Data 2"
			},
			{
				"data": "Things Data 3"
			},
			{
				"data": "Things Data 4"
			}
		]
	}
	`
	newData := `
	{
		"data": "New Things Data"
	}`

	var oldData map[string]interface{}
	var newDt map[string]interface{}
	_ = json.Unmarshal([]byte(originalData), &oldData)
	_ = json.Unmarshal([]byte(newData), &newDt)
	newThings := oldData["things"].([]interface{})
	newThings = append(newThings, newDt)
	fmt.Println("New Things: ", newThings)

	oldData["things"] = newThings
	jsonData, _ := json.Marshal(oldData)
	fmt.Println("All data: ", string(jsonData))

}

func CheckStatus(ch chan string) {
	// apiURL := "https://www.google.com"

	res, err := http.Get(<-ch)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	// ch <- res.Status
	fmt.Printf("HTTP Status of '%s' is '%s'\n: ", string(<-ch), res.Status)
}

func RunFunc1() {
	ch := make(chan string)
	urls := []string{"https://www.google.com", "https://www.facebook.com", "https://www.instagram.com"}
	for _, url := range urls {
		go CheckStatus(ch)
		ch <- url
	}
	time.Sleep(3 * time.Second)
}

func RunFunc11() {
	data := []int{1, 2, 22, 3, 4, 5, 0, 5, 6, 6, 7, 8, 9, 10, 11}
	res := map[int]int{}
	for _, char := range data {
		res[char]++
	}
	for i, v := range res {
		fmt.Printf("The number '%d' is '%d' times\n", i, v)
	}
}

func CountCharacter() {
	data := []int{1, 2, 22, 3, 4, 5, 0, 5, 6, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 55, 5, 6, 7, 8, 9, 10, 11}
	var inp = 5
	times := 0
	for _, v := range data {
		if v == inp {
			times++
		}
	}
	fmt.Printf("The number '%d' is '%d' times", inp, times)
}

type Fruit struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Price int    `json:"price"`
}

func RunFunc() {
	orgData := []byte(`
	[
  {"name": "orange", "type": "fruit", "price":20},
  { "name": "apple", "type": "fruit", "price": 70 },
  { "name": "kiwi", "type": "fruit", "price": 50 },
  { "name": "potato", "type": "vegetables", "price": 10 }
]`)

	var newData []Fruit
	if err := json.Unmarshal(orgData, &newData); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("newData: ", SortByPrice("Asc", newData))

}
func SortByPrice(sortIn string, data []Fruit) []Fruit {
	newData := data
	if strings.ToLower(sortIn) == "asc" {
		for i := 0; i < len(newData)-1; i++ {
			for j := 0; j < len(newData)-i-1; j++ {
				// for j := 0; j < len(newData)-j-1; j++ {
				if newData[j].Price > newData[j+1].Price {
					newData[j], newData[j+1] = newData[j+1], newData[j]
				}
			}
		}
	}
	return newData
}

// for i := 0; i < n-1; i++ {
// 	for j := 0; j < n-i-1; j++ {
// 		if arr[j] > arr[j+1] {
// 			// Swap elements if they are in the wrong order
// 			arr[j], arr[j+1] = arr[j+1], arr[j]
// 		}
// 	}
