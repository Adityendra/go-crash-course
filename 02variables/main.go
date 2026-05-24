package main

import "fmt"

const LoginToken string = "123456789" //caps is public variable and can be accessed outside the package

func main() {
	var name string = "Adi"
	fmt.Println(name)
	fmt.Printf("Variable of type: %T \n", name)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable of type: %T \n", smallValue)

	//Implicit type
	var test = "Test"
	fmt.Println(test)
	fmt.Printf("Variable of type: %T \n", test)

	//no var style walrus operator only inside functions
	noOfUsers := 3000
	fmt.Println(noOfUsers)
	fmt.Printf("Variable of type: %T \n", noOfUsers)

	fmt.Println(LoginToken)
}
