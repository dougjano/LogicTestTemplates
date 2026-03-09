package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {

	t.Run("Stack Test", func(t *testing.T) {

		stack := []int{1, 2, 3}

		stack = push(stack, 4)

		assert.Equal(t, []int{1, 2, 3, 4}, stack)

		stack = pop(stack)
		stack = pop(stack)

		assert.Equal(t, []int{1, 2}, stack)

	})

	t.Run("String Test", func(t *testing.T) {

		testString := FixStringParentheses("abc(def(g)))hi(j(")

		assert.Equal(t, "abc(def(g))hij", testString)

		testString = FixStringParentheses("())))((())a(b))c(((d))")

		assert.Equal(t, "()((())a(b))c((d))", testString)

	})

}
