package main

import (
	"slices"
)

func FixStringParentheses(value string) string {

	var stack []int
	var indexesToDelete []int

	for i, char := range value {
		switch char {
		case '(':
			stack = push(stack, i)
		case ')':
			if len(stack) == 0 {
				indexesToDelete = append(indexesToDelete, i)
			} else {
				stack = pop(stack)
			}
		}
	}

	indexesToDelete = append(indexesToDelete, stack...)

	slices.Sort(indexesToDelete)

	for i := len(indexesToDelete) - 1; i >= 0; i -= 1 {
		value = value[:indexesToDelete[i]] + value[indexesToDelete[i]+1:]
	}

	return value
}

func push(a []int, n int) []int {
	return append(a, n)
}

func pop(a []int) []int {
	return a[:len(a)-1]
}
