package main

import "fmt"

func main() {
	integers := []int{1, 5, 7, 9}

	for i := 0; i < len(integers); i++ {
		fmt.Println(integers[i])
	}

	strings := [3]string{"Hello", "World", "!"}

	for i, s := range strings {
		fmt.Printf("%d: %s\n", i, s)
	}
}
