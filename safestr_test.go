package safestr

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestMask(t *testing.T) {
	card := "1234321232456789"
	expected := "************6789"

	t.Run("should return the parsed string", func(t *testing.T) {
		value, err := Parse(card, "*", 4)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		assert.Equal(t, expected, *value)
	})

	t.Run("should return an error of wrong length", func(t *testing.T) {
		_, err := Parse(card, "*", 17)
		assert.NotNil(t, err)
	})

	t.Run("should apply the default visible value when zero is passed", func(t *testing.T) {
		value, err := Parse(card, "*", 0)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		assert.Equal(t, expected, *value)
	})
}