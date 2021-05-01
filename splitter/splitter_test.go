package splitter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplit(t *testing.T) {
	t.Run("no_space", func(t *testing.T) {
		result := Split("foobar", 3)
		expected := []string{"foo", "bar"}

		assert.Equal(t, result, expected)
	})

	t.Run("correct", func(t *testing.T) {
		result := Split("foo bar", 3)
		expected := []string{"foo", "bar"}

		assert.Equal(t, result, expected)
	})

	t.Run("trim_spaces", func(t *testing.T) {
		result := Split("foo    bar", 3)
		expected := []string{"foo", "bar"}

		assert.Equal(t, result, expected)
	})
}
