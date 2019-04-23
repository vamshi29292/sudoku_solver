package main

import (
	"errors"
	"fmt"
	"log"
)

type sudokuArray [9][9]int

func buildSudokuArrayFromInput() sudokuArray {
	var result sudokuArray
	for i := 0; i < 9; i++ {
		var row int
		_, err := fmt.Scan(&row)
		if err != nil {
			log.Fatal(err)
		}
		for j := range result[i] {
			result[i][8-j] = row % 10
			row = row / 10
		}
	}
	return result
}

func findSlotToFill(s sudokuArray) (int, int, error) {
	for i, v := range s {
		for j, x := range v {
			if x == 0 {
				return i, j, nil
			}
		}
	}
	return 0, 0, errors.New("No slots to fill")
}

func checkRows(i, j int, s sudokuArray, values map[int]bool) {
	for _, v := range s[i] {
		if v > 0 {
			values[v] = false
		}
	}
}

func checkColumns(i, j int, s sudokuArray, values map[int]bool) {
	for x := 0; x < 9; x++ {
		if s[x][j] > 0 {
			values[s[x][j]] = false
		}
	}
}

func checkSquare(i, j int, s sudokuArray, values map[int]bool) {
	up := (i / 3) * 3
	down := up + 3
	left := (j / 3) * 3
	right := left + 3
	for up < down {
		left := (j / 3) * 3
		for left < right {
			if s[up][left] > 0 {
				values[s[up][left]] = false
			}
			left++
		}
		up++
	}
}

func findPotentialValues(i, j int, s sudokuArray) ([]int, error) {
	values := map[int]bool{1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true, 9: true}
	checkRows(i, j, s, values)
	checkColumns(i, j, s, values)
	checkSquare(i, j, s, values)
	potentialValues := make([]int, 0)
	for k, v := range values {
		if v == true {
			potentialValues = append(potentialValues, k)
		}
	}
	if len(potentialValues) == 0 {
		return []int{}, errors.New("No potential values")
	}
	return potentialValues, nil
}

func sudokuSolver(sArr sudokuArray) (sudokuArray, error) {
	i, j, err := findSlotToFill(sArr)
	if err != nil {
		return sArr, nil
	}
	potentialValues, err := findPotentialValues(i, j, sArr)
	if err != nil {
		return sArr, errors.New("No possible entries")
	}
	for len(potentialValues) > 0 {
		newSArr := sArr
		newSArr[i][j] = potentialValues[0]
		potentialValues = potentialValues[1:]
		result, err := sudokuSolver(newSArr)
		if err == nil {
			return result, nil
		}
	}
	return sArr, errors.New("Ran out of possible entries")
}

func printSudokuLine(l [9]int) {
	fmt.Printf("%d %d %d | %d %d %d | %d %d %d\n", l[0], l[1], l[2], l[3], l[4], l[5], l[6], l[7], l[8])
}

func printSudoku(s sudokuArray) {
	i := 0
	for i < 9 {
		printSudokuLine(s[i])
		printSudokuLine(s[i+1])
		printSudokuLine(s[i+2])
		print("- - - - - - - - - - - \n")
		i = i + 3
	}
	for i < 3 {
		fmt.Print("\n")
		i++
	}
}

func main() {
	inp := buildSudokuArrayFromInput()
	result, err := sudokuSolver(inp)
	if err != nil {
		log.Fatal(err)
	}
	printSudoku(result)
}
