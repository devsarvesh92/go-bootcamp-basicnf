package basicnf

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
