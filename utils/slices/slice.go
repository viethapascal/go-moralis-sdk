package slices

type CompareFunc func(field string, value interface{}) bool

func FilterObject[T any](s []T, f func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range s {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

func FirstObject[T any](s []T, f func(T) bool) *T {
	for _, v := range s {
		if f(v) {
			return &v
		}
	}
	return nil
}
