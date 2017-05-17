package main

import (
//	"fmt"
//	"os"
)

/*func main() {
	str := os.Getenv("INPUT")
	if str != "" {
		fmt.Println(ReverseString(str))
	} else {
		fmt.Println("working with docker's scratch image sucks")
	}
}*/

func ReverseString(str string) string {
	if len(str) <= 1 {
		return str
	} else {
		return ReverseString(str[1:]) + str[:1]
	}
}

