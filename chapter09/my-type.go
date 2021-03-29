package main

import "fmt"

type MyType string

func (m MyType) sayHi() {
	fmt.Println("Hi from", m)
}

func (m MyType) WithReturn() int {
	return len(m)
}

func main() {
	value := MyType("a MyType value")
	value.sayHi()
	fmt.Println(value.WithReturn())
	anotherValue := MyType("another value")
	anotherValue.sayHi()
}
