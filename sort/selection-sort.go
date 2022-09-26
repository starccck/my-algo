package sort

type SelSort[T ~int] struct{}

func (sort SelSort[T]) Sort(data []T) {
	selSort(data)
}

func selSort[T ~int](data []T) {
	for i := 0; i < len(data); i++ {
		minIdx := i
		for j := i; j < len(data); j++ {
			if data[j] < data[minIdx] {
				minIdx = j
			}
		}

		data[i], data[minIdx] = data[minIdx], data[i]
	}
}
