package main

import "fmt"

func main() {
	const (
		Zero = iota
		One
		Two
		Three
		Four
	)

	fmt.Println(Zero, One, Two, Three, Four)
}
