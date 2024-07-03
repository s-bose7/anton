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

import (
	"fmt"
	"resolve-on-go/util"
	"resolve-on-go/dns"
	. "resolve-on-go/lib"
)

func printHeader(dnsHeader DNSHeader) {
	fmt.Print("\n")
	fmt.Printf("Id:      %d\n", dnsHeader.Id)
	fmt.Printf("Flag:    %d\n", dnsHeader.Flags)
	fmt.Printf("QdCount: %d\n", dnsHeader.QdCount)
	fmt.Printf("AnCount: %d\n", dnsHeader.AnCount)
	fmt.Printf("NsCount: %d\n", dnsHeader.NsCount)
	fmt.Printf("ArCount: %d\n", dnsHeader.ArCount)
}

func printQuestion(dnsQuestion DNSQuestion) {
	fmt.Print("\n")
	name, err := util.DecodeString(dnsQuestion.Name)
	if err != nil {
		return
	}
	fmt.Println("Name:", name)
	fmt.Printf("Type: %d\n", dnsQuestion.Type)
	fmt.Printf("Class: %d\n", dnsQuestion.Class)
}

func main() {

	domain := "dns.google.com"
	
	query, err := dns.BuildMessage(domain)
	if err != nil {
		util.PrintStackTrace("Error building query: ", err)
		return
	}
	util.PrintBytes("Encoded DNS Message:", query)
	
	server := "8.8.8.8:53"
	response, err := dns.Query(server, query)
	if err != nil {
		util.PrintStackTrace("Error Querying Google's DNS server: ", err)
		return
	}
	util.PrintBytes("\nGoogle's DNS server response:", response)
	
	resolvedMessage, err := DecodeResponse(response)
	if err != nil {
		util.PrintStackTrace("Error parsing response: ", err)
		return
	}

	printHeader(resolvedMessage.Header)
	printQuestion(resolvedMessage.Question)
}

