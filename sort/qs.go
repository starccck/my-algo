package sort

type QS[T ~int] struct{}

func (sort QS[T]) Sort(data []T) {
	quickSort(data)
}

// type IQS QS[MyIntType]

// var iqsV IQS

// func (q IQS) Sort(data []MyIntType) {
// 	if len(data) <= 1 {
// 		return
// 	}

// 	pivot := partition(data, 0, len(data)-1)
// 	q.Sort(data[:pivot])
// 	q.Sort(data[pivot+1:])
// }

func quickSort[T ~int](data []T) {
	if len(data) <= 1 {
		return
	}

	pivot := partition(data, 0, len(data)-1)
	quickSort(data[:pivot])
	quickSort(data[pivot+1:])
}

func partition[T ~int](data []T, left, right int) (parIdx int) {
	pivot := getPivot(left, right)
	data[0], data[pivot] = data[pivot], data[0]
	flag := data[0]

	l, r := 1, right
	for l <= r {
		for l <= r && data[l] < flag {
			l++
		}

		for l <= r && data[r] >= flag {
			r--
		}

		if l >= r {
			break
		}

		data[l], data[r] = data[r], data[l]
	}

	data[0], data[l-1] = data[l-1], data[0]

	return l - 1
}

func getPivot(left, right int) (pivot int) {
	return (left + right) / 2
}
