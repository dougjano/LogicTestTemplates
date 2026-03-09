package main

import (
	"fmt"
)

func BubbleSortCounter(vector []int) int {
	hasChanged := false
	iterationCounter := 0

	for i := 0; i < (len(vector) * len(vector)); i += 1 {
		vector, hasChanged = BubbleSort(vector)

		if !hasChanged {
			iterationCounter = i
			break
		}
	}

	if iterationCounter > 2 {
		fmt.Println("Too chaotic")
	}

	return iterationCounter
}

func BubbleSort(vector []int) ([]int, bool) {

	tempValue := 0
	hasChanged := false

	for i := 0; i < len(vector); i += 1 {
		if i != len(vector)-1 {
			if vector[i] > vector[i+1] {
				tempValue = vector[i]
				vector[i] = vector[i+1]
				vector[i+1] = tempValue

				i += 1
				hasChanged = true
			}
		}
	}

	return vector, hasChanged
}
