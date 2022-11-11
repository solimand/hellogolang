package main

import "fmt"

// fibonacci (sum o prev 2 numbers, starting from 0 1) is a function that returns
// a function that returns an int.
// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, ...
func fibonacci() func() int {
	fibo := 0
	prev := 0
	startseq := false

	return func() int {
		if (fibo == 0) && (prev == 0) && !startseq {
			fibo = 1
			return 0
		} else if (fibo == 1) && (prev == 0) {
			if !startseq {
				startseq = true
				return fibo
			} else {
				res := prev + fibo
				prev = 1
				return res
			}
		} else {
			res := prev + fibo
			prev = fibo
			fibo = res
			return res
		}
	}
}

// shorter version with shift [a,b] in [b,a+b]
// starts from 1,1
func fibonacciShort() func() int {
	fibo := 0
	prev := 1
	return func() int {
		fibo, prev = prev, fibo+prev
		return fibo
	}
}

func main() {
	// f := fibonacci()
	fs := fibonacciShort()
	for i := 0; i < 13; i++ {
		// fmt.Println(f())
		fmt.Println(fs())
	}
}
