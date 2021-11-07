package main

import "fmt"

func fizzbuzz(num int) string {
	switch 0 {
	case num % 15:
		return "fizz buzz"
	case num % 5:
		return "buzz"
	case num % 3:
		return "fizz"
	default:
		return fmt.Sprint(num)
	}
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(fizzbuzz(i))
	}
}
