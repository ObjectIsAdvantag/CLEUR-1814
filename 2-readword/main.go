package main

import (
	"fmt"
)

func readword() (word string) {
	fmt.Println("Type something")
	fmt.Scanf("%s", &word)
	return
}

func main() {
	w := readword()
	fmt.Println("You entered", w)
}
