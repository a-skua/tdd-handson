package example

func Add(a, b int) int {
	return sum(a, b)
}

func Add3(a, b, c int) int {
	return sum(a, b, c)
}

func sum(ns ...int) (result int) {
	for _, n := range ns {
		result += n
	}
	return
}
