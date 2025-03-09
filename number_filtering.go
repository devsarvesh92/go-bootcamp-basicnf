package basicnf

import "math"

func GetEvenNumbers(numbers []int) []int {
	evenNumbers := make([]int, 0, len(numbers))
	for _, num := range numbers {
		if num%2 == 0 {
			evenNumbers = append(evenNumbers, num)
		}
	}
	return evenNumbers
}

func GetOddNumbers(numbers []int) []int {
	oddNumbers := make([]int, 0, len(numbers))
	for _, val := range numbers {
		if val%2 != 0 {
			oddNumbers = append(oddNumbers, val)
		}
	}
	return oddNumbers
}

func GetPrimeNumbers(numbers []int) []int {
	primeNumbers := make([]int, 0, len(numbers))

	for _, num := range numbers {
		if num <= 1 {
			continue
		}

		isPrime := true
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 {
				isPrime = false
			}
		}

		if isPrime {
			primeNumbers = append(primeNumbers, num)
		}
	}
	return primeNumbers
}
