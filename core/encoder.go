package core

import "bytes"

func encodeToByteArray(domain string) []byte {
	var buf bytes.Buffer
	for _, label := range bytes.Split([]byte(domain), []byte(".")) {
		buf.WriteByte(byte(len(label)))
		buf.Write(label)
	}
	buf.WriteByte(0) // End of domain name
	return buf.Bytes()
}

func decodeToString(encoded []byte) (string, error) {
	var domain bytes.Buffer
	reader := bytes.NewReader(encoded)

	for {
		length, err := reader.ReadByte()
		if err != nil {
			return "", err
		}
		if length == 0 {
			break
		}

		label := make([]byte, length)
		if _, err := reader.Read(label); err != nil {
			return "", err
		}

		if domain.Len() > 0 {
			domain.WriteByte('.') // Add dot separator
		}
		domain.Write(label)
	}

	return domain.String(), nil
}
