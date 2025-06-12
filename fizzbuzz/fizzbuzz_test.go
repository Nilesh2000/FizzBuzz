package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzzNEquals3(t *testing.T) {
	n := 3
	want := []string{
		"1",
		"2",
		"Fizz",
	}

	got := FizzBuzz(n)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FizzBuzz(%d) = %v, want %v", n, got, want)
	}
}

func TestFizzBuzzNEquals5(t *testing.T) {
	n := 5
	want := []string{
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
	}

	got := FizzBuzz(n)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FizzBuzz(%d) = %v, want %v", n, got, want)
	}
}
