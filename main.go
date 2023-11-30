package main

import (
	"fmt"
	"runtime"

	"github.com/hernan0073/testgolang/variables"
)

func main() {
	estado, texto := variables.ConviertetoTex(1544)

	fmt.Println("convert", estado, texto)

	os := runtime.GOOS

	if os == "Linux ." {

		fmt.Println(" Esto es linux")

	} else {

		fmt.Println("esto es", os)
	}

}
