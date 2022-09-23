package sort

type SortI[T comparable] interface {
	Sort(data []T)
}
