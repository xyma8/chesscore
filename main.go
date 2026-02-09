package main

import (
	"fmt"
)

type Position struct {
	WP, WN, WB, WR, WQ, WK uint64
	BP, BN, BB, BR, BQ, BK uint64
}

func bitboardToArray(bb uint64) (board [64]uint8) {
	for i := 0; i < 64; i++ {
		board[i] = uint8((bb >> i) & 1)
	}
	return
}

func InitStartPos() Position {
	const (
		Rank1 uint64 = 0x00000000000000FF
		Rank2 uint64 = 0x000000000000FF00
		Rank7 uint64 = 0x00FF000000000000
		Rank8 uint64 = 0xFF00000000000000

		FileA uint64 = 0x0101010101010101
		FileB uint64 = 0x0202020202020202
		FileC uint64 = 0x0404040404040404
		FileD uint64 = 0x0808080808080808
		FileE uint64 = 0x1010101010101010
		FileF uint64 = 0x2020202020202020
		FileG uint64 = 0x4040404040404040
		FileH uint64 = 0x8080808080808080
	)
	return *&Position{}
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

	fmt.Printf("%b", (0x0101010101010101|0x8080808080808080)&0x00000000000000FF)
}
