package main

import (
	"fmt"
)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func loopOverRange() {
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}
}

func main() {
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	arr := [3]int{1, 2, 3}
	copyArr := arr
	copyArr[0] = 99

	fmt.Println("Array original:", arr) // [1 2 3]
	fmt.Println("Array copy:", copyArr) // [99 2 3]

	// --- Slices ---
	slice := []int{1, 2, 3}
	copySlice := slice
	copySlice[0] = 99

	fmt.Println("Slice original:", slice) // [99 2 3]
	fmt.Println("Slice copy:", copySlice) // [99 2 3]

	// --- Append creates a new backing array if needed ---
	newSlice := append(slice, 100)
	newSlice[0] = 7

	fmt.Println("After append:")
	fmt.Println("Original slice:", slice) // [99 2 3]
	fmt.Println("New slice:", newSlice)   // [7 2 3 100]
}
