package main

func main() {
	testFn(6)
	// fmt.Println(testFn(6))
}

func simplest(x int) int {
	return x * 2
}

func testFn(x int) (int, int) {
	if x < 0 {
		x *= -1
	}
	return x, 1
}

func fnTwo(x int) int {
	y := 0
	for i := 0; i < x; i++ {
		y += i
	}
	return y
}
