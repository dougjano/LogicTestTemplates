package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	//just one method
	t.Run("Triplet Sum Test", func(t *testing.T) {
		assert.Equal(t, 4, TripleSumLogic([]int{3, 5, 7}, []int{3, 6}, []int{4, 6, 9}))
		assert.Equal(t, 8, TripleSumLogic([]int{3, 2, 5}, []int{2, 3}, []int{1, 2, 3}))
		assert.Equal(t, 5, TripleSumLogic([]int{1, 4, 5}, []int{2, 3, 3}, []int{1, 2, 3}))
		assert.Equal(t, 12, TripleSumLogic([]int{1, 3, 5, 7}, []int{5, 7, 9}, []int{7, 9, 11, 13}))
	})

	t.Run("Optimized Triplet Sum Test", func(t *testing.T) {
		assert.Equal(t, 4, OptimizedTripleSumLogic([]int{3, 5, 7}, []int{3, 6}, []int{4, 6, 9}))
		assert.Equal(t, 8, OptimizedTripleSumLogic([]int{3, 2, 5}, []int{2, 3}, []int{1, 2, 3}))
		assert.Equal(t, 5, OptimizedTripleSumLogic([]int{1, 4, 5}, []int{2, 3, 3}, []int{1, 2, 3}))
		assert.Equal(t, 12, OptimizedTripleSumLogic([]int{1, 3, 5, 7}, []int{5, 7, 9}, []int{7, 9, 11, 13}))
	})
}
