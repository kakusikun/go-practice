package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum", sum)

	for i, num := range nums {
		sum += num
		fmt.Println(i, ". ", "sum", sum)
	}

	mp := map[string]int{"a": 1, "b": 6, "c": 100}
	for k, v := range mp {
		fmt.Println("key of", k, "has value of", v)
	}

}
