package Controllers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MyArr() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the grid description: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inpData := strings.Split(input, " ")
	fmt.Println(inpData)
	// inpData := []string{"12#45#33", "94#54#23", "98#59#27"}
	var arrData [][]int
	for _, v := range inpData {
		strVal := strings.SplitAfter(v, "#")
		var rowInts []int
		for _, num := range strVal {
			val := 0
			fmt.Sscanf(num, "%d", &val)
			rowInts = append(rowInts, val)
		}
		arrData = append(arrData, rowInts)

	}
	resp := funcMaxMinimum(arrData)
	fmt.Println("My response: ", resp)
}

func funcMaxMinimum(grid [][]int) []string {
	n := len(grid)
	maxMin := 0
	var positions []string

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			min := grid[i][j]

			for x := -1; x <= 1; x++ {
				for y := -1; y <= 1; y++ {
					if x == 0 && y == 0 {
						continue
					}
					ni, nj := i+x, j+y
					if ni >= 0 && ni < n && nj >= 0 && nj < n {
						if grid[ni][nj] < min {
							min = grid[ni][nj]
						}
					}
				}
			}

			if min > maxMin {
				maxMin = min
				positions = []string{fmt.Sprintf("%d,%d", i+1, j+1)}
			} else if min == maxMin {
				positions = append(positions, fmt.Sprintf("%d,%d", i+1, j+1))
			}
		}
	}

	return positions
}
