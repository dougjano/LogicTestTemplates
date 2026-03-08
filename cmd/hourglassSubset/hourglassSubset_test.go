package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	t.Run("Large Matrix", func(t *testing.T) {
		matrix := [][]int{
			{-9, -9, -9, 1, 1, 1},
			{0, -9, 0, 4, 3, 2},
			{-9, -9, -9, 1, 2, 3},
			{0, 0, 8, 6, 6, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}
		hourglasses := GetAllHourglasses(matrix)
		assert.Equal(t, 16, len(hourglasses))
	})

	t.Run("Small Matrix", func(t *testing.T) {
		matrix := [][]int{
			{-9, -9, -9},
			{0, -9, 0},
			{-9, -9, -9},
		}
		hourglasses := GetAllHourglasses(matrix)
		assert.Equal(t, 1, len(hourglasses))
		assert.Equal(t, []int{-9, -9, -9, -9, -9, -9, -9}, hourglasses[0])
	})

	t.Run("Sum Hourglass Values", func(t *testing.T) {
		matrix := [][]int{
			{-9, -9, -9, 1, 1, 1},
			{0, -9, 0, 4, 3, 2},
			{-9, -9, -9, 1, 2, 3},
			{0, 0, 8, 6, 6, 0},
			{0, 0, 0, -2, 0, 0},
			{0, 0, 1, 2, 4, 0},
		}
		hourglasses := GetAllHourglasses(matrix)
		firstHourglassSum := GetHourglassSum(hourglasses[0])
		assert.Equal(t, -63, firstHourglassSum)

		highestHourglassSum := GetHighestHourglassSum(hourglasses)
		assert.Equal(t, 28, highestHourglassSum)

		highestHourglassSum = OptimizedHighestHourglassSum(matrix)
		assert.Equal(t, 28, highestHourglassSum)

	})
}
