package vector

func Reverse(v []int) {
	l := len(v)
	for i := 0; i < l/2; i++ {
		v[i], v[l-1-i] = v[l-1-i], v[i]
	}
}
