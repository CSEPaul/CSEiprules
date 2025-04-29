package main

import (
	"fmt"

	"github.com/CSEPaul/CSEiprules/libs"
)

func main() {
	// Get the token from the environment variable
	token := libs.TokenReader()

	// Check if the token is valid length and alphanumeric
	if !libs.TokenLengthChecker(token) {
		fmt.Println("Token length is not valid")
		return
	}
	if !libs.TokenAlphanumericChecker(token) {
		fmt.Println("Token contains non-alphanumeric characters")
		return
	}

	fmt.Println("Token: ", token)

}
