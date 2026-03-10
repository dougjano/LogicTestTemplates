package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {

	/*t.Run("Exponent Test", func(t *testing.T) {
		assert.Equal(t, 3, GetExpoent(125, 5))
		assert.Equal(t, -1, GetExpoent(123, 5))
		assert.Equal(t, 2, GetExpoent(16, 4))
	})

	t.Run("Map Test", func(t *testing.T) {
		mapCounter, _ := UpsertMapCounter([]int{1, 4, 16, 64}, 4)
		assert.Equal(t, map[int]int{0: 1, 1: 1, 2: 1, 3: 1}, mapCounter)
		mapCounter, _ = UpsertMapCounter([]int{1, 5, 5, 25, 125}, 5)
		assert.Equal(t, map[int]int{0: 1, 1: 2, 2: 1, 3: 1}, mapCounter)
	})

	t.Run("Triplet Count Test", func(t *testing.T) {
		assert.Equal(t, 2, CheckTripletsCount([]int{1, 4, 16, 64}, 4))
		assert.Equal(t, 4, CheckTripletsCount([]int{1, 5, 5, 25, 125}, 5))
		assert.Equal(t, 6, CheckTripletsCount([]int{1, 3, 9, 9, 27, 81}, 3))
		assert.Equal(t, 8, CheckTripletsCount([]int{1, 1, 4, 4, 16, 16}, 4))
	})*/

	t.Run("Optimized Triplet Count Test", func(t *testing.T) {
		assert.Equal(t, 2, OptimizedTripletsCount([]int{1, 4, 16, 64}, 4))
		assert.Equal(t, 4, OptimizedTripletsCount([]int{1, 5, 5, 25, 125}, 5))
		assert.Equal(t, 6, OptimizedTripletsCount([]int{1, 3, 9, 9, 27, 81}, 3))
		assert.Equal(t, 8, OptimizedTripletsCount([]int{1, 1, 4, 4, 16, 16}, 4))
	})
}
