package main

import "fmt"

func main() {
	fmt.Printf("%.2f", 3.6785)
	fmt.Println()

	fmt.Printf("%T", 3.6785)
	fmt.Println()

	// \t - istab character
	fmt.Printf("%s:\t%d", "Decimal", 11)
	fmt.Println()

	fmt.Printf("%s:\t%b", "Binary", 11)
	fmt.Println()

	// \n - is new line character
	fmt.Printf("%s:\t%o\n", "Octal", 11)

	fmt.Printf("%s:\t%o\n", "Hexadecimal", 11)

}
