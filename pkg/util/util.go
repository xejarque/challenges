package util

func GetPointer[T any](t T) *T {
	return &t
}
