package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("slice", s)

	s[0] = "w"
	s[1] = "t"
	s[2] = "f"
	fmt.Println("slice changes", s)

	fmt.Println("length", len(s))

	s = append(s, "!")
	s = append(s, "s", "h", "i", "t")
	fmt.Println("slice append", s)

	s2 := make([]string, len(s))
	copy(s2, s)
	fmt.Println("slice copy", s2)

	ss := s2[2:5]
	fmt.Println("subslice", ss)

	ss = s2[:5]
	fmt.Println("subslice from start", ss)

	ss = s2[2:]
	fmt.Println("subslice to end", ss)

	s3 := []string{"o", "h"}
	fmt.Println("init slice", s3)

	s2D := make([][]int, 2)
	for i := 0; i < 2; i++ {
		s2D[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			s2D[i][j] = 3*i + j
		}
	}
	fmt.Println("2d slice", s2D)

}
