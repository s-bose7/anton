package core

import (
	"fmt"
	"bytes"
 	"encoding/binary"
)

const (
	// Header Constants
	QR_QUERY = 0
	QR_RESPONSE = 1
	OPCODE_QUERY  = 0  	// Standard query
	OPCODE_IQUERY = 1  	// Inverse query
	OPCODE_STATUS = 2  	// Server status request
	RD = 1              // Recusion Desired
	AA = 0              // Authoritative Answer
	TC = 0              // Trucation
	RA = 0              // Recursion Available
	Z  = 0              // Reserved for future use. Must be zero in all queries and responses.
	
	// Header response code types
	RCODE_NO_ERROR = 0
	RCODE_FORMAT_ERROR = 1
	RCODE_SERVER_FAILURE = 2
	RCODE_NAME_ERROR = 3
	RCODE_NOT_IMPLEMENTED = 4
	RCODE_REFUSED = 5

	// Memory In Bytes
	SIZE_H = 12 
)

type DNSHeader struct {
	Id      uint16 
	Flags   uint16 
	QdCount uint16  	// The number of entries in the question section.
	AnCount uint16  	// The number of resource records in the answer section.
	NsCount uint16  	// The number of name server resource records in the authority section.
	ArCount uint16  	// The number of resource records in the additional section.
}

func (h *DNSHeader) Encode() ([]byte, error) {
	
	encodedHeader := new(bytes.Buffer)
	
	binary.Write(encodedHeader, binary.BigEndian, h.Id)
	binary.Write(encodedHeader, binary.BigEndian, h.Flags)
	binary.Write(encodedHeader, binary.BigEndian, h.QdCount)
	binary.Write(encodedHeader, binary.BigEndian, h.AnCount)
	binary.Write(encodedHeader, binary.BigEndian, h.NsCount)
	binary.Write(encodedHeader, binary.BigEndian, h.ArCount)

	return encodedHeader.Bytes(), nil
}

func DecodeHeader(data []byte) DNSHeader {
	var header DNSHeader
	reader := bytes.NewReader(data)

	binary.Read(reader, binary.BigEndian, &header.Id)
	binary.Read(reader, binary.BigEndian, &header.Flags)
	binary.Read(reader, binary.BigEndian, &header.QdCount)
	binary.Read(reader, binary.BigEndian, &header.AnCount)
	binary.Read(reader, binary.BigEndian, &header.NsCount)
	binary.Read(reader, binary.BigEndian, &header.ArCount)

	return header
}


func (h *DNSHeader) setFlags(qr, opcode, rd uint16) {
	h.Flags = (qr << 15) | (opcode << 11) | (rd << 8)
}


func NewDNSHeader() DNSHeader {
	header := DNSHeader{
		Id:      1209,
		QdCount: 1,
		AnCount: 0,
		NsCount: 0,
		ArCount: 0,
	}
	header.setFlags(QR_QUERY, OPCODE_QUERY, RD)
	return header
}


func (header *DNSHeader) PrintHeader() {
	fmt.Print("\n")
	fmt.Printf("Id:      %d\n", header.Id)
	fmt.Printf("Flag:    %d\n", header.Flags)
	fmt.Printf("QdCount: %d\n", header.QdCount)
	fmt.Printf("AnCount: %d\n", header.AnCount)
	fmt.Printf("NsCount: %d\n", header.NsCount)
	fmt.Printf("ArCount: %d\n", header.ArCount)
}