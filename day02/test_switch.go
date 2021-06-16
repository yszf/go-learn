package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("1. case is false")
		fallthrough
	case true:
		fmt.Println("2. case is true")
		fallthrough
	case false:
		fmt.Println("3. case is false")
		fallthrough
	case true:
		fmt.Println("4. case is true")
	default:
		fmt.Println("case is default")
	}

	var grade string = "D"
	var score int = 90

	switch score {
	case 100, 90:
		grade = "A"
	case 80, 70:
		grade = "B"
	case 60:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println("grade =", grade)

	var x interface{}
	switch x.(type) {
	case int:
		fmt.Println("int")
	case float32, float64:
		fmt.Println("float32 or floate64")
	case nil:
		fmt.Println("nil")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("未知类型")
	}
}
