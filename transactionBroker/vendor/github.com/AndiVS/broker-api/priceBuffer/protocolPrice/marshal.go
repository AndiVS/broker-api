package protocolPrice

import "encoding/json"

// MarshalBinary Marshal currency to byte
func (c *Currency) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

// UnmarshalBinary Marshal currency to byte
func (c *Currency) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}
