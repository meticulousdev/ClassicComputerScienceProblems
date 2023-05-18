package main

import (
	"fmt"
	"math/bits"
	"strings"
	"unsafe"
)

func compression(gene string) uint64 {
	var bit_string uint64 = 0b01

	for i := 0; i < len(gene); i++ {
		var nucleotide = string(gene[i])
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
			fmt.Printf("Invalid nucleotide: %s", nucleotide)
		}
	}

	return bit_string
}

func decompress(bit_string uint64) string {
	gene := ""
	cnt := 0
	for i := 0; i < bits.Len64(bit_string)-1; i = i + 2 {
		var bits = (bit_string >> i) & 0b11
		if bits == 0b00 {
			gene += "A"
		} else if bits == 0b01 {
			gene += "C"
		} else if bits == 0b10 {
			gene += "G"
		} else if bits == 0b11 {
			gene += "T"
		} else {
			fmt.Printf("Invalid bits: %b", bits)
		}
		cnt++
	}

	gene = reverse_string(gene)
	return gene
}

func reverse_string(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	gene := "AGCT"
	original := ""

	const LENGTH int = 40

	for i := 0; i < LENGTH; i++ {
		original += string(gene[i%4])
	}

	fmt.Printf("ORIGINAL        : %s\n", original)
	fmt.Printf("ORIGINAL SIZE   : %d BYTE\n\n", unsafe.Sizeof(original))

	bit_string := compression(original)
	fmt.Printf("COMPRESSED      : %d\n", bit_string)
	fmt.Printf("COMPRESSED SIZE : %d BYTE\n\n", unsafe.Sizeof(bit_string))

	gene_ret := decompress(bit_string)
	fmt.Printf("DECOMPRESSED    : %s\n", gene_ret)
}

// ORIGINAL        : AGCTAGCTAGCTAGCTAGCTAGCTAGCTAG
// ORIGINAL SIZE   : 16 BYTE

// COMPRESSED      : 1329250675899658866
// COMPRESSED SIZE : 8 BYTE

// DECOMPRESSED    : AGCTAGCTAGCTAGCTAGCTAGCTAGCTAG
