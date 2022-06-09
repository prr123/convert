// convert.go
// author pr azul software
// date: 9. June 2022
// update:
// copyright prr, azul softwre
// license: see github
//

package main

import (
	"fmt"
	"os"
	"convert/cvtLib"
	)


func main () {
// function that parses a json structure and create a go structure

	var inpfilNam, gofilNam string

	numArg := len(os.Args)

	switch numArg {
	case 1:
		fmt.Println("insufficient arguments!")
		fmt.Printf("Usage is:\n  %s infile gofile!\n", os.Args[0])
		os.Exit(1)

	case 2:
		inpfilNam = os.Args[1]
	case 3:
		inpfilNam = os.Args[1]
		gofilNam = os.Args[2]

	default:
		fmt.Printf("too many arguments in cmd line: %d!", numArg)
		fmt.Printf("Usage is:\n  %s infile gofile!\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Println("******** Convert ********")
	fmt.Printf("input file: %s\n", inpfilNam)
	fmt.Printf("go file:    %s\n", gofilNam)

	err := cvtLib.CreGoFil(inpfilNam, gofilNam)
	if err != nil {
		fmt.Printf("error CreGoFil: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("**** Convert Success ****")
}
