package main

import "fmt"

func sumAll(numbers ...int) int { //...int adalah vararg atau variable args
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 10, 10, 10, 10, 10)
	fmt.Println(total)

	// ketika ada slice kita masih bisa menggunakan variadic func
	numbers := []int{10,10,10,10}
	fmt.Println(sumAll(numbers...)) //cara conversi
}
