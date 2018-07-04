package main

import (
	"fmt"
	"os"
)

func prepareMatrix(rowstr string, columnstr string) [][]int {
	row := len(rowstr) + 1
	column := len(columnstr) + 1
	matrix := make([][]int, row)

	for i := 0; i < row; i++ {
		matrix[i] = make([]int, column)
		matrix[i][0] = i
	}
	for i := 0; i < column; i++ {
		matrix[0][i] = i
	}

	return matrix
}

func printMatrix(rowstr string, columnstr string, matrix [][]int) {
	_rowstr := " " + rowstr
	_columnstr := " " + " " + columnstr

	for i, row := range matrix {
		if i == 0 {
			for _, c := range _columnstr {
				fmt.Printf(" %c ", c)
			}
			fmt.Println()
		}

		for j, column := range row {
			if j == 0 {
				fmt.Printf(" %c ", _rowstr[i])
			}

			fmt.Printf("%02d ", column)
		}
		fmt.Println()
	}
}

func getMin(insert int, delete int, replace int) int {
	min := insert
	if min > delete {
		min = delete
	}
	if min > replace {
		min = replace
	}
	return min
}

func getCost(rowc byte, colc byte) int {
	if rowc != colc {
		return 1
	}
	return 0
}

// LevenshteinDistance : Levenshtein distance sample code
func LevenshteinDistance(rowstr string, columnstr string, matrix [][]int) (int, error) {
	if len(rowstr)+1 != len(matrix) || len(columnstr)+1 != len(matrix[0]) {
		return -1, fmt.Errorf("matrix size is not matched\n")
	}

	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(matrix[0]); col++ {
			matrix[row][col] = getMin(
				matrix[row][col-1]+1,
				matrix[row-1][col]+1,
				matrix[row-1][col-1]+getCost(rowstr[row-1], columnstr[col-1]))
		}
	}

	return matrix[len(matrix)-1][len(matrix[0])-1], nil
}

func main() {
	row := "kitten"
	column := "sitting"

	if len(os.Args) > 2 {
		row = os.Args[1]
		column = os.Args[2]
	}

	matrix := prepareMatrix(row, column)
	// printMatrix(row, column, matrix)

	if distance, err := LevenshteinDistance(row, column, matrix); err == nil {
		fmt.Printf("'%s' - '%s' distance is %d\n", row, column, distance)
		printMatrix(row, column, matrix)
	} else {
		fmt.Printf("err: %s", err)
	}
}
