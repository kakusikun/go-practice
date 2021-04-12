package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["k1"] = 123
	m["k2"] = 666
	fmt.Println("map", m)

	v1 := m["k1"]
	fmt.Println("v1", v1)

	fmt.Println("length", len(m))

	delete(m, "k1")
	fmt.Println("map", m)

	if _, prs := m["k1"]; prs {
		fmt.Println("k1 exists")
	} else {
		fmt.Println("k1 does not exist")
	}

	m2 := map[string]int{"foo": 2, "bar": 1}
	fmt.Println("map 2", m2)

}
