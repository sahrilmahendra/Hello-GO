package main

import "fmt"

func main() {
	// 1. luas segitiga
	// input
	var (
		alas   float64 = 20
		tinggi float64 = 25
	)

	// kode
	luas := (alas * tinggi) / 2

	// output
	fmt.Println(luas)
}
