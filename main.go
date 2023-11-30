package main

import (
	"fmt"

	"github.com/hernan0073/testgolang/variables"
)

func main() {
	estado, texto := variables.ConviertetoTex(1544)

	fmt.Println("convert", estado, texto)

}
