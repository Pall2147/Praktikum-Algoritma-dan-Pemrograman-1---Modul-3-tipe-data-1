package main

import "fmt"

func main() {
	var N, sisa, detik, menit, jam int

	fmt.Scan(&N)
	jam = N / 3600
	sisa = N % 3600
	menit = sisa / 60
	detik = sisa % 60

	fmt.Println(jam,menit,detik)
}