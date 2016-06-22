package main

import (
	"fmt"
	"os"
)

func print(whatever interface{}) {
	fmt.Fprintf(os.Stdout, "%v\n", whatever)
}

func makeFibbonaciFn() func(n int) int {
	m := make(map[int]int)
	m[1] = 1;
	m[2] = 1;

	// apparently Go doesn't support named functions in functions
	// nor anonymous recursive functions
	var fibb func(n int) int;
	fibb = func(n int) int {
		_, ok := m[n]

		if (!ok) {
			m[n] = fibb(n-1) + fibb(n-2)
		}

		return m[n]
	}

	return fibb
}

func main() {
	myfibb := makeFibbonaciFn()

	for i := 1; i < 10; i+= 1 {
		print(myfibb(i))
	}
}