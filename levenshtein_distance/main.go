package main

import (
	"fmt"
	"os"
)

// Op is the type of operation type
type Op rune

const (
	NOSET Op = 'N'
	INS   Op = 'I'
	DEL   Op = 'D'
	REP   Op = 'R'
)

// Matrix of Levenshtein Distance
type Matrix struct {
	distance int
	op       Op
}

func prepareMatrix(rowstr string, columnstr string) [][]Matrix {
	row := len(rowstr) + 1
	column := len(columnstr) + 1
	matrix := make([][]Matrix, row)

	for i := 0; i < row; i++ {
		matrix[i] = make([]Matrix, column)
		matrix[i][0].distance = i
	}
	for i := 0; i < column; i++ {
		matrix[0][i].distance = i
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			matrix[i][j].op = NOSET
		}
	}

	return matrix
}

func printMatrix(rowstr string, columnstr string, matrix [][]Matrix) {
	_rowstr := " " + rowstr
	_columnstr := " " + " " + columnstr

	for i, row := range matrix {
		if i == 0 {
			for _, c := range _columnstr {
				fmt.Printf(" %c  ", c)
			}
			fmt.Println()
		}

		for j, m := range row {
			if j == 0 {
				fmt.Printf("  %c ", _rowstr[i])
			}

			fmt.Printf("%02d%c ", m.distance, m.op)
		}
		fmt.Println()
	}
}

func getMin(insert int, delete int, replace int) (int, Op) {
	min := insert
	op := INS
	if min > delete {
		min = delete
		op = DEL
	}
	if min > replace {
		min = replace
		op = REP
	}
	return min, op
}

func getCost(rowc byte, colc byte) int {
	if rowc != colc {
		return 1
	}
	return 0
}

// LevenshteinDistance : Levenshtein distance sample code
func LevenshteinDistance(rowstr string, columnstr string, matrix [][]Matrix) (int, error) {
	if len(rowstr)+1 != len(matrix) || len(columnstr)+1 != len(matrix[0]) {
		return -1, fmt.Errorf("matrix size is not matched\n")
	}

	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(matrix[0]); col++ {
			matrix[row][col].distance, matrix[row][col].op = getMin(
				matrix[row][col-1].distance+1,
				matrix[row-1][col].distance+1,
				matrix[row-1][col-1].distance+getCost(rowstr[row-1], columnstr[col-1]))
		}
	}

	return matrix[len(matrix)-1][len(matrix[0])-1].distance, nil
}

func main() {
	row := "kitten"
	column := "sitting"

	if len(os.Args) > 2 {
		row = os.Args[1]
		column = os.Args[2]
	}

	matrix := prepareMatrix(row, column)

	if distance, err := LevenshteinDistance(row, column, matrix); err == nil {
		fmt.Printf("'%s' - '%s' distance is %d\n", row, column, distance)
		printMatrix(row, column, matrix)
	} else {
		fmt.Printf("err: %s", err)
	}
}
