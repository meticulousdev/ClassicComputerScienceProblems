package main

import (
	"fmt"
	"strings"
)

func compression(gene []string) uint {
	var bit_string uint = 0b01

	for _, nucleotide := range gene {
		bit_string <<= 2
		if strings.Compare(nucleotide, "A") == 0 {
			bit_string |= 0b00
		} else if strings.Compare(nucleotide, "C") == 0 {
			bit_string |= 0b01
		} else if strings.Compare(nucleotide, "G") == 0 {
			bit_string |= 0b10
		} else if strings.Compare(nucleotide, "T") == 0 {
			bit_string |= 0b11
		} else {
			fmt.Printf("Invalid nucleotide:%s", nucleotide)
		}
	}

	return bit_string
}

// func decompress(bit_string uint) []string {
// 	gene := []string{}
// 	for i := 0; i < bits.Len(bit_string); i++ {

// 	}
// }

func main() {
	const LENGTH int = 10
	gene := []string{"A", "G", "C", "T"}
	original := [LENGTH]string{}
	for i := 0; i < LENGTH; i++ {
		original[i] = gene[i%4]
	}
	fmt.Printf("%s", original)
	fmt.Println()

	var bit_string = compression(gene)
	fmt.Printf("%b \n", bit_string)
}
