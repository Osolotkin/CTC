package main

import (
	"fmt"
)

const N int = 20;

func main() {

	var f func() int = fibonacci();

	for i := 0; i < N; i++ {
		fmt.Println(f());
	}

}

func fibonacci() func() int {

	var fn1 int = -1; // as f(n - 1)
	var fn2 int = -2; // as f(n - 2)

	return func() int {

		var fn int;
		// ternary, where are you (
		if fn1 < 1 {
			fn = fn1 + 1;
		} else {
			fn = fn1 + fn2;
		}

		fn2 = fn1;
		fn1 = fn;

		return fn;

	}

}
