package main

import (
	"fmt"
	"os"
)

func main()  {
	programName := os.Args[0] 

	fmt.Printf("Hello %s \n", programName)

	if len(os.Args) > 1 {
		fmt.Println("Hello,", os.Args[1])
	}
}	