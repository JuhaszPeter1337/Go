package main

import "fmt"

func forLoop(){
	// var sum = 0
	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)
}

func whileLoop(){
	sum := 0
	// If you skip the init and post statements, you get a while loop.
	for sum < 5 {
		sum += 1
	}
	fmt.Println(sum)
}

func infiniteLoop(){
	sum := 0
	// If you skip the condition as well, you get an infinite loop.
	for {
		sum++
		fmt.Println(sum)
	}
}

func forEachRangeLoop(){
	strings := [] string {"hello", "world", "!"}
	length := len(strings)
	for i, s := range strings {
		if i < length - 1 {
			fmt.Printf("%d. %s ", i, s)
			// fmt.Print(i, ". ", s, " ")
		} else {
			fmt.Printf("%d. %s\n", i, s)
			// fmt.Print(i, ". ", s, "\n")
		}
	}
}

func forEachRangeLoop2(){
	strings := [] string {
		"hello",
		"world",
		"!",
	}
	for _, s := range strings {
		fmt.Println(s)
	}
}

func main(){
	var taskOne = "1. Watch Go crash course"
	var taskTwo = "2. Learn about Go language"
	var taskThree = "3. Go to the store"

	var testNumber = 5

	var taskItems = [] string {
		"1. Watch Go crash course", 
		"2. Learn about Go language",
	}

	var taskItems2 = [] string {taskOne, taskTwo}

	var taskItems3 = [2] string {
		taskOne, taskThree,
	}

	fmt.Println("##### Welcome to our TODO List Application! #####")
	fmt.Println(taskOne)
	fmt.Println(taskTwo)
	fmt.Println(taskThree)
	fmt.Println(testNumber)
	fmt.Println(taskItems)
	fmt.Println(taskItems2[0])
	fmt.Println(taskItems3)
	forLoop()
	whileLoop()
	forEachRangeLoop()
	forEachRangeLoop2()
}