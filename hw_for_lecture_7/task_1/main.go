package main

import (
	"fmt"
)

func main() {
	arr := []int{4, 1, 4, -4, 6, 3, 8, 8}
	var result []int
	uniqueNumbers := make(map[int]bool)

	for _, i := range arr {
		if _, value := uniqueNumbers[i]; !value {
			uniqueNumbers[i] = true
			result = append(result, i)
		}
	}
	fmt.Printf("Here is the result = %v\n", result)
}