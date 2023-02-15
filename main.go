// Filename: main.go
package main

//demonstrate flags

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	//set flags
	msg := flag.String("msg", "Howdy, stranger!", "the message to display")
	num := flag.Int("num", 1, "how many times to print the message")
	caps := flag.Bool("U", false, "specify whether to uppercase the string (true or false)")
	reverse := flag.Bool("r", false, "reverse the string (true or false)")
	flag.Parse()
	//check if it should be uppercase before printing it
	if *caps {
		*msg = strings.ToUpper(*msg)
	}
	//check if we should reverse the string
	if *reverse {
		var result string
		for _, value := range *msg {
			result = string(value) + result
		}
		//write the reversed string to msg
		*msg = result
	}
	//print the string
	for i := 0; i < *num; i++ {
		fmt.Println(*msg)
	}
}
