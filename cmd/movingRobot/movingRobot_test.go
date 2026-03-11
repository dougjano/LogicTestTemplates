package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	t.Run("Test Arrays", func(t *testing.T) {
		assert.True(t, CheckMovement([][]int{{1, 0, 0, 1}, {0, 1, 1, 0}}))
		assert.True(t, CheckMovement([][]int{{1, 0, 0, 0, 1}, {1, 0, 0, 1, 0}}))
		assert.True(t, CheckMovement([][]int{
			{1, 0, 0, 1},
			{0, 1, 1, 0},
			{1, 0, 0, 1},
		}))
	})
}
