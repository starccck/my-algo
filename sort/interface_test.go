package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

const testCaseCnt = 40

type testArg[T ~int] struct {
	data []T
}

type SortTestCase[T ~int] struct {
	name   string
	Sorter SortI[T]
	args   testArg[T]
}

func TestSortI(t *testing.T) {
	sortHelper(t, "quick_sort", SortI[MyIntType](IQS{}))
}

func sortHelper[T ~int](t *testing.T, name string, sorter SortI[T]) {
	tests := []SortTestCase[T]{}

	for i := 1; i <= testCaseCnt; i++ {
		slice := generate[T]()

		tests = append(tests, SortTestCase[T]{
			name: fmt.Sprintf("%s_%d", name, i),
			args: testArg[T]{
				data: slice,
			},
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newSlice := append([]T{}, tt.args.data...)
			// my sort interface
			t1 := time.Now()
			sorter.Sort(tt.args.data)
			sorterInterval := time.Now().Sub(t1)

			stdSorted, stdInterval := stdSortHelper(newSlice)

			if !compare(tt.args.data, stdSorted) {
				t.Errorf("sort is unavailable, name: %s, data: %#v", tt.name, tt.args.data)
				t.FailNow()
			}

			// t.Logf("test case passed, name: %s, data: %v\n", tt.name, tt.args.data)
			t.Logf("test case passed, name: %s, use time(microseconds), sorter: %d, std: %d\n", tt.name, sorterInterval.Microseconds(), stdInterval.Microseconds())
		})
	}
}

func compare[T1 ~int, T2 ~int](a []T1, b []T2) (euqal bool) {
	if len(a) != len(b) {
		return
	}

	for i := 0; i < len(a); i++ {
		if int(a[i]) != int(b[i]) {
			return
		}
	}

	return true
}

func generate[T ~int]() (slice []T) {
	cnt := rand.Int()%10000 + 10000
	slice = make([]T, 0, cnt)
	for i := 0; i < cnt; i++ {
		slice = append(slice, T(rand.Int()%1000))
	}

	return
}

func stdSortHelper[T ~int](data []T) (sortedData []T, useTime time.Duration) {
	var intSlice = make(intSlice, 0, len(data))
	for _, v := range data {
		intSlice = append(intSlice, int(v))
	}
	t := time.Now()
	sort.Sort(intSlice)
	useTime = time.Now().Sub(t)

	for _, v := range intSlice {
		sortedData = append(sortedData, T(v))
	}
	return
}

type intSlice []int

func (slice intSlice) Len() int {
	return len(slice)
}

func (slice intSlice) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice intSlice) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
