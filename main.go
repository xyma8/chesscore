package main

import (
	"fmt"
)

func printMatrix(matrix [64]uint64) {
	for i := 0; i < len(matrix); i++ {
		fmt.Print(matrix[i])
		if (i+1)%8 == 0 {
			fmt.Println()
		}
	}
}

func bitboardToArray(bb uint64) (board [64]uint8) {
	for i := 0; i < 64; i++ {
		board[i] = uint8((bb >> i) & 1)
	}
	return
}

func main() {
	var board uint64 = 0
	board |= 1 << a1
	board |= 1 << 2
	board |= 1 << f1

	var whitePawns uint64 = 1<<a2 | 1<<b2 | 1<<c2 | 1<<d2 | 1<<e2 | 1<<f2 | 1<<g2 | 1<<h2
	var blackPawns uint64 = 1<<a7 | 1<<b7 | 1<<c7 | 1<<d7 | 1<<e7 | 1<<f7 | 1<<g7 | 1<<h7
	//fmt.Printf("%b \n", board)
	//board = board &^ 0b111011
	fmt.Printf("%b \n", whitePawns)
	fmt.Printf("%b \n", blackPawns)
	//test := 0b100100 &^ 0b110111
	//fmt.Printf("%b \n", test)

	//fmt.Printf("%b \n", mask)
	//fmt.Printf("%b \n", board&^mask>>0)
	//fmt.Println()
	matrix := bitboardToArray(whitePawns)
	fmt.Println(matrix)
}
