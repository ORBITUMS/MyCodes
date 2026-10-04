package main

import "fmt"

func main() {
	var input string
	fmt.Print("Введите текст: ")
	fmt.Scanln(&input)

	const a string = "("
	const b string = ")"

	fmt.Println( )
	// if 3 [ ((())), (())(),()(()), ()()()]
	// aaabbb aabbab abaabb ababab
}