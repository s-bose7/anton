package dns

import (
	. "resolve-on-go/lib"
	"resolve-on-go/core"
)

func BuildMessage(domain string) ([]byte, error) {
	message := DNSMessage{ 
		Header: NewDNSHeader(), Question: NewDNSQuestion(domain), 
	}
	encodedMessage, err := message.Encode()
	if err != nil {
		return nil, err
	}
	return encodedMessage, nil
}


func Query(server string, query []byte) ([]byte, error){
	return core.QueryNameServer(server, query)
}

func DecodeResponse(response []byte) (DNSMessage, error) {
	return Decode(response)
}