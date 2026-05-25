package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please rate our pizza between 1 and 5: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for rating %s", input)
	numRating, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		fmt.Println("Error converting string to int", err)
	}
	fmt.Println("Rating in int is ", numRating)
}
