package basic

import (
	"fmt"
	"github.com/wesovilabs/koazee"
	"github.com/wesovilabs/koazee/stream"
	"strings"
	"testing"
)

func BenchmarkKoazeeSumElements5000(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = koazee.StreamOf(numbers5000).Reduce(func(acc int, val int) int { return acc + val }).Int()
	}
	var expected = 0
	for _, val := range numbers5000 {
		expected += val
	}
	if expected != result {
		b.FailNow()
	}
}

func BenchmarkKoazeeDuplElements5000(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = koazee.StreamOf(numbers5000).Map(func(val int) int { return val * 2 }).Do().Out().Val().([]int)
	}
	for index := range numbers5000 {
		if numbers5000[index]*2 != result[index] {
			b.FailNow()
		}
	}
}

func BenchmarkKoazeeReverseElements5000(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = koazee.StreamOf(numbers5000).Reverse().Do().Out().Val().([]int)
	}
	for index := range numbers5000 {
		if numbers5000[index] != result[4999-index] {
			b.FailNow()
		}
	}
}

func BenchmarkKoazeeSortElements5000(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = koazee.StreamOf(numbers5000).Sort(func(i, j int) int {
			if i <= j {
				return -1
			}
			return 1
		}).Do().Out().Val().([]int)
	}
	for index := 0; index < 4999; index++ {
		if result[index] > result[index+1] {
			b.FailNow()
		}
	}
}

func BenchmarkKoazeeFilterElements5000(b *testing.B) {
	var result []int
	for i := 0; i < b.N; i++ {
		result = koazee.StreamOf(numbers5000).Filter(func(val int) bool {
			return val%2 == 0
		}).Do().Out().Val().([]int)
	}
	for _, val := range result {
		if val%2 != 0 {
			b.FailNow()
		}
	}
}

func BenchmarkKoazeeContainsElements5000(b *testing.B) {
	var result bool
	for i := 0; i < b.N; i++ {
		result, _ = koazee.StreamOf(numbers5000).Contains(numbers5000[4999])
	}
	if !result {
		b.FailNow()
	}
}

func BenchmarkKoazeeOperationgWithStrings50008b(b *testing.B) {
	var result stream.Output
	for i := 0; i < b.N; i++ {
		result = koazee.StreamOf(strings5000).
			Sort(strings.Compare).
			Reverse().
			Filter(func(val string) bool {
				return len(val) == 4
			}).
			Map(func(val string) int {
				return len(val)
			}).
			Reduce(func(acc, val int) int {
				return acc + val
			})
	}
	var expected int
	for _, val := range strings5000 {
		if len(val) == 4 {
			expected += 4
		}
	}
	if result.Err() != nil {
		fmt.Println(result.Err())
	}
	if result.Int() != expected {
		b.FailNow()
	}
}
