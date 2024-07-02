package main

/*
Copyright (c) 2024 s-bose7

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

import "fmt"
import "bytes"


// Encodes a domain name into the DNS wire format.
func encodeDNSName(name string) []byte {
	var buffer bytes.Buffer
	parts := bytes.Split([]byte(name), []byte(".")) 
	for _, part := range parts { 
		buffer.WriteByte(byte(len(part)))
		buffer.Write(part)
	}
	// The domain name terminates with the zero length octet f
	// for the null label of the root.
	buffer.WriteByte(0)
	return buffer.Bytes() 
}

// Prints the byte slice in a readable format
func printByteSlice(name string, b []byte) {
	fmt.Println(name)

	for _, byte := range b {
		fmt.Printf("%02x", byte)
	}
	fmt.Println()
}

func main() {

	domain := "dns.google.com"
	fmt.Println("Provided domain: ", domain)
	encodedName := encodeDNSName(domain)
	printByteSlice("Encoded domain:", encodedName)
	
}
