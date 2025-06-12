package fizzbuzz

import "strconv"

const (
	FIZZ     = "Fizz"
	BUZZ     = "Buzz"
	FIZZBUZZ = "FizzBuzz"
)

func FizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result[i-1] = FIZZBUZZ
		} else if i%3 == 0 {
			result[i-1] = FIZZ
		} else if i%5 == 0 {
			result[i-1] = BUZZ
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}
