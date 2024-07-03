package util

import (
	"fmt"
)


// Prints the byte slice in a readable format
func PrintBytes(name string, bytes []byte) {
	fmt.Println(name)
	for _, b := range bytes {
		fmt.Printf("%02x ", b)
	}
	fmt.Println()
}

func PrintStackTrace(msg string, err error) {
	fmt.Println(msg, err)
}
