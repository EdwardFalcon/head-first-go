package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0] = 42
	numbers[2] = 108
	var letters = [3]string{"a", "b", "c"}

	fmt.Println(numbers[0])
	fmt.Println(numbers[1])
	fmt.Println(numbers[2])
	fmt.Println(letters[0])
	fmt.Println(letters[1])
	fmt.Println(letters[2])

	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}

	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}

	for index, note := range notes {
		fmt.Println(index, note)
	}
}
