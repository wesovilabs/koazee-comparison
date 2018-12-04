package basic

import (
	"github.com/ahmetb/go-linq"
	"strings"
	"testing"
)

func BenchmarkGoLinqSumElements5000(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = int(linq.From(numbers5000).SumInts())
	}
	var expected = 0
	for _, val := range numbers5000 {
		expected += val
	}
	if expected != result {
		b.FailNow()
	}
}

func BenchmarkGoLinqDuplElements5000(b *testing.B) {
	var result []interface{}
	for i := 0; i < b.N; i++ {
		result = linq.From(numbers5000).Select(func(val interface{}) interface{} {
			return val.(int) * 2
		}).Results()
	}
	for index := range numbers5000 {
		if numbers5000[index]*2 != result[index] {
			b.FailNow()
		}
	}
}

func BenchmarkGoLinqReverseElements5000(b *testing.B) {
	var result []interface{}
	for i := 0; i < b.N; i++ {
		result = linq.From(numbers5000).Reverse().Results()
	}
	for index := range numbers5000 {
		if numbers5000[index] != result[4999-index] {
			b.FailNow()
		}
	}
}

func BenchmarkGoLinqSortElements5000(b *testing.B) {
	var result []interface{}
	for i := 0; i < b.N; i++ {
		result = linq.From(numbers5000).Sort(func(i, j interface{}) bool {
			return i.(int) <= j.(int)
		}).Results()
	}
	for index := 0; index < 4999; index++ {
		if result[index].(int) > result[index+1].(int) {
			b.FailNow()
		}
	}
}

func BenchmarkGoLinqFilterElements5000(b *testing.B) {
	var result []interface{}
	for i := 0; i < b.N; i++ {
		result = linq.From(numbers5000).Where(func(val interface{}) bool {
			return val.(int)%2 == 0
		}).Results()
	}
	for _, val := range result {
		if val.(int)%2 != 0 {
			b.FailNow()
		}
	}
}

func BenchmarkGoLinqContainsElements5000(b *testing.B) {
	var result bool
	for i := 0; i < b.N; i++ {
		result = linq.From(numbers5000).Contains(numbers5000[4999])
	}
	if !result {
		b.FailNow()
	}
}

func BenchmarkGoLinqOperationgWithStrings50008b(b *testing.B) {
	var result int
	for i := 0; i < b.N; i++ {
		result = int(linq.From(strings5000).
			Sort(func(i, j interface{}) bool {
				return strings.Compare(i.(string), j.(string)) <= 0
			}).
			Reverse().
			Where(func(val interface{}) bool {
				return len(val.(string)) == 4
			}).
			Select(func(val interface{}) interface{} {
				return len(val.(string))
			}).
			SumInts())
	}
	var expected int
	for _, val := range strings5000 {
		if len(val) == 4 {
			expected += 4
		}
	}
	if result != expected {
		b.FailNow()
	}
}
