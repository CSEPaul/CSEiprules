package main

import (
	"fmt"

	"github.com/CSEPaul/CSEiprules/libs"
)

func main() {
	token := libs.TokenReader()
	fmt.Println(token)
}
