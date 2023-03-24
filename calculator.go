package bobbygogo

func Cross(n ...int) int {
	a := 0
	for _, val := range n {
		a += val
	}
	return a
}
