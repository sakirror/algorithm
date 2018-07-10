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
		matrix[i][0].op = DEL
	}
	for i := 0; i < column; i++ {
		matrix[0][i].distance = i
		matrix[0][i].op = INS
	}

	matrix[0][0].op = NOSET
	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
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

func printOperation(rowstr string, columnstr string, matrix [][]Matrix) {
	row := len(matrix) - 1
	column := len(matrix[0]) - 1
	target := func() string {
		return fmt.Sprintf("%s[%d]", rowstr, row-1)
	}

	for row > 0 || column > 0 {
		switch matrix[row][column].op {
		case INS:
			fmt.Printf("%s insert '%c'\n", target(), columnstr[column-1])
			column--
		case DEL:
			fmt.Printf("%s delete '%c'\n", target(), rowstr[row-1])
			row--
		case REP:
			if matrix[row][column].distance != matrix[row-1][column-1].distance {
				fmt.Printf("%s replace '%c' to '%c'\n", target(), rowstr[row-1], columnstr[column-1])
			}
			column--
			row--
		default:
			fmt.Printf("[%d][%d] is invalid operation %c.\n", row, column, matrix[row][column].op)
			break
		}
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
	rowstr := "kitten"
	columnstr := "sitting"

	if len(os.Args) > 2 {
		rowstr = os.Args[1]
		columnstr = os.Args[2]
	}

	matrix := prepareMatrix(rowstr, columnstr)

	if distance, err := LevenshteinDistance(rowstr, columnstr, matrix); err == nil {
		fmt.Printf("'%s' -> '%s' distance is %d\n", rowstr, columnstr, distance)
		printMatrix(rowstr, columnstr, matrix)
		printOperation(rowstr, columnstr, matrix)
	} else {
		fmt.Printf("err: %s", err)
	}
}
