package main

import (
	"errors"
	"fmt"
	"math"
)

type person struct {
	name string
	age  int
}

func main() {
	x := 5
	y := 10
	sum := x + y

	if sum < 10 {
		fmt.Println("Less than 10")
	} else {
		fmt.Println("Not less than 10")
	}

	if sum > 10 {
		fmt.Println("More than 10")
	} else {
		fmt.Println("Not more than 10")
	}

	fmt.Println(sum)

	a := []int{5, 4, 3, 2, 1}
	a = append(a, 13)
	fmt.Println(a)

	vertices := make(map[string]int)
	vertices["Triangle"] = 2
	vertices["Square"] = 3
	vertices["Dodgecoin"] = 12

	fmt.Println(vertices)

	delete(vertices, "Square")

	fmt.Println(vertices)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	m := make(map[string]string)
	m["a"] = "alpha"
	m["b"] = "beta"

	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}

	resultOfFunc := add(2, 3)
	fmt.Println("result:", resultOfFunc)

	result, err := sqrt(-16)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	p := person{name: "Bryant", age: 39}
	fmt.Println(p)
	fmt.Println(p.age)
	fmt.Println(p.name)

	i := 7
	inc(&i)
	fmt.Println(i)

}

func inc(x *int) {
	*x++
}

func add(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined")
	}
	return math.Sqrt(x), nil
}
