package main

import "fmt"

func main() {
	var num []int
	num = make([]int, 100)

	for i := range num {
		num[i] = i + 1
	}

	ShowFizzBuzz(num)
}

//ShowFizzBuzz is function what writes to console
func ShowFizzBuzz(numbers []int) {
	for _, n := range numbers {
		if n%15 == 0 {
			fmt.Println("FizzBuzz")
		} else if n%3 == 0 {
			fmt.Println("Fizz")
		} else if n%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(n)
		}
	}
}
