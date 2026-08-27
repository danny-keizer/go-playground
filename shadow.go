package main

import (
	"fmt"
)

func add(x, y int) (int, error) {
	return x + y, nil
}

func main() {
	var number int64 = 7
	fmt.Println("Like your never departing shadow.", number)
	fmt.Printf("Your final number is: %d", number)
}
