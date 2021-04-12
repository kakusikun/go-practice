package main

import (
	"fmt"
	"math"
)

const s string = "const"

func main() {
	fmt.Println(s)

	const n = 999

	const d = n / 10
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
