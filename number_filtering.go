package basicnf

import "math"

func GetEvenNumbers(numbers []int) []int {
	evenNumbers := make([]int, 0, len(numbers))
	for _, num := range numbers {
		if IsEven(num) {
			evenNumbers = append(evenNumbers, num)
		}
	}
	return evenNumbers
}

func GetOddNumbers(numbers []int) []int {
	oddNumbers := make([]int, 0, len(numbers))
	for _, val := range numbers {
		if IsOdd(val) {
			oddNumbers = append(oddNumbers, val)
		}
	}
	return oddNumbers
}

func GetPrimeNumbers(numbers []int) []int {
	primeNumbers := make([]int, 0, len(numbers))

	for _, num := range numbers {
		isPrime := IsPrime(num)
		if isPrime {
			primeNumbers = append(primeNumbers, num)
		}
	}
	return primeNumbers
}

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}

	isPrime := true
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			isPrime = false
		}
	}
	return isPrime
}

func IsOdd(num int) bool {
	return num%2 != 0
}

func IsEven(num int) bool {
	return num%2 == 0
}

type FilterCondition func(n int) bool

func All(numbers []int, condtions ...FilterCondition) []int {
	results := make([]int, 0, len(numbers))
	for _, num := range numbers {
		conditionMatch := true
		for _, condition := range condtions {
			if !condition(num) {
				conditionMatch = false
			}
		}
		if conditionMatch {
			results = append(results, num)
		}
	}
	return results
}
