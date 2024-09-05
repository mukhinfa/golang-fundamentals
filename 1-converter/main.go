package main

import "fmt"

func main() {
	const fromUSDtoEUR = 0.9
	const fromUSDtoRUB = 88.25
	fromEURtoRUB := fromUSDtoEUR / fromUSDtoRUB
	fmt.Printf("EUR в RUB: %.2F", fromEURtoRUB)
}
