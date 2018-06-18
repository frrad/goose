package main

func main() {
}

func simplest(x int) int {
	return x * 2
}

func dagFn(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func fnTwo(x int) int {
	y := 0
	for i := 0; i < x; i++ {
		y += i
	}
	return y
}
