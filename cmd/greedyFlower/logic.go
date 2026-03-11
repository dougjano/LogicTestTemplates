package main

import (
	"slices"
)

func GetMinPrice(k int, arr []int) int {
	result := 0
	personIndex := 1

	slices.Sort(arr)
	slices.Reverse(arr)

	for i, price := range arr {
		if i > 0 && i%k == 0 {
			personIndex += 1
		}

		result += price * personIndex
	}

	return int(result)
}
