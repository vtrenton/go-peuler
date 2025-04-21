package main

import (
	"slices"
	"testing"
)

func TestGetPrimesToLimit(t *testing.T) {
	cases := []struct {
		Description string
		limit       int
		expected    []int
	}{
		{
			"get primes up to 100",
			100,
			[]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97},
		},
		{
			"return nothing if input number < 2",
			1,
			[]int{},
		},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := getprimestolimit(test.limit)

			if !slices.Equal(got, test.expected) {
				t.Errorf("got %d but wanted %d", got, test.expected)
			}
		})
	}
}

func TestIsPrime(t *testing.T) {
	cases := []struct {
		description string
		input       int
		result      bool
	}{
		{
			"test 7 a simple prime number",
			7,
			true,
		},
		{
			"test 1 which is not a prime number",
			1,
			false,
		},
	}

	for _, test := range cases {
		t.Run(test.description, func(t *testing.T) {
			got := isprime(test.input)
			want := test.result

			if got != want {
				t.Errorf("given %d, got %v but want %v", test.input, got, want)
			}
		})
	}
}
