package utils

func Map[T any, U any](intput []T, fn func(T) U) []U {
	r := make([]U, 0, len(intput))
	for _, v := range intput {
		r = append(r, fn(v))
	}
	return r
}
