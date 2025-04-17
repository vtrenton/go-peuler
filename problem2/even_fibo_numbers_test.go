package main

import (
	"slices"
	"testing"
)

func TestFiboNumbers(t *testing.T) {
	t.Run("fibogen generates all the correct finonacci numbers up to the limit", func(t *testing.T) {
		got := fibogen(13)
		want := []int{1, 1, 2, 3, 5, 8, 13}

		if !slices.Equal(got, want) {
			t.Errorf("got %d but wanted %d", got, want)
		}
	})

	t.Run("fibeven will only return even numbers from given input", func(t *testing.T) {
		input := []int{1, 1, 2, 3, 5, 8, 13, 21, 34}

		got := fibeven(input)
		want := []int{2, 8, 34}

		if !slices.Equal(got, want) {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("sumfibeven will output the correct sum", func(t *testing.T) {
		input := []int{2, 8, 34}

		got := sumfibeven(input)
		want := 44

		// sums of evens will always be even
		if got%2 != 0 {
			t.Errorf("result was not an even number: %d", got)
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
