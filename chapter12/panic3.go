package main

import "fmt"

func main() {
	one()
}
func one() {
	defer fmt.Println("deffered in one()")
	two()
}
func two() {
	defer fmt.Println("deffered in two()")
	panic("Let's see what's been deferred!")
}
