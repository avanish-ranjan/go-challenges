package Controllers

import (
	"encoding/json"
	"fmt"
	"strings"
)

func StringAppend() {
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
	var new map[string]interface{}
	_ = json.Unmarshal([]byte(originalData), &oldData)
	_ = json.Unmarshal([]byte(newData), &new)
	things := oldData["things"].([]interface{})
	fmt.Println("Old: ", len(things))
	things = append(things, new)
	fmt.Println("New: ", len(things), things)
}

//Problem: Reverse a String

func ReverseSentence() {
	sentence := "My Name Is Aryan"
	ch := make(chan string)
	go func() {
		ruineStr := []rune(sentence)
		for i, j := 0, len(ruineStr)-1; i < j; i, j = i+1, j-1 {
			ruineStr[i], ruineStr[j] = ruineStr[j], ruineStr[i]
		}
		ch <- string(ruineStr)
	}()
	result := <-ch
	fmt.Println("Reversed: ", result)
}

func CountChar() {
	input := "hii my name is avanish"
	ch := make(chan map[rune]int)
	go func() {
		charCount := make(map[rune]int)
		for _, char := range input {
			charCount[char]++
		}
		ch <- charCount
	}()
	res := <-ch
	for char, count := range res {
		fmt.Printf("Character '%c' occurs %d times\n", char, count)

	}

}

func IsPolidrome() {
	input := "Jahaj"
	normalized := strings.ToLower(strings.ReplaceAll(input, " ", ""))
	reversed := []rune(normalized)
	for i, j := 0, len(normalized)-1; i < j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = reversed[j], reversed[i]
	}
	if normalized == string(reversed) {
		fmt.Println("PolinDrome : ", string(reversed))
	} else {
		fmt.Println("Not PolinDrome : ", string(reversed))
	}
}

//Problem: Count Vowels in a String

func VowelsCount() {
	input := "my name is AvAnIsh"
	vowels := "aeiou"
	count := 0
	for _, char := range input {
		if strings.Contains(vowels, strings.ToLower(string(char))) {
			count++
		}
	}
	fmt.Println("Vowels in sentence: ", count)
}

// Problem: Count Words in a String
func WordInSentence() {
	input := "My Name is Avanish"
	words := strings.Fields(input)
	fmt.Println("words: ", len(words))

}

//Problem: Find Maximum Element

func MaxElement() {
	input := []int{23, 45, 67, 12}
	max := 0
	for _, v := range input {
		if v > max {
			max = v
		}
	}
	fmt.Println("Max: ", max)
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap elements if they are in the wrong order
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func Sorting() {
	// Example array to be sorted
	arr := []int{64, 25, 12, 22, 11}

	// Print original array
	fmt.Println("Original array:", arr)

	// Sort the array using bubble sort
	bubbleSort(arr)

	// Print sorted array
	fmt.Println("Sorted array:", arr)
}
