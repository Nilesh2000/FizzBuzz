package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzzNEquals1(t *testing.T) {
	n := 1
	want := []string{
		"1",
	}

	got := FizzBuzz(n)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FizzBuzz(%d) = %v, want %v", n, got, want)
	}
}

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

func TestFizzBuzzNEquals15(t *testing.T) {
	n := 15
	want := []string{
		"1",
		"2",
		"Fizz",
		"4",
		"Buzz",
		"Fizz",
		"7",
		"8",
		"Fizz",
		"Buzz",
		"11",
		"Fizz",
		"13",
		"14",
		"FizzBuzz",
	}

	got := FizzBuzz(n)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FizzBuzz(%d) = %v, want %v", n, got, want)
	}
}
