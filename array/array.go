package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("array", a)

	a[4] = 666
	fmt.Println("now array is", a)

	fmt.Println("length of array", len(a))

	b := [5]int{6, 7, 8, 9, 11}
	fmt.Println("another array is", b)

	var array2D [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			array2D[i][j] = 3*i + j
		}
	}
	fmt.Println("2d array", array2D)
}
