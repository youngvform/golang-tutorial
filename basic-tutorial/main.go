package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func multiply(a int, b int) int {
	return a * b
}

func nakedLenAndUpper(word string) (length int, name string) {
	defer fmt.Println("done!")
	length = len(word)
	name = strings.ToUpper(word)
	return
}

func superAdd(numbers ...int) int {
	sum:= 0
	for _, number := range numbers{
		sum += number
	}

	return sum
}

type person struct {
	name string
	age int
	job []string
}

func main() {
	length, name := lenAndUpper("start")
	fmt.Println(length, name)
	result := multiply(2, 3);
	fmt.Println(result);

	len2, name2 := nakedLenAndUpper("newStart")
	fmt.Println(len2)
	fmt.Println(name2)

	total := superAdd(1,2,3,4,5,6,7)
	println(total)

	arr := []string{"1","2","3"}
	fmt.Println(arr)

	obj := map[string]int{"name": 1, "task": 2}
	fmt.Println(obj["name"])

	job := []string{"developer", "watcher", "doctor"}
	p := person{"kim", 21, job}
	fmt.Println(p)
}
