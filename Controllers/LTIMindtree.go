package Controllers

import (
	"fmt"
	"strings"
)

// find the not repeated words from the sentence
func NotRepeated() {
	sentence1 := "I am a doctor"
	sentence2 := "I am a software engineer"

	wordArr1 := strings.Fields(sentence1)
	wordArr2 := strings.Fields((sentence2))

	wrdCount1 := make(map[string]int)
	wrdCount2 := make(map[string]int)

	for _, wrd1 := range wordArr1 {
		wrdCount1[wrd1]++
	}
	for _, wrd2 := range wordArr2 {
		wrdCount2[wrd2]++
	}

	for wrd := range wrdCount1 {
		if wrdCount2[wrd] == 0 {
			fmt.Println(wrd)
		}
	}
	for wrd := range wrdCount2 {
		if wrdCount1[wrd] == 0 {
			fmt.Println(wrd)
		}
	}

}

// Employee =>  id, name, salaryId
//Salary  =>   id, amount, gradeId
// Grade  =>   id, name

//empName, GradeName, Amount=> 50k

//SELECT emp.name as empName, g.name as GradeName, s.amount as Amount FROM Employee emp JOIN Salary s s.id=emp.salaryId AND Grade g g.id=s.gradeId WHERE s.amount > 50000
// SELECT
//     emp.name AS empName,
//     g.name AS GradeName,
//     s.amount AS Amount
// FROM
//     Employee emp
// JOIN
//     Salary s ON s.id = emp.salaryId
// JOIN
//     Grade g ON g.id = s.gradeId
// WHERE
//     s.amount > 50000;
// https://go.dev/play/p/NfR_hJsSIuK
