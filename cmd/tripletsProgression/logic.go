package main

import (
	"fmt"
)

func UpsertMapCounter(values []int, ratio int) (map[int]int, int) {
	result := map[int]int{}
	maxExpoent := 0

	for _, val := range values {
		expoent := GetExpoent(val, ratio)

		if expoent != -1 {
			if expoent > maxExpoent {
				maxExpoent = expoent
			}

			result[expoent] += 1
		}
	}

	return result, maxExpoent
}

func GetExpoent(value int, ratio int) int {
	result := 0

	if value%ratio != 0 && value != 1 {
		return -1
	}

	for value != 1 {
		result += 1
		value = value / ratio
	}

	return result
}

func CheckTripletsCount(values []int, ratio int) int {

	result := 0
	myMap, maxExpoent := UpsertMapCounter(values, ratio)

	for i := maxExpoent; i > 1; i -= 1 {
		tripletCounter := 0

		if tripletCounter < myMap[i] {
			tripletCounter = myMap[i]
		}

		if tripletCounter < myMap[i-1] {
			tripletCounter = myMap[i-1]
		}

		if tripletCounter < myMap[i-2] {
			tripletCounter = myMap[i-2]
		}

		result += tripletCounter
	}

	return result
}

func OptimizedTripletsCount(values []int, ratio int) int {
	result := 0
	mapCounter := map[int]int{}
	mapPairCounter := map[string]int{}

	for _, val := range values {
		prev := val / ratio
		prevPrev := prev / ratio

		if val%ratio == 0 && prev%ratio == 0 {
			pairKey := fmt.Sprintf("%d,%d", prevPrev, prev)
			result += mapPairCounter[pairKey]
		}

		if val%ratio == 0 {
			pairKey := fmt.Sprintf("%d,%d", prev, val)
			mapPairCounter[pairKey] += mapCounter[prev]
		}

		mapCounter[val]++
	}

	return result
}
