package main

import (
	"fmt"
	"log"
)

func perimeter(length, width float64) (float64, error) {
	if length < 0 {
		return 0, fmt.Errorf("a lenght of %0.2f is invalid", length)
	}
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	return (length + width) * 2, nil
}

func main() {
	p, err := perimeter(8.2, -10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}
