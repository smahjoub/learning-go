package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#008000",
		"blue":  "#0000FF",
	}

	for color, hex := range colors {
		fmt.Printf("[%v]=%v\n", color, hex)
	}

	anotherMap := make(map[int]string)

	anotherMap[0] = "hello"

	anotherMap[1] = "world"

	fmt.Println(colors)

	fmt.Println(anotherMap)

	delete(anotherMap, 1)

	fmt.Println(anotherMap)
}
