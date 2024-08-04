package main

import "fmt"

// ! Not Allowed
// token:= "123456"

// * Allowed
var token = "abc"
var token1 int = 123456

// ? Constants
const LoginToken string = "abcd"

//    ^ -> First character capital means PUBLIC

func main() {
	var username string = "John Doe"
	var isLoggedIn bool = true
	var smallValue uint8 = 255
	var floatValue float32 = 255.32132132132132132132132123
	var anotherVariable int
	var website = "golang.org"
	//         ^ -> Type Inference
	// Type is automatically assigned to the variable

	numberOfUsers := 10
	//            ^ -> Walrus Operator
	//  Used to declare variables and assign a value in a single line
	//  Can't be used outside a function

	fmt.Print(username)
	fmt.Printf(" is of type: %T \n", username)
	fmt.Print(isLoggedIn)
	fmt.Printf(" is of type: %T \n", isLoggedIn)
	fmt.Print(smallValue)
	fmt.Printf(" is of type: %T \n", smallValue)
	fmt.Print(floatValue)
	fmt.Printf(" is of type: %T \n", floatValue)
	fmt.Print(anotherVariable)
	fmt.Printf(" is of type: %T \n", anotherVariable)
	fmt.Print(website)
	fmt.Printf(" is of type: %T \n", website)
	fmt.Print(numberOfUsers)
	fmt.Printf(" is of type: %T \n", numberOfUsers)
	fmt.Print(token)
	fmt.Printf(" is of type: %T \n", token)
	fmt.Print(token1)
	fmt.Printf(" is of type: %T \n", token1)
	fmt.Print(LoginToken)
	fmt.Printf(" is of type: %T \n", LoginToken)
}
