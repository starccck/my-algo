package sort

type MyIntType int

type SliceSortInterface[T ~int] struct {
	data []T
}

type MyIntSliceType []MyIntType

func (slice MyIntSliceType) Len() int {
	return len(slice)
}

func (slice MyIntSliceType) Less(i, j int) bool {
	return slice[i] < slice[j]
}

func (slice MyIntSliceType) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
