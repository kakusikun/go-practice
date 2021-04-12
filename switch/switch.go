package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("i is", i)
	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("weekday")
	}

	t := time.Now().Add(time.Hour * 12)
	switch {
	case t.Hour() < 12 && t.Hour() >= 6:
		fmt.Println("morning")
	case t.Hour() == 12:
		fmt.Println("noon")
	case t.Hour() > 12 && t.Hour() < 18:
		fmt.Println("afternoon")
	case t.Hour() >= 18 && t.Hour() <= 23:
		fmt.Println("night")
	case t.Hour() >= 0 && t.Hour() < 6:
		fmt.Println("midnight")
	}

	checkType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("unknown type %T\n", t)
		}
	}

	checkType(true)
	checkType(666)
	checkType("WTF")
}
