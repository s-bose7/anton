package util

import (
	"fmt"
)


func PrintRaw(name string, bytes []byte) {
	fmt.Println(name, bytes)
	fmt.Println()
}

func PrintStackTrace(msg string, err error) {
	fmt.Println(msg, err)
}
