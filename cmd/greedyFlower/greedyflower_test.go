package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	t.Run("Get Min Price Test", func(t *testing.T) {
		assert.Equal(t, 11, GetMinPrice(3, []int{1, 2, 3, 4}))
		assert.Equal(t, 13, GetMinPrice(3, []int{2, 5, 6}))
		assert.Equal(t, 15, GetMinPrice(2, []int{2, 5, 6}))
		assert.Equal(t, 29, GetMinPrice(3, []int{1, 3, 5, 7, 9}))
	})
}
