package basic

import (
	"sort"
	"testing"
)

type goBasic struct{}

func (fk goBasic) Reduce(numbers []int, f func(acc, elem int) int, acc int) int {
	for _, v := range numbers {
		acc = f(acc, v)
	}
	return acc
}
func (fk goBasic) Map(numbers []int, f func(val int) int) []int {
	result := make([]int, len(numbers))
	for i, v := range numbers {
		result[i] = f(v)
	}
	return result
}

func (fk goBasic) Reverse(numbers []int) []int {
	result := make([]int, len(numbers))
	for i, v := range numbers {
		result[len(numbers)-1-i] = v
	}
	return result
}
func (fk goBasic) Filter(numbers []int, f func(val int) bool) []int {
	result := make([]int, 0, len(numbers))
	for _, v := range numbers {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}
func (fk goBasic) Sort(numbers []int, f func(i, j int) int) []int {
	result := append(make([]int, 0, len(numbers)), numbers...)
	sort.Sort(byFunc{Slice: result, Func: f})
	return result
}
func (fk goBasic) Contains(numbers []int, needle int) bool {
	for _, v := range numbers {
		if v == needle {
			return true
		}
	}
	return false
}

var basic goBasic

func BenchmarkGoBasicSumElements5000(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = int(basic.Reduce(numbers5000, sum, 0))
	}
	var expected = 0
	for _, val := range numbers5000 {
		expected += val
	}
	if expected != result {
		b.FailNow()
	}
}

func BenchmarkGoBasicDuplElements5000(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = basic.Map(numbers5000, func(val int) int {
			return val * 2
		})
	}
	for index := range numbers5000 {
		if numbers5000[index]*2 != result[index] {
			b.FailNow()
		}
	}
}

func BenchmarkGoBasicReverseElements5000(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = basic.Reverse(numbers5000)
	}
	for index := range numbers5000 {
		if numbers5000[index] != result[4999-index] {
			b.FailNow()
		}
	}
}

func BenchmarkGoBasicFilterElements5000(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = basic.Filter(numbers5000, func(val int) bool {
			return val%2 == 0
		})
	}
	for _, val := range result {
		if val%2 != 0 {
			b.FailNow()
		}
	}
}
func BenchmarkGoBasicSortElements5000(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = basic.Sort(numbers5000, func(i, j int) int {
			if i <= j {
				return -1
			}
			return 1
		})
	}
	for index := 0; index < 4999; index++ {
		if result[index] > result[index+1] {
			b.Logf("[%d]=%d > %d=[%d]", index, result[index], result[index+1], index+1)
			b.FailNow()
		}
	}
}

func BenchmarkGoBasicContainsElements5000(b *testing.B) {
	var result bool
	for i := 0; i < b.N; i++ {
		result = basic.Contains(numbers5000, numbers5000[4999])
	}
	if !result {
		b.FailNow()
	}
}

type byFunc struct {
	Slice []int
	Func  func(i, j int) int
}

func (by byFunc) Len() int           { return len(by.Slice) }
func (by byFunc) Swap(i, j int)      { by.Slice[i], by.Slice[j] = by.Slice[j], by.Slice[i] }
func (by byFunc) Less(i, j int) bool { return by.Func(by.Slice[i], by.Slice[j]) < 0 }
