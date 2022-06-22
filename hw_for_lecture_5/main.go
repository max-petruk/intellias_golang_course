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
		applePrice = 5.99
		pearPrice = 7.00
	)

	fmt.Printf("Одне яблуко коштує - %v грн.\n", applePrice)
	fmt.Printf("Одна груша коштує - %v грн.\n", pearPrice)
	fmt.Printf("1. Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш? - %v грн.\n",  applePrice*float64(applesToBuy)+pearPrice*float64(pearsToBuy))
	fmt.Printf("2. Скільки груш ми можемо купити? - %v шт.\n", math.Floor(balance/pearPrice))
	fmt.Printf("3. Скільки яблук ми можемо купити? -%v шт.\n", math.Floor(balance/applePrice))
	
	applesToBuy = 2
	pearsToBuy = 2

	if (applePrice*float64(applesToBuy)+pearPrice*float64(pearsToBuy) <= balance) {
		fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука? - Так")
	} else {
		fmt.Println("4. Чи ми можемо купити 2 груші та 2 яблука? - Ні")
	}
}