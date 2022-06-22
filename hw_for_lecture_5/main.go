package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		applesToBuy = 9
		pearsToBuy = 8
		balance = 23.00
	)

	const (
		// currency string = "грн"
		// pieces string = "шт."
		applePrice = 5.99
		pearPrice = 7.00
	)

	// fmt.Println("Одне яблуко коштує", applePrice, currency)
	fmt.Printf("Одне яблуко коштує - %v грн.\n", applePrice)
	// fmt.Println("Одна груша коштує", pearPrice, currency)
	fmt.Printf("Одна груша коштує - %v грн.\n", pearPrice)

	// fmt.Println("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?",  applePrice*float64(applesToBuy)+pearPrice*float64(pearsToBuy), currency)
	fmt.Printf("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш? - %v грн.\n",  applePrice*float64(applesToBuy)+pearPrice*float64(pearsToBuy))


	
	// fmt.Println("2. Скільки груш ми можемо купити?", math.Floor(balance/pearPrice), pieces)
	fmt.Printf("2. Скільки груш ми можемо купити? - %v шт.\n", math.Floor(balance/pearPrice))
	// fmt.Println("3. Скільки яблук ми можемо купити?", math.Floor(balance/applePrice), pieces)
	fmt.Printf("3. Скільки яблук ми можемо купити? -%v шт.\n", math.Floor(balance/applePrice))
	applesToBuy = 2
	pearsToBuy = 2

	// 4. Чи ми можемо купити 2 груші та 2 яблука?

	if (applePrice*float64(applesToBuy)+pearPrice*float64(pearsToBuy) <= balance) {
		fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука? - Так")
	} else {
		fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука? - Ні")
	}
}