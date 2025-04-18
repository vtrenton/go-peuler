package main

import (
	"slices"
	"testing"
)

func TestHighestPrimeFactor(t *testing.T) {
	t.Run("getFactors returns a list of all the factors of a number up to the Sqrt", func(t *testing.T) {
		got := getFactors(20)
		want := []int{1, 2}

		if !slices.Equal(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("prime numbers should only return 1", func(t *testing.T) {

	})
}
