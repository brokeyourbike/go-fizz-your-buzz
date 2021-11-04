package main

import "fmt"

func fizzbuzz(num int) string {
	switch {
	case num%15 == 0:
		return "fizz buzz"
	case num%5 == 0:
		return "buzz"
	case num%3 == 0:
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
