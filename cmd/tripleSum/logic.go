package main

import (
	"slices"
)

func ConvertToUnique(arr []int) []int {
	intMap := map[int]struct{}{}
	result := []int{}

	for _, val := range arr {
		intMap[val] = struct{}{}
	}

	for uniqueIndex := range intMap {
		result = append(result, uniqueIndex)
	}

	return result
}

func TripleSumLogic(arr1 []int, arr2 []int, arr3 []int) int {
	result := 0

	arr1 = ConvertToUnique(arr1)
	arr2 = ConvertToUnique(arr2)
	arr3 = ConvertToUnique(arr3)

	slices.Sort(arr1)
	slices.Sort(arr2)
	slices.Sort(arr3)

	for _, val := range arr2 {
		for leftIndex := 0; leftIndex < len(arr1); leftIndex += 1 {
			if arr1[leftIndex] > val {
				break
			} else {
				for rightIndex := 0; rightIndex < len(arr3); rightIndex += 1 {
					if arr3[rightIndex] > val {
						break
					} else {
						result += 1
					}
				}
			}
		}
	}

	return result
}

func OptimizedTripleSumLogic(arr1 []int, arr2 []int, arr3 []int) int {
	result := 0

	arr1 = ConvertToUnique(arr1)
	arr2 = ConvertToUnique(arr2)
	arr3 = ConvertToUnique(arr3)

	slices.Sort(arr1)
	slices.Sort(arr2)
	slices.Sort(arr3)

	for _, val := range arr2 {
		indexCountLeft, _ := slices.BinarySearch(arr1, val+1)
		indexCountRight, _ := slices.BinarySearch(arr3, val+1)
		result += (indexCountLeft * indexCountRight)
	}

	return result
}
