package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	t.Run("Tree Creation Test", func(t *testing.T) {
		rootTree := BuildTree([]int{10, 5, 15, 3, 7, 13, 18})
		val := rootTree.rightChild.leftChild.value
		assert.Equal(t, 13, val)

		rootTree = BuildTree([]int{3, 0, 4, -1, 2, -1, -1, 1})
		val = rootTree.leftChild.rightChild.value
		assert.Equal(t, 2, val)
	})

	t.Run("Tree Trim Test", func(t *testing.T) {
		rootTree := BuildTree([]int{10, 5, 15, 3, 7, 13, 18})
		trimedTree := TreeTrim([]int{1, 3}, rootTree)
		nilNode := trimedTree.leftChild
		assert.Nil(t, nilNode)
	})
}
