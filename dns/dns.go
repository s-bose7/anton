package dns

import (
	"net"
	"resolve-on-go/core"
)

const MAX_PACKET_SIZE = 512


func BuildMessage(domain string) ([]byte, error) {
	message := core.DNSMessage{ 
		Header: core.NewDNSHeader(), Question: core.NewDNSQuestion(domain), 
	}
	encodedMessage, err := message.Encode()
	if err != nil {
		return nil, err
	}
	return encodedMessage, nil
}


func Query(server string, query []byte) ([]byte, error){
	return queryNameServer(server, query)
}

func ParseResponse(response []byte) (core.DNSMessage, error) {
	return core.Decode(response)
}


func queryNameServer(server string, query []byte) ([]byte, error) {
	conn, err := net.Dial("udp", server)
	if err != nil {
		return nil, err
	}
	// Ensures that the connection is closed when the function completes, 
	// whether it succeeds or fails.
	defer conn.Close() 
	
	_, err = conn.Write(query)
	if err != nil {
		return nil, err
	}
	// Buffer to store the response
	response := make([]byte, MAX_PACKET_SIZE)
	n, err := conn.Read(response)
	if err != nil {
		return nil, err
	}

	return response[:n], nil
}