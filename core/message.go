package core

import "fmt"


type DNSMessage struct {
	Header      DNSHeader
	Question    DNSQuestion
	Answer     []DNSResourceRecord
	Authority  []DNSResourceRecord
	Additional []DNSResourceRecord
}


func (m *DNSMessage) Encode() ([]byte, error) {
	header, err := m.Header.Encode()
	if err != nil {
		return nil, err
	}

	question, err := m.Question.Encode()
	if err != nil {
		return nil, err
	}

	message := append(header, question...)
	return message, nil
}

func Decode(response []byte) (DNSMessage, error) {
	rangeH := []int {0, SIZE_H}
	rangeQ := []int {SIZE_H, SIZE_H+sizeOfDomainBytes+4}

	header := DecodeHeader(response[rangeH[0]:rangeH[1]])
	question, err := DecodeQuestion(response[rangeQ[0]:rangeQ[1]])
	if err != nil {
		return DNSMessage{}, fmt.Errorf("failed to decode question: %w", err)
	}

	dnsMessageResponse := DNSMessage {
		Header: header,
		Question: question,
	}

	return dnsMessageResponse, nil
}

func (m DNSMessage) Print() {
	fmt.Print("\n")
	fmt.Printf("DNSMessageState: %+v\n", m)
}