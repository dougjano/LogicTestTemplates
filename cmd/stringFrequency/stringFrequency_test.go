package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	t.Run("Test Phrases", func(t *testing.T) {
		assert.True(t, CheckPhrases([]string{"give", "me", "one", "grand", "today", "night"}, []string{"give", "one", "grand", "today"}))
	})
}
