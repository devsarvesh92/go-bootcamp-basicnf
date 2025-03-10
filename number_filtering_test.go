package basicnf_test

import (
	basicnf "basic-nf"
	"reflect"
	"testing"
)

func TestGetEvenNumbers(t *testing.T) {
	test := []struct {
		nums     []int
		expected []int
	}{{
		nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		expected: []int{2, 4, 6, 8, 10},
	}}
	for _, x := range test {
		got := basicnf.GetEvenNumbers(x.nums)
		if !reflect.DeepEqual(x.expected, got) {
			t.Errorf("Vals not matching")
		}
	}
}

func TestGetOddNumbers(t *testing.T) {
	test := []struct {
		nums     []int
		expected []int
	}{{
		nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		expected: []int{1, 3, 5, 7, 9},
	}}
	for _, x := range test {
		got := basicnf.GetOddNumbers(x.nums)
		if !reflect.DeepEqual(x.expected, got) {
			t.Errorf("Vals not matching")
		}
	}
}

func TestGetPrimeNumbers(t *testing.T) {
	test := []struct {
		nums     []int
		expected []int
	}{{
		nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		expected: []int{2, 3, 5, 7},
	}}
	for _, x := range test {
		got := basicnf.GetPrimeNumbers(x.nums)
		if !reflect.DeepEqual(x.expected, got) {
			t.Errorf("Vals not matching")
		}
	}
}

func TestGetOddPrimeNumbers(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, expected: []int{3, 5, 7}}}

	for _, val := range tests {
		got := basicnf.All(val.input, basicnf.IsPrime, basicnf.IsOdd)
		if !reflect.DeepEqual(got, val.expected) {
			t.Errorf("Got %v, Expected %v", got, val.expected)
		}
	}
}

func TestEvenMultiplesOf5(t *testing.T) {
	MultipleOf5 := func(num int) bool {
		return num%5 == 0
	}

	tests := []struct {
		input    []int
		expected []int
	}{{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, expected: []int{10, 20}}}

	for _, val := range tests {
		got := basicnf.All(val.input, MultipleOf5, basicnf.IsEven)
		if !reflect.DeepEqual(got, val.expected) {
			t.Errorf("Got %v, Expected %v", got, val.expected)
		}
	}
}

func TestOddMultiplesOf3Gt10(t *testing.T) {
	MultipleOf3 := func(num int) bool {
		return num%3 == 0
	}
	Gt10 := func(num int) bool {
		return num > 10
	}

	tests := []struct {
		input    []int
		expected []int
	}{{input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}, expected: []int{15}}}

	for _, val := range tests {
		got := basicnf.All(val.input, MultipleOf3, Gt10, basicnf.IsOdd)
		if !reflect.DeepEqual(got, val.expected) {
			t.Errorf("Got %v, Expected %v", got, val.expected)
		}
	}
}
