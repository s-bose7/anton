package core

import (
	"fmt"
	"bytes"
 	"encoding/binary"
)


type DNSQuestion struct {
	Name  []byte 		// A domain name represented as a sequence of labels.
	Type  uint16 		// The type of the query
	Class uint16 		// The class of the query.
}

var (
	sizeOfDomainBytes = 0
)

func (q *DNSQuestion) Encode() ([]byte, error) {
	buf := new(bytes.Buffer)
	
	buf.Write(q.Name)
	binary.Write(buf, binary.BigEndian, q.Type)
	binary.Write(buf, binary.BigEndian, q.Class)

	return buf.Bytes(), nil
}

func DecodeQuestion(data []byte) (DNSQuestion, error) {
	var question DNSQuestion
	reader := bytes.NewReader(data)

	// Decode the domain name
	name := make([]byte, sizeOfDomainBytes)
	n, err := reader.Read(name)
	if err != nil {
		return DNSQuestion{}, fmt.Errorf("failed to read domain name: %w", err)
	}
	question.Name = name[:n]
	// Decode the type
	if err := binary.Read(reader, binary.BigEndian, &question.Type); err != nil {
		return question, fmt.Errorf("failed to read question type: %w", err)
	}
	// Decode the class
	if err := binary.Read(reader, binary.BigEndian, &question.Class); err != nil {
		return question, fmt.Errorf("failed to read question class: %w", err)
	}

	return question, nil
}


func NewDNSQuestion(domain string) DNSQuestion {
	encodedDomain := encodeToByteArray(domain)
	sizeOfDomainBytes = len(encodedDomain)
	return DNSQuestion{
		Name:  encodedDomain,
		Type:  1, // IPv4
		Class: 1, // 'IN'
	}
}

func (question *DNSQuestion) PrintQuestion() {
	fmt.Print("\n")
	name, err := decodeToString(question.Name)
	if err != nil {
		return
	}
	fmt.Println("Name:", name)
	fmt.Printf("Type: %d\n",  question.Type)
	fmt.Printf("Class: %d\n", question.Class)
}