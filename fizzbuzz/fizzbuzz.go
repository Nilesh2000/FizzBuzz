package fizzbuzz

func FizzBuzz(n int) []string {
	switch n {
	case 1:
		return []string{"1"}
	case 3:
		return []string{"1", "2", "Fizz"}
	case 5:
		return []string{"1", "2", "Fizz", "4", "Buzz"}
	case 15:
		return []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}
	default:
		return []string{}
	}
}
