package encode

import "strings"

// The Encoding
// Charactor code: UTF-8
func (e *Encode) UTF8() ([]byte, error) {
	b, err := encodeReader(strings.NewReader(e.Target), "utf-8")
	return b, err
}
