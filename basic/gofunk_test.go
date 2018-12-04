package basic

import (
	"github.com/thoas/go-funk"
	"testing"
)

func BenchmarkGoFunkSumElements5000(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = int(funk.Reduce(numbers5000, sum, 0))
	}
	var expected = 0
	for _, val := range numbers5000 {
		expected += val
	}
	if expected != result {
		b.FailNow()
	}
}

func BenchmarkGoFunkDuplElements5000(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = funk.Map(numbers5000, func(val int) int {
			return val * 2
		}).([]int)
	}
	for index := range numbers5000 {
		if numbers5000[index]*2 != result[index] {
			b.FailNow()
		}
	}
}

func BenchmarkGoFunkReverseElements5000(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = funk.Reverse(numbers5000).([]int)
	}
	for index := range numbers5000 {
		if numbers5000[index] != result[4999-index] {
			b.FailNow()
		}
	}
}

func BenchmarkGoFunkFilterElements5000(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = funk.Filter(numbers5000, func(val int) bool {
			return val%2 == 0
		}).([]int)
	}
	for _, val := range result {
		if val%2 != 0 {
			b.FailNow()
		}
	}
}

func BenchmarkGoFunkContainsElements5000(b *testing.B) {
	var result bool
	for i := 0; i < b.N; i++ {
		result = funk.Contains(numbers5000, numbers5000[4999])
	}
	if !result {
		b.FailNow()
	}
}
