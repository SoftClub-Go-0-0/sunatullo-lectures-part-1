package main

import "fmt"

func main() {
	var averageScore int
	fmt.Scan(&averageScore)

	if averageScore < 80 {
		fmt.Println("Sorry, you can not continue the course")
	} else if averageScore > 100 {
		fmt.Println("Hmm... I think you've made a mistake. Enter the number from 0 to 100")
	} else {
		fmt.Println("Congrats! You passed to the next stage.")
	}
}
