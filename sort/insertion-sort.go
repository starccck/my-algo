package sort

type InsertionSortS[T ~int] struct{}

func (sort InsertionSortS[T]) Sort(data []T) {
	InsertionSort(data)
}

func InsertionSort[T ~int](data []T) {
	for i := 1; i < len(data); i++ {
		n := data[i]
		j := i
		for ; j > 0 && data[j-1] > n; j-- {
			data[j] = data[j-1]
		}
		data[j] = n
	}
}
