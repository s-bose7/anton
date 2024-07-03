package lib

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
	var name []byte
	for {
		length, err := reader.ReadByte()
		if err != nil {
			return question, err
		}
		if length == 0 {
			break
		}

		label := make([]byte, length)
		if _, err := reader.Read(label); err != nil {
			return question, err
		}
		name = append(name, length)
		name = append(name, label...)
	}
	name = append(name, 0) // Append the null byte at the end of the name
	question.Name = name

	// Decode the type
	if err := binary.Read(reader, binary.BigEndian, &question.Type); err != nil {
		return question, err
	}
	// Decode the class
	if err := binary.Read(reader, binary.BigEndian, &question.Class); err != nil {
		return question, err
	}

	return question, nil
}


func NewDNSQuestion(domain string) DNSQuestion {
	encodedDomain := encodeString(domain)
	return DNSQuestion{
		Name:  encodedDomain,
		Type:  1, // IPv4
		Class: 1, 
	}
}

func (question *DNSQuestion) PrintQuestion() {
	fmt.Print("\n")
	name, err := decodeString(question.Name)
	if err != nil {
		return
	}
	fmt.Println("Name:", name)
	fmt.Printf("Type: %d\n",  question.Type)
	fmt.Printf("Class: %d\n", question.Class)
}