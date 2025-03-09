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
