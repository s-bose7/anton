package lib


type DNSResourceRecord struct {
	Name     []byte    // A domain name to which this resource record pertains.
	Type     uint16    // This field specifies the meaning of the data in the RDATA field. 
	Class    uint16    // The class of the data in the RDATA field.
	TTL      uint32    // The time interval (in seconds) that the resource record may be cached 
					   // before it should be discarded. 
	RDLength uint16    // Specifies the length in octets of the RDATA field.
	RData    []byte    // A a variable length string of octets that describes the resource.
}


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
	header := DecodeHeader(response[:12])
	question, err := DecodeQuestion(response[12:34])
	if err != nil {
		return DNSMessage{}, err
	}
	
	dnsMessageResponse := DNSMessage {
		Header: header,
		Question: question,
	}

	return dnsMessageResponse, nil
}