package core


type DNSResourceRecord struct {
	Name     DNSQuestion
	TTL      uint32    // The time interval (in seconds) that the resource record may be cached 
					   // before it should be discarded. 
	RDLength uint16    // Specifies the length in octets of the RDATA field.
	RData    []byte    // A a variable length string of octets that describes the resource.
}



// func decodeResource(data []byte, startPosition int) (DNSResourceRecord, error) {

// }
