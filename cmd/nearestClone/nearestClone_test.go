package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {

	t.Run("Search Test", func(t *testing.T) {
		assert.Equal(t, 1, FindNeighbours(1, []int{1, 1, 4}, []int{2, 3, 2}, []int{1, 2, 1, 1}))
		assert.Equal(t, -1, FindNeighbours(2, []int{1, 1, 4}, []int{2, 3, 2}, []int{1, 2, 3, 4}))
		assert.Equal(t, 3, FindNeighbours(2, []int{1, 1, 2, 3}, []int{2, 3, 4, 5}, []int{1, 2, 3, 3, 2}))
	})

}
