package main

import (
	"fmt"
	"math"
)

func main() {
	var currency string = "грн"
	var pieces string = "шт."
	var applesToBuy = 9
	var pearsToBuy = 8
	const applePrice = 5.99
	const pearPrice = 7.00

	var balance = 23.00

	fmt.Println("Одне яблуко коштує", applePrice, currency)
	fmt.Println("Одна груша коштує", pearPrice, currency)
	fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?",  applePrice*float64(applesToBuy)+pearPrice*float64(pearsToBuy), currency)
	fmt.Println("2. Скільки груш ми можемо купити?", math.Floor(balance/pearPrice), pieces)
	fmt.Println("3. Скільки яблук ми можемо купити?", math.Floor(balance/applePrice), pieces)
	applesToBuy = 2
	pearsToBuy = 2

	if (applePrice*float64(applesToBuy)+pearPrice*float64(pearsToBuy) <= balance) {
		fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука? - Так")
	} else {
		fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука? - Ні")
	}
}