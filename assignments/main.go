package main

import "fmt"

func main() {
	numbers := makeRange(0, 10)

	for _, i := range numbers {
		if i%2 == 0 {
			fmt.Printf("%d is even\n", i)
		} else {
			fmt.Printf("%d is odd\n", i)
		}
	}
}

func makeRange(min int, max int) []int {
	r := []int{}

	for i := min; i <= max; i++ {
		r = append(r, i)
	}

	return r
}
