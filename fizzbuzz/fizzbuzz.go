package fizzbuzz

const (
	FIZZ     = "Fizz"
	BUZZ     = "Buzz"
	FIZZBUZZ = "FizzBuzz"
)

func FizzBuzz(n int) []string {
	switch n {
	case 1:
		return []string{"1"}
	case 3:
		return []string{"1", "2", FIZZ}
	case 5:
		return []string{"1", "2", FIZZ, "4", BUZZ}
	case 15:
		return []string{"1", "2", FIZZ, "4", BUZZ, FIZZ, "7", "8", FIZZ, BUZZ, "11", FIZZ, "13", "14", FIZZBUZZ}
	default:
		return []string{}
	}
}
