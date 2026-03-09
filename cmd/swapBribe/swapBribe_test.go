package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {

	t.Run("Bubble Sort Test", func(t *testing.T) {

		inputQ := []int{1, 2, 3, 4, 5}

		sortedQ, hasChanged := BubbleSort(inputQ)

		assert.Equal(t, sortedQ, inputQ)
		assert.False(t, hasChanged)

		inputQ = []int{1, 2, 3, 5, 4}

		sortedQ, hasChanged = BubbleSort(inputQ)

		assert.Equal(t, sortedQ, []int{1, 2, 3, 4, 5})
		assert.True(t, hasChanged)
	})

	t.Run("Queue Swap Logic Test", func(t *testing.T) {

		inputQ := []int{1, 2, 3, 5, 4, 6, 7, 8}

		assert.Equal(t, BubbleSortCounter(inputQ), 1)

		inputQ = []int{2, 1, 5, 3, 4}

		assert.Equal(t, BubbleSortCounter(inputQ), 2)

		inputQ = []int{2, 5, 1, 3, 4}

		assert.Equal(t, BubbleSortCounter(inputQ), 3)

		inputQ = []int{5, 1, 2, 3, 7, 8, 6, 4}

		assert.Equal(t, BubbleSortCounter(inputQ), 5)
	})
}
