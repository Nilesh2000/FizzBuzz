package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{1, []string{"1"}},
		{3, []string{"1", "2", "Fizz"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}

	for _, test := range tests {
		got := FizzBuzz(test.n)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("FizzBuzz(%d) = %v, want %v", test.n, got, test.want)
		}
	}
}
