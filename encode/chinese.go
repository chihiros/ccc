package encode

import "strings"

// Legacy multi-byte Chinese (simplified) encodings
// Charactor code: GBK
func (e *Encode) GBK() ([]byte, error) {
	b, err := encodeReader(strings.NewReader(e.Target), "gbk")
	return b, err
}

// Charactor code: gb18030
func (e *Encode) GB18030() ([]byte, error) {
	b, err := encodeReader(strings.NewReader(e.Target), "gb18030")
	return b, err
}

// Legacy multi-byte Chinese (traditional) encodings
func (e *Encode) Big5() ([]byte, error) {
	b, err := encodeReader(strings.NewReader(e.Target), "big5")
	return b, err
}
