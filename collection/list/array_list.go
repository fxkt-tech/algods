package list

func ReverseArray[T any](v []T) {
	l := len(v)
	for i := 0; i < l/2; i++ {
		v[i], v[l-1-i] = v[l-1-i], v[i]
	}
}
