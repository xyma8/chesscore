package main

import (
	"fmt"
	"math/bits"
)

func printMatrix(matrix [64]uint64) {
	for i := 0; i < len(matrix); i++ {
		fmt.Print(matrix[i])
		if (i+1)%8 == 0 {
			fmt.Println()
		}
	}
}

func bitboardToMatrix(board uint64) (matrix [64]uint64) {
	const MAX uint64 = 18446744073709551615
	var mask uint64 = MAX
	bitLen := bits.Len64(board)

	for i := 0; i < bitLen; i++ {
		mask = MAX ^ 1<<i
		square := board &^ mask
		if square != 0 {
			square = square >> i
		}

		//fmt.Printf("%b", square)
		matrix[i] = square
	}

	return matrix
}

func main() {
	var board uint64 = 0
	board |= 1 << 0
	board |= 1 << 2
	board |= 1 << 5
	//fmt.Printf("%b \n", board)
	//board = board &^ 0b111011
	fmt.Printf("%b \n", board)
	//fmt.Println(bits.Len64(board))
	//test := 0b100100 &^ 0b110111
	//fmt.Printf("%b \n", test)

	var mask uint64 = 18446744073709551615
	mask ^= 1 << 0
	//fmt.Printf("%b \n", mask)
	//fmt.Printf("%b \n", board&^mask>>0)
	//fmt.Println()
	matrix := bitboardToMatrix(board)
	fmt.Println(matrix)
	//printMatrix(matrix)
}
