package main

import (
	"fmt"
	"math"
)

func main() {
	var currency string = "грн"
	var pieces string = "шт."
	const applePrice = 5.99
	const pearPrice = 7.00

	var balance = 23.00

	fmt.Println("Одне яблуко коштує", applePrice, currency)
	fmt.Println("Одна груша коштує", pearPrice, currency)
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?",  math.Round(float64(applePrice*9+float32(pearPrice)*8)), currency)
	fmt.Println("2. Скільки груш ми можемо купити?", math.Floor(balance/pearPrice), pieces)
	fmt.Println("3. Скільки яблук ми можемо купити?", float32(balance)/applePrice, pieces)
	fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука?", "Так")
}