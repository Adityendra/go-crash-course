package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcomeMsg := "Welcome user to GoLang"
	fmt.Println(welcomeMsg)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello %s, nice to meet you! \n", name)
}
